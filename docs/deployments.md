# Deployments (release-artifact model)

> Status: **draft / iterating.** Defines how the prover is built, released, and deployed
> deterministically across chains. Not yet implemented.

## Principle

**Deploy bytes, not source.** A tagged release is built once into an immutable bytecode
artifact; every chain deployment consumes that exact artifact. The compiler, `foundry.toml`,
and source tree are out of the deploy path entirely, so a contract's CREATE2 address is fixed
forever and cannot drift from a compiler bump, config edit, comment, or refactor.

The corollary — and the reason the existing deploy script must be redone:

> **The deploy step MUST NOT recompile the contract.** It reads the frozen `creationCode`
> from the release and CREATE2-deploys it. It must never reference `type(Contract).creationCode`
> or run `forge build` against the source — doing so silently puts the compiler back in the
> deploy path and reintroduces drift.

## Compiler recipe (metadata-free)

The `default` profile uses **`bytecode_hash = "none"` + `cbor_metadata = false`** — the same
metadata-free recipe as `infra/contract-deploys-v2` / ProverV2. This strips the source hash from
the bytecode, so **comments, formatting, NatSpec, and variable names never change the CREATE2
address — only a real logic (opcode) change does.** (Verified: adding a comment leaves
`keccak256(creationCode)` byte-identical.) A constant `solc`-version tag remains in the trailer;
it's source-independent, so immunity holds. `bytecode_hash="none"` is the effective lever;
`cbor_metadata=false` is kept for parity with the v2 recipe.

> ⚠️ **One-time relaunch.** This is *not* the recipe that produced the existing v1 addresses
> (`0x85e9…`, `0x95cc…`) — those used `bytecode_hash="ipfs"` and are permanently welded to their
> exact source bytes; this recipe does **not** reproduce them. Adopting it means a deliberate,
> one-time move to **new** canonical addresses. We're prepping this now, not executing it (TBD).
> To verify/reproduce the legacy v1 contracts, use the `ipfs` recipe (still live in
> `infra/contract-deploys`).

## Address model

```
address = keccak256(0xff ++ factory ++ salt ++ keccak256(initCode))
initCode = creationCode (frozen, from release)  ++  abi.encode(constructorArgs)
```

CREATE2 has no `chainId`, so the address is identical on every chain for a fixed
`(factory, salt, initCode)`. What this means here:

- **`creationCode`** is frozen per release and is **environment-independent** (one artifact
  serves testnet and mainnet).
- **Constructor args differ by environment** (different sequencer address for testnet vs
  mainnet). Different args → different `initCode` → **a distinct canonical address per
  environment**, but consistent across all chains within that environment.
- So there is one canonical **testnet** prover address and one canonical **mainnet** prover
  address, each stable across all chains in its env.

### Canonical addresses & per-chain overrides

The current canonical per-env addresses (the **v1 / `ipfs`** line) are the source of truth in
**infra**, at `overlays/<env>/contracts-configmap-patch.yaml` → `PROVER_V2_ADDRESS`. The
metadata-free relaunch (see Compiler recipe) will mint **new** canonical addresses, recorded
the same way:

| env | canonical `PROVER_V2_ADDRESS` |
|---|---|
| devnet (testnet) | `0x85e9506fd24F9B588dcf2A5AaEF7069e34D99fCE` |
| mainnet | `0x95ccEAE71605c5d97A0AC0EA13013b058729d075` |
| shadownet | `0x584dD7A65796c9245ABEF95EEbf8029B4255155F` |

The same file also holds **per-chain overrides** (`PROVER_V2_ADDRESS_MEGAETH`,
`_REDBELLY`, `_TRON`, …) that differ from the canonical address. These exist because those
chains **did not get the canonical CREATE2 address** — typically the canonical Arachnid
factory `0x4e59…4956C` isn't deployed there (so the old script's vanilla-deploy fallback fired,
giving a nonce-based address), or the chain is non-EVM (TRON). This is exactly the
factory-presence failure mode below, already visible in production. The new deploy must treat a
missing factory as a hard stop (deploy the factory first, or record a deliberate override) — not
a silent fallback.

## What's frozen vs. supplied at deploy

| Item | Where it comes from | Frozen? |
|---|---|---|
| `creationCode` (compiled bytecode) | release artifact | **yes** (per tag) |
| compiler version + settings | release manifest | **yes** (per tag) |
| factory address | constant `0x4e59…4956C` | yes |
| `clientType` (`"proof_api"`) | constant | yes |
| `sequencer` (per-env pubkey) | **env var at deploy** | no — env-specific |
| `peptideChainId` | **env var at deploy** | no |
| `CREATE2_SALT` | **env var at deploy** (set by infra) | no |
| RPC URL, deployer key | **env var at deploy** (set by infra) | no |

All runtime inputs are **read from the environment** — this repo provides the deploy logic;
the infra repo (or any caller) provides salt, args, RPC, and key when it invokes the action.
The deployer key does **not** affect the address (CREATE2 goes through the factory), so it can
differ per chain.

## Release artifacts

A release for tag `vX.Y.Z` contains, per deployable contract (prover scope — see Non-goals):

- `creationCode.hex` — the frozen creation bytecode (no constructor args).
- `creationCodeHash` — `keccak256(creationCode)`, the immutability anchor.
- `abi.json` — for interaction.
- `metadata.json` / standard-JSON input — for explorer verification. With `bytecode_hash="none"`
  there's no embedded metadata hash, so explorers verify from the standard-JSON source rather than
  by an automatic metadata match.
- `manifest.json` — see below.

`manifest.json` (env-independent; **no addresses**, since salt + args are runtime inputs):

