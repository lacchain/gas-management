pragma solidity 0.5.9;


contract AccountRulesList {
    event AccountAdded(
        bool accountAdded,
        address accountAddress
    );

    event AccountRemoved(
        bool accountRemoved,
        address accountAddress
    );

    event TargetAdded(
        bool targetAdded,
        address accountAddress
    );

    event TargetRemoved(
        bool targetRemoved,
        address accountAddress
    );

    address[] public accountAllowList;
    address[] public targetAllowList;
    mapping (address => uint256) private accountIndexOf; //1 based indexing. 0 means non-existent
    mapping (address => uint256) private targetIndexOf; //1 based indexing. 0 means non-existent

    function sizeAccounts() internal view returns (uint256) {
        return accountAllowList.length;
    }

    function sizeTargets() internal view returns (uint256) {
        return targetAllowList.length;
    }

    function existsAccount(address _account) internal view returns (bool) {
        return accountIndexOf[_account] != 0;
    }

    function existsTarget(address _target) internal view returns (bool) {
        return targetIndexOf[_target] != 0;
    }

    function addNewAccount(address _account) internal returns (bool) {
        if (accountIndexOf[_account] == 0) {
            accountAllowList.push(_account);
            accountIndexOf[_account] = accountAllowList.length; 
            return true;
        }
        return false;
    }

    function addAllAccounts(address[] memory accounts) internal returns (bool) {
        bool allAdded = true;
        for (uint i = 0; i < accounts.length; i++) {
            bool added = addNewAccount(accounts[i]);
            emit AccountAdded(added, accounts[i]);
            allAdded = allAdded && added;
        }

        return allAdded;
    }

    function removeOldAccount(address _account) internal returns (bool) {
        uint256 index = accountIndexOf[_account];
        if (index > 0 && index <= accountAllowList.length) { //1-based indexing
            //move last address into index being vacated (unless we are dealing with last index)
            if (index != accountAllowList.length) {
                address lastAccount = accountAllowList[accountAllowList.length - 1];
                accountAllowList[index - 1] = lastAccount;
                accountIndexOf[lastAccount] = index;
            }

            //shrink array
            accountAllowList.length -= 1;
            accountIndexOf[_account] = 0;
            return true;
        }
        return false;
    }

    function addNewTarget(address _target) internal returns (bool) {
        if (targetIndexOf[_target] == 0) {
            targetAllowList.push(_target);
            targetIndexOf[_target] = targetAllowList.length; 
            return true;
        }
        return false;
    }

    function addAllTargets(address[] memory targets) internal returns (bool) {
        bool allAdded = true;
        for (uint i = 0; i < targets.length; i++) {
            bool added = addNewTarget(targets[i]);
            emit TargetAdded(added, targets[i]);
            allAdded = allAdded && added;
        }

        return allAdded;
    }

    function removeOldTarget(address _target) internal returns (bool) {
        uint256 index = targetIndexOf[_target];
        if (index > 0 && index <= targetAllowList.length) { //1-based indexing
            //move last address into index being vacated (unless we are dealing with last index)
            if (index != targetAllowList.length) {
                address lastAccount = targetAllowList[targetAllowList.length - 1];
                targetAllowList[index - 1] = lastAccount;
                targetIndexOf[lastAccount] = index;
            }

            //shrink array
            targetAllowList.length -= 1;
            targetIndexOf[_target] = 0;
            return true;
        }
        return false;
    }
}
