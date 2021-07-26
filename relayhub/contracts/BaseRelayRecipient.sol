// SPDX-License-Identifier:MIT
pragma solidity >=0.8.0 <0.9.0;

/**
 * A base contract to be inherited by any contract that want to obtain the original msg.sender 
 * A subclass must use "_msgSender()" instead of "msg.sender"
 */
abstract contract BaseRelayRecipient{

    /*
     * Forwarder singleton we accept calls from
     */
    address internal trustedForwarder;

    /**
     * return the sender of this call.
     * if the call came through our Relay Hub, return the original sender.
     * should be used in the contract anywhere instead of msg.sender
     */
    function _msgSender() internal virtual returns (address sender) {
        bytes memory bytesSender;
        (,bytesSender) = trustedForwarder.call(abi.encodeWithSignature("getMsgSender()"));

        return abi.decode(bytesSender, (address));
    }
}