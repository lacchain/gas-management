pragma solidity 0.5.9;

import "./AccountRulesProxy.sol";
import "./AccountRulesList.sol";
import "./AccountIngress.sol";
import "./Admin.sol";


contract AccountRules is AccountRulesProxy, AccountRulesList {

    address constant public ON_CHAIN_PRIVACY_ADDRESS = 0x000000000000000000000000000000000000007E;
    // in read-only mode rules can't be added/removed
    // this will be used to protect data when upgrading contracts
    bool private readOnlyMode = false;
    // version of this contract: semver like 1.2.14 represented like 001002014
    uint private version = 1000000;

    AccountIngress private ingressContract;

    address private relayHub;

    modifier onlyOnEditMode() {
        require(!readOnlyMode, "In read only mode: rules cannot be modified");
        _;
    }

    modifier onlyAdmin() {
        address adminContractAddress = ingressContract.getContractAddress(ingressContract.ADMIN_CONTRACT());

        require(adminContractAddress != address(0), "Ingress contract must have Admin contract registered");
        require(Admin(adminContractAddress).isAuthorized(msg.sender), "Sender not authorized");
        _;
    }

    constructor (AccountIngress _ingressContract) public {
        ingressContract = _ingressContract;
        addNewAccount(msg.sender);
    }

    // VERSION
    function getContractVersion() public view returns (uint) {
        return version;
    }

    // READ ONLY MODE
    function isReadOnly() public view returns (bool) {
        return readOnlyMode;
    }

    function enterReadOnly() public onlyAdmin returns (bool) {
        require(readOnlyMode == false, "Already in read only mode");
        readOnlyMode = true;
        return true;
    }

    function exitReadOnly() public onlyAdmin returns (bool) {
        require(readOnlyMode == true, "Not in read only mode");
        readOnlyMode = false;
        return true;
    }

    function transactionAllowed(
        address sender,
        address target,
        uint256, // value
        uint256, // gasPrice
        uint256, // gasLimit
        bytes memory // payload
    ) public view returns (bool) {
        if (accountPermitted(sender) && destinationPermitted(target)) {
            //decreaseGasLimit(sender, target);
            return true;
        } else {
            return false;
        }
    }

    function accountPermitted(
        address _account
    ) public view returns (bool) {
        return existsAccount(_account);
    }

    function destinationPermitted(
        address _target
    ) public view returns (bool) {
        return existsTarget(_target);
    }

    function addAccount(
        address account
    ) public onlyAdmin onlyOnEditMode returns (bool) {
        bool added = addNewAccount(account);
        emit AccountAdded(added, account);
        //call add gasLimit
        if (added){
            bytes memory payload = abi.encodeWithSignature("addNode(address)",account);
            (bool cResponse, bytes memory result) = relayHub.call(payload);
            added = cResponse;
            require (cResponse, "Node haven't been added to GasLimit");
        }
        return added;
    }

    function removeAccount(
        address account
    ) public onlyAdmin onlyOnEditMode returns (bool) {
        bool removed = removeOldAccount(account);
        emit AccountRemoved(removed, account);
        return removed;
    }

    function getSize() public view returns (uint) {
        return sizeAccounts();
    }

    function getByIndex(uint index) public view returns (address account) {
        return accountAllowList[index];
    }

    function getAccounts() public view returns (address[] memory){
        return accountAllowList;
    }

    function getTargets() public view returns (address[] memory){
        return targetAllowList; 
    }

    function addAccounts(address[] memory accounts) public onlyAdmin returns (bool) {
        return addAllAccounts(accounts);
    }

    function addTarget(
        address target
    ) public onlyAdmin onlyOnEditMode returns (bool) {
        bool added = addNewTarget(target);
        emit TargetAdded(added, target);
        return added;
    }

    function removeTarget(
        address target
    ) public onlyAdmin onlyOnEditMode returns (bool) {
        bool removed = removeOldTarget(target);
        emit TargetRemoved(removed, target);
        return removed;
    }

    function addTargets(address[] memory targets) public onlyAdmin returns (bool) {
        return addAllTargets(targets);
    }

    function setRelay(address _relayHub) public onlyAdmin returns (bool) {
        relayHub = _relayHub;
    }

    function decreaseGasLimit(address _sender, address _target) private returns (bool){
        emit AccountVerified(true, _sender);
        uint256 gasUsed = 300000;
        if (_target == ON_CHAIN_PRIVACY_ADDRESS){
            gasUsed = 25000;    
        } 
        
        bytes memory payload = abi.encodeWithSignature("decreaseGasUsed(address,uint256)",_sender,gasUsed);
        (bool cResponse, bytes memory result) = relayHub.call(payload);
        return cResponse;
    }

    event AccountVerified(
        bool accountAdded,
        address accountAddress
    );
}
