// SPDX-License-Identifier: MIT
// @unsupported: evm
pragma solidity >0.5.0 <0.8.0;

/* Library Imports */
import { Lib_ErrorUtils } from "../../libraries/utils/Lib_ErrorUtils.sol";

/**
 * @title OVM_ExecutionManagerWrapper
 *
 * Compiler used: optimistic-solc
 * Runtime target: OVM
 */
contract OVM_ExecutionManagerWrapper {
    fallback()
        external
    {
        bytes memory data = msg.data;
        assembly {
            // kall is a custom yul builtin within optimistic-solc that allows us to directly call
            // the execution manager (since `call` would be compiled).
            kall(add(data, 0x20), mload(data), 0x0, 0x0)
            let size := returndatasize()
            let returndata := mload(0x40)
            mstore(0x40, add(returndata, and(add(add(size, 0x20), 0x1f), not(0x1f))))
            mstore(returndata, size)
            returndatacopy(add(returndata, 0x20), 0x0, size)
            return(add(returndata, 0x20), mload(returndata))
        }
    }
}
