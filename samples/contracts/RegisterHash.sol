pragma solidity >=0.4.22 <0.7.0;

import "./BaseRelayRecipient.sol";

/**
 * @title RegisterHash
 * @dev Store & retreive value in a variable
 */
contract RegisterHash is BaseRelayRecipient{

    mapping (address => bytes32) records;
    address owner;

    constructor() public {
        owner = _msgSender();
    }

    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(bytes32 memory hash) public {
        address memory _msgSender = _msgSender();
        records[_msgSender] = hash;

        emit HashSeted(_msgSender,hash);
    }

    /**
     * @dev Return value 
     * @return value of 'number'
     */
    function retreiveHash(address sender) public view returns (bytes32){
        return records[sender];
    }

    function getOwner() public view returns (address){
        return owner;
    }

    event HashSeted(address sender, bytes32 hash);
}
