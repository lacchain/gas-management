// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.0 <0.9.0;

/**
 * @dev Interface for `RelayHub`,
 *
 */
interface IRelayHub {

    enum ErrorCode {
        MaxBlockGasLimit,
        BadOriginalSender,
        BadNonce,
        NotEnoughGas,
        IsNotContract,
        EmptyCode,
        InvalidSignature,
        InvalidDestination,
        OK
    }

     struct TransactionData {
        address from;
        uint256 nonce;
        uint256 gasLimit;
        address to;
        bytes encodedFunction;
    }

    /**
     * @dev Relays a protected transaction.
     *
     * Parameters:
     *  - `signingData`: client's RLP transaction + nodeAddress + exìration, without signature
     *  - `v`: client's signature transaction v parameter
     *  - `r`: client's signature transaction r parameter
     *  - `s`: client's signature transaction s parameter
     */

    function relayMetaTx(
        uint256 gasLimit,
        bytes memory signingData,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns(ErrorCode);

    /**
     * @dev Relays a protected transaction to deploy a new contract.
     *
     * Parameters:
     *  - `signingData`: client's RLP transaction + nodeAddress + exìration, without signature
     *  - `v`: client's signature transaction v parameter
     *  - `r`: client's signature transaction r parameter
     *  - `s`: client's signature transaction s parameter
     */

    function deployMetaTx(
        uint256 gasLimit,
        bytes memory signingData,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns(ErrorCode, address);

    /**
     * @dev Returns an account's nonce in `RelayHub`.
     */
    function getNonce(address from) external view returns (uint256);

     /**
     * @dev Emitted when a transaction is relayed.
     */
    event TransactionRelayed(address indexed relay, address indexed from, address indexed to, bool executed, bytes output);
}