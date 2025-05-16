pragma solidity 0.8.28;

import {IL1Block} from "../../../interfaces/IL1Block.sol";
import {IBuffer} from "block-hash-pusher/contracts/interfaces/IBuffer.sol";

contract NitroL1BlockWrapper is IL1Block {
    IBuffer private BLOCK_HASH_PUSHER;

    constructor(address blockHashPusher) {
        BLOCK_HASH_PUSHER = IBuffer(blockHashPusher);
    }

    function hash() external view returns (bytes32) {
        return BLOCK_HASH_PUSHER.parentChainBlockHash(BLOCK_HASH_PUSHER.newestBlockNumber());
    }
}
