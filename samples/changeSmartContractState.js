const {web3} = require("./pantheon_utils/web3")
// set to 1 for faster validation in this course.
web3.transactionConfirmationBlocks = 1
const sha3 = require("js-sha3").keccak_256

let set

const setRecipientContract = () => { //Set for RelayHub Contract
  console.log("#######################Preparing value for Recipient SmartContract #######################")
  //function store(uint256)
  const functionName = "store"
  const typeOfData = "uint256"

  const number = 10
  const nodeAddress = "0xfaae4e8e9dabf9859db1601024191f3c97302230"
  const expiration = 1636394529

  let set = web3.eth.abi.encodeFunctionSignature(`${functionName}(${typeOfData})`)//function name to use
  
  let value = web3.eth.abi.encodeParameters(
    ["uint256","address","uint256"],
    [number,nodeAddress,expiration])//setting the value
  
  const txData = set + value.substr(2)
  console.log("txData:"+txData)
  return txData
}

const chooseSmartContractSetter = () => {
    set = setRecipientContract
}

chooseSmartContractSetter()

module.exports = {set}
