// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import {Test} from "forge-std/Test.sol";
import {Registry} from "../../contracts/core/native_fallback/L1/Registry.sol";

// Owner-scoped action space for the invariant fuzzer. Every call is made as the owner, so the
// fuzzer explores what an authorized admin can do — never what a grantee can voluntarily renounce.
contract RegistryOwnerHandler is Test {
    Registry public registry;
    address public owner;

    constructor(Registry _registry, address _owner) {
        registry = _registry;
        owner = _owner;
    }

    function grantRevocable(address grantee, uint256 chainID) public {
        vm.prank(owner);
        try registry.grantChainID(grantee, chainID) {} catch {}
    }

    function revokeRevocable(address grantee, uint256 chainID) public {
        vm.prank(owner);
        try registry.revokeChainID(grantee, chainID) {} catch {}
    }

    function grantIrrevocable(address grantee, uint256 chainID) public {
        vm.prank(owner);
        try registry.grantChainIDIrrevocable(grantee, chainID) {} catch {}
    }

    function tryRevokeRole(bytes32 role, address account) public {
        vm.prank(owner);
        try registry.revokeRole(role, account) {} catch {}
    }
}

contract RegistryInvariantTest is Test {
    Registry registry;
    RegistryOwnerHandler handler;

    address constant OWNER = address(0x1);
    address constant GRANTEE = address(0xBEEF);
    uint256 constant PROTECTED_CHAIN = 777;

    function setUp() public {
        Registry.InitialL2Configuration[] memory l2 = new Registry.InitialL2Configuration[](0);
        Registry.InitialL1Configuration[] memory l1 = new Registry.InitialL1Configuration[](0);

        // No prank needed: Registry sets its owner from the constructor arg, not msg.sender.
        registry = new Registry(OWNER, l2, l1);

        vm.prank(OWNER);
        registry.grantChainIDIrrevocable(GRANTEE, PROTECTED_CHAIN);

        handler = new RegistryOwnerHandler(registry, OWNER);
        targetContract(address(handler));
    }

    // Once granted, no owner action sequence can strip an irrevocable grant.
    function invariant_ownerCannotRemoveIrrevocableGrant() public view {
        assertTrue(registry.isIrrevocableGrantee(GRANTEE, PROTECTED_CHAIN));
    }
}
