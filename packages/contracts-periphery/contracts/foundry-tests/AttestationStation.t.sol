//SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/* Testing utilities */
import { Test } from "forge-std/Test.sol";
import { TestERC20 } from "../testing/helpers/TestERC20.sol";
import { TestERC721 } from "../testing/helpers/TestERC721.sol";
import { AssetReceiver } from "../universal/AssetReceiver.sol";
import { AttestationStation } from "../universal/op-nft/AttestationStation.sol";

contract AssetReceiver_Initializer is Test {
    address alice_attestor = address(128);
    address bob = address(256);
    address sally = address(512);
    function _setUp() public {
        // Give alice and bob some ETH
        vm.deal(alice_attestor, 1 ether);

        vm.label(alice_attestor, "alice_attestor");
        vm.label(bob, "bob");
        vm.label(sally, "sally");
    }
}

contract AssetReceiverTest is AssetReceiver_Initializer {
    function setUp() public {
        super._setUp();
    }

    function test_attest() external {
        AttestationStation attestationStation = new AttestationStation();

        // alice is going to attest about bob
        vm.prank(alice_attestor);
        address creator = alice_attestor;
        AttestationStation.AttestationData memory attestationData = AttestationStation
            .AttestationData({
            about: bob,
            key: bytes32("test-key:string"),
            val: bytes("test-value")
        });

        // assert the attestation starts empty
        assertEq(
            attestationStation.attestations(address(this), address(this), "test"),
            ""
        );

        // make attestation
        attestationStation.attest(attestationData.about, attestationData.key, attestationData.val);

        // assert the attestation is there
        assertEq(
            attestationStation.readAttestation(
                creator,
                attestationData.about,
                attestationData.key,
            ),
            attestationData.val
        );

        bytes memory new_val = bytes32("new updated value");
        // make a new attestations to same about and key
        attestationData = AttestationStation.AttestationData({
            about: attestationData.about,
            key: attestationStation.key,
            val: new_val
        });

        attestationStation.attest(attestationData.about, attestationData.key, attestationData.val);

        // assert the attestation is updated
        assertEq(
            attestationStation.readAttestation(
                creator,
                attestationData.about,
                attestationData.key,
            ),
            attestationData.val
        );
    }
}
