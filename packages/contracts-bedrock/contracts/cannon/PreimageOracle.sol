// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

import { PreimageKeyLib } from "./PreimageKeyLib.sol";

/// @title PreimageOracle
/// @notice A contract for storing permissioned pre-images.
contract PreimageOracle {
    /// @notice Mapping of pre-image keys to pre-image lengths.
    mapping(bytes32 => uint256) public preimageLengths;
    /// @notice Mapping of pre-image keys to pre-image parts.
    mapping(bytes32 => mapping(uint256 => bytes32)) public preimageParts;
    /// @notice Mapping of pre-image keys to pre-image part offsets.
    mapping(bytes32 => mapping(uint256 => bool)) public preimagePartOk;

    /// @notice Reads a pre-image from the oracle.
    /// @param _key The key of the pre-image to read.
    /// @param _offset The offset of the pre-image to read.
    /// @return dat_ The pre-image data.
    /// @return datLen_ The length of the pre-image data.
    function readPreimage(bytes32 _key, uint256 _offset)
        external
        view
        returns (bytes32 dat_, uint256 datLen_)
    {
        require(preimagePartOk[_key][_offset], "pre-image must exist");

        // Calculate the length of the pre-image data
        // Add 8 for the length-prefix part
        datLen_ = 32;
        uint256 length = preimageLengths[_key];
        if (_offset + 32 >= length + 8) {
            datLen_ = length + 8 - _offset;
        }

        // Retrieve the pre-image data
        dat_ = preimageParts[_key][_offset];
    }

    /// TODO(CLI-4104):
    /// we need to mix-in the ID of the dispute for local-type keys to avoid collisions,
    /// and restrict local pre-image insertion to the dispute-managing contract.
    /// For now we permit anyone to write any pre-image unchecked, to make testing easy.
    /// This method is DANGEROUS. And NOT FOR PRODUCTION.
    function cheat(
        uint256 partOffset,
        bytes32 key,
        bytes32 part,
        uint256 size
    ) external {
        preimagePartOk[key][partOffset] = true;
        preimageParts[key][partOffset] = part;
        preimageLengths[key] = size;
    }

    /// @notice Loads a word of local data into the preimage oracle in two separate parts.
    /// @param _ident The identifier of the local data.
    /// @param _word The local data word.
    /// @param _size The number of bytes in `_word` to load.
    /// @dev The local data parts are loaded into the preimage oracle under the context
    ///      of the caller - no other account can write to the caller's context
    ///      specific data.
    ///
    ///      There are 5 local data identifiers:
    ///      ┌────────────┬────────────────────────┐
    ///      │ Identifier │      Data              │
    ///      ├────────────┼────────────────────────┤
    ///      │          1 │ L1 Head Hash (bytes32) │
    ///      │          2 │ Output Root (bytes32)  │
    ///      │          3 │ Root Claim (bytes32)   │
    ///      │          4 │ L2 Block Number (u64)  │
    ///      │          5 │ Chain ID (u64)         │
    ///      └────────────┴────────────────────────┘
    function loadLocalData(
        uint256 _ident,
        bytes32 _word,
        uint8 _size
    ) external returns (bytes32 key_) {
        // Compute the localized key from the given local identifier.
        key_ = PreimageKeyLib.localizeIdent(_ident);

        // Load both parts of the local data word into storage for future
        // reads.
        bytes32 part1;
        assembly {
            // The first part is prepended with an 8 byte length prefix and contains
            // the first 24 bytes of the passed word.
            part1 := or(shl(192, _size), shr(64, _word))
        }

        // Store the first part with offset 0.
        preimagePartOk[key_][0] = true;
        preimageParts[key_][0] = part1;

        // If the size is greater than 24, we need to store a second part as well.
        if (_size > 24) {
            bytes32 part2 = _word << 192;
            preimagePartOk[key_][32] = true;
            preimageParts[key_][32] = part2;
        }

        // Assign the length of the preimage at the localized key.
        preimageLengths[key_] = _size;
    }

    /// @notice Prepares a pre-image to be read by keccak256 key, starting at
    ///         the given offset and up to 32 bytes (clipped at pre-image length, if out of data).
    /// @param _partOffset The offset of the pre-image to read.
    /// @param _preimage The preimage data.
    function loadKeccak256PreimagePart(uint256 _partOffset, bytes calldata _preimage) external {
        uint256 size;
        bytes32 key;
        bytes32 part;
        assembly {
            // len(sig) + len(partOffset) + len(preimage offset) = 4 + 32 + 32 = 0x44
            size := calldataload(0x44)

            // revert if part offset > size+8 (i.e. parts must be within bounds)
            if gt(_partOffset, add(size, 8)) {
                revert(0, 0)
            }
            // we leave solidity slots 0x40 and 0x60 untouched,
            // and everything after as scratch-memory.
            let ptr := 0x80
            // put size as big-endian uint64 at start of pre-image
            mstore(ptr, shl(192, size))
            ptr := add(ptr, 8)
            // copy preimage payload into memory so we can hash and read it.
            calldatacopy(ptr, _preimage.offset, size)
            // Note that it includes the 8-byte big-endian uint64 length prefix.
            // this will be zero-padded at the end, since memory at end is clean.
            part := mload(add(sub(ptr, 8), _partOffset))
            let h := keccak256(ptr, size) // compute preimage keccak256 hash
            // mask out prefix byte, replace with type 2 byte
            key := or(and(h, not(shl(248, 0xFF))), shl(248, 2))
        }
        preimagePartOk[key][_partOffset] = true;
        preimageParts[key][_partOffset] = part;
        preimageLengths[key] = size;
    }
}
