// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Semver } from "@eth-optimism/contracts-bedrock/contracts/universal/Semver.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract AttestationStation is Initializable, Semver {
    struct AttestationData {
        address about;
        bytes32 key;
        bytes val;
    }

    mapping(address => mapping(address => mapping(bytes32 => bytes))) public attestations;

    event AttestationCreated(
        address indexed creator,
        address indexed about,
        bytes32 indexed key,
        bytes val
    );

    constructor() Semver(0, 0, 1) {}

    /**
     * @notice  Initialize the AttestationStation contract.
     */
    function initialize() public initializer {}

    /**
     * @notice  Attest to the given data.
     * @dev     Attests to the given data from the sender.
     * @param   _attestations  The array of attestation data.
     */
    function attestBulk(AttestationData[] memory _attestations) public {
        for (uint256 i = 0; i < _attestations.length; ++i) {
            attest(_attestations[i]);
        }
    }

    /**
     * @notice  Attest to the given data.
     * @dev     Attests to the given data from the sender.
     * @param   _attestation  The attestation data.
     */
    function attest(AttestationData memory _attestation) public {
        attestations[msg.sender][_attestation.about][_attestation.key] = _attestation.val;
        emit AttestationCreated(
            msg.sender,
            _attestation.about,
            _attestation.key,
            _attestation.val
        );
    }

    /**
     * @notice  Reads an attestation
     * @return  bytes  The attestation
     */
    function readAttestation(address memory creator, address memory about,  bytes32 memory key) public view returns (bytes memory val) {
        val = attestationStation.attestations(admin, about, key);
        return val;
    }

    /**
     * @notice Reads an attestation string
     * @dev This is very convenient for people using etherscan
     * @dev Takes in a string key instead of a bytes32
     * @return The attestation formatted as a string
     */
    function readAttestationString(address memory creator, address memory about,  string memory keyStr) public view returns (string memory valStr) {
        val = attestationStation.attestations(
            admin,
            about,
            bytes32(keyStr)
        );
        valStr = string(abi.encodePacked((val)));
        return valStr;
    }
}