```jsonc
{
  "tag": "vX.Y.Z",
  "commit": "<sha>",
  "compiler": {
    "solc": "0.8.15",
    "profile": "default",
    "via_ir": false, "optimizer": true, "optimizer_runs": 200,
    "evm_version": "london", "bytecode_hash": "none", "cbor_metadata": false
  },
  "factory": "0x4e59b44847b379578588920cA78FbF26c0B4956C",
  "contracts": {
    "CrossL2ProverV2": {
      "creationCodeHash": "0x…",
      "constructorSignature": "constructor(string,address,bytes32)",
      "creationCode": "CrossL2ProverV2.creationCode.hex"
    }
  }
}
```

The expected on-chain **address is derived at deploy** from `creationCode + env args + env salt`
and recorded per-env in infra (see Decisions).

## Release workflow (on tag push `vX.Y.Z`)

1. Build the prover under the `default` profile (the deploy recipe).
2. Extract `creationCode`, compute `creationCodeHash`, collect ABI + verification metadata.
3. Emit `manifest.json` + a `SHA256SUMS` checksum file over all artifacts.
4. **Reproducibility check (CI):** re-run the build from the tag and assert
   `keccak256(creationCode)` matches the manifest. This — not the uploaded blob — is the
   immutability guarantee; anyone can re-derive and verify.
5. Publish a GitHub Release with the artifacts, manifest, checksums, and notes.

Immutability also requires **protected release tags** (no force-push); a moved tag silently
changes the source of truth.

## Deploy workflow (`deploy vX.Y.Z → chain`)

Bytes-in, no recompile:

1. Download the release artifact for `vX.Y.Z`; verify against `SHA256SUMS`.
2. Build `initCode = creationCode ++ abi.encode("proof_api", $SEQUENCER, $PEPTIDE_CHAIN_ID)`.
3. Compute predicted address `keccak256(0xff ++ factory ++ $CREATE2_SALT ++ keccak256(initCode))`.
   If `$EXPECTED_ADDRESS` is provided, **assert it matches** and abort otherwise.
4. **Factory presence:** require the canonical CREATE2 factory `0x4e59…4956C` on the chain
   (the Arachnid deployer, `code.length == 69`). If absent, deploy it first (its keyless
   pre-signed tx) — never silently fall back to a nonce-based `new` (that yields a different,
   non-deterministic address).
5. If code already exists at the predicted address, skip.
6. Deploy by sending `$CREATE2_SALT ++ initCode` to the factory
   (`cast send 0x4e59…4956C 0x<salt><initCode>`), **not** via a recompiling `forge script`.
7. Assert deployed address == predicted; verify on the chain's explorer using the bundled
   standard-JSON.

## Inputs (env contract between infra and this repo)

The deploy entrypoint reads:

| Env var | Meaning |
|---|---|
| `CREATE2_SALT` | bytes32 salt (set by infra) |
| `SEQUENCER_PUB_KEY` | env-specific sequencer address |
| `PEPTIDE_CHAIN_ID` | bytes32 peptide chain id |
| `RPC_URL` | target chain RPC |
| `DEPLOYER_PRIVATE_KEY` | funded key (does not affect address) |
| `EXPECTED_ADDRESS` (optional) | canonical per-env address to assert against — infra supplies it from `overlays/<env>/contracts-configmap-patch.yaml` (`PROVER_V2_ADDRESS`) |

`CREATE2_SALT` is held as an env var/secret in the infra GitHub repo (fixed per env), which is
what keeps the canonical address stable and predictable. This repo never hardcodes it.

## Scope & non-goals

- **Release artifacts:** `CrossL2ProverV2` **and** `CrossL2Executor` (both build under the
  `default` profile). Freeze and publish both.
- **Deploy (v1):** `CrossL2ProverV2` only. The executor artifact ships in the release but its
  deployment is deferred (TBD).
- **Out of scope (for now):** native_fallback contracts (`NativeProver`, `Registry`, OPStack
  provers). They need `via_ir` (the `native` profile) and there's a chance they're removed.
  TBD. If kept, they need their own parity verification (their deployed recipe is unconfirmed)
  before joining the release.

## Migration from infra

- The **deploy logic + release build move here**; infra stays the **orchestrator**: it holds
  secrets, picks chains, sets the env vars above, and invokes this repo's deploy action.
- The current `infra/contract-deploys` recompiles at deploy — it is replaced by the bytes-based
  deploy. Cut over one chain, confirm the address matches, then retire it.

## Decisions

- **Compiler recipe:** Option B — `bytecode_hash="none"` + `cbor_metadata=false` (metadata-free,
  matches ProverV2). Comments/formatting never move the address; only logic does. Adopting it is a
  one-time relaunch from the v1 `ipfs` addresses — prepped now, not executed.
- **Address source of truth:** infra `overlays/<env>/contracts-configmap-patch.yaml`
  (`PROVER_V2_ADDRESS` + per-chain overrides). No address registry in this repo; the deploy
  asserts against an `EXPECTED_ADDRESS` that infra supplies from that configmap.
- **Salt:** fixed per-env secret in the infra repo, read from env at deploy time.
- **Executor:** included in the release artifacts; deployment deferred (TBD).

## Still TBD

1. **Executor deployment** — when/whether to deploy it (its artifact is already frozen per release).
2. **Native_fallback** — keep or remove; if kept, parity-verify the deployed recipe and add the
   `native` profile to the release build.
3. **Non-canonical chains** (MEGAETH, REDBELLY, TRON) — policy for chains lacking the canonical
   factory: deploy the factory first to land the canonical address, or keep deliberate per-chain
   overrides in infra.
