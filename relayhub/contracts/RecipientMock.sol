// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.0 <0.9.0;

/**
 * @title RecipientMock
 * @dev Store & retreive value in a variable
 */
contract RecipientMock{

    uint256 number;

    constructor() {
    }

    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        number = num;

        emit ValueSeted(msg.sender,number);
    }

    /**
     * @dev Return value 
     * @return value of 'number'
     */
    function retreive() public view returns (uint256){
        return number;
    }

    event ValueSeted(address sender, uint256 newNumber);
}