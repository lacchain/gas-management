// test/TxRelay.test.js
// Load dependencies
const { expectEvent } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');

// Load compiled artifacts
const TxRelay = artifacts.require('TxRelay');
const RecipientMock = artifacts.require('RecipientMock');


// Start test block
contract('Receipient', function (accounts) {
  const [ node, admin, permissioningContract] = accounts;

  beforeEach(async function () {
    // Deploy a new Box contract for each test
    this.txRelay = await TxRelay.new(10,permissioningContract, {from: admin});
    this.repicientMock = await RecipientMock.new();
  });

  // Test case
  it('deploy contract', async function () {
    const encodedFunction = '0xf861808082ea6094fd32cfc2e71611626d6368a41f915d0077a306a180b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0xc799efa03b5704c80758759aa0844544a5afba312f1d6f4a49b27dcbd75d14fa';
    const s = '0x0f1ab4f7308759e83db3820377fc595f10bdf103273a54fc061c33746a296cd0';

    const byteCode = '0xf8f08080833d09008080b8e6608060405234801561001057600080fd5b5060c78061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80636057361d146037578063b05784b8146062575b600080fd5b606060048036036020811015604b57600080fd5b8101908080359060200190929190505050607e565b005b60686088565b6040518082815260200191505060405180910390f35b8060008190555050565b6000805490509056fea26469706673582212208595760c8d4272f1711ffb94441dddc78e22630630efc2bc60b984de1caeb06c64736f6c63430006030033';
    const v1 = 28;
    const r1 = '0x3171f4477a944e826d8153e514e37517492f254daf5b365424f1b52f9508cbf5';
    const s1 = '0x66212041f82f45d2cb8bd57cdaf211b8ba5bfe037cd61fd5b3fc06b6b296febf';

    const receipt = await this.txRelay.addNode(node, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: node});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({from: node})).toString()).to.equal('0');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s,{from: node});

    expect(await this.txRelay.getGasLimit({from: node})).to.be.bignumber.gt('150000000');

    const receipt2 = await this.txRelay.deployMetaTx(byteCode,v1,r1,s1, {from: node});

    expectEvent(receipt2, 'Relayed', {from: node });

    expectEvent(receipt2, 'ContractDeployed', {from: "0x63949701cD0e1Cc04Dfea0AFBf410968F10fF4b6" });
  });

  // Test case
  it('forward to recipient', async function () {
    const encodedFunction = '0xf861808082ea6094fd32cfc2e71611626d6368a41f915d0077a306a180b8446057361d000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000173cf75f0905338597fcd38f5ce13e6840b230e9';
    const v = 27;
    const r = '0xc799efa03b5704c80758759aa0844544a5afba312f1d6f4a49b27dcbd75d14fa';
    const s = '0x0f1ab4f7308759e83db3820377fc595f10bdf103273a54fc061c33746a296cd0';

    const encodedFunction2 = '0xf8418080833877fb94'+this.repicientMock.address.substr(2)+'80a46057361d000000000000000000000000000000000000000000000000000000000000003c'
    const v2 = 27;
    const r2 = '0x69f6d236c76111f5026ff68ed37bef369598a75603657fa02f93a87bbddc62c6';
    const s2 = '0x3896e8f7353ec4e9a5bca4fb1be00449c49d2a015256a09e0dc7a376668016b5';

    const receipt = await this.txRelay.addNode(node, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: node});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({from: node})).toString()).to.equal('0');

    await this.txRelay.relayMetaTx(encodedFunction,v,r,s,{from: node});

    expect(await this.txRelay.getGasLimit({from: node})).to.be.bignumber.gt('150000000');

    const receipt2 = await this.txRelay.relayMetaTx(encodedFunction2,v,r,s, {from: node});

    expectEvent(receipt2, 'TransactionRelayed', {to: this.repicientMock.address, output:null});
    
  });

  // Test case
  it('forward to recipient and get revert', async function () {
    const encodedFunction = '0xe18080832dc6c094d8Be627F46F79EC31e47B9632e5B88f5b89491De80843c3da97d'
    const byteCodeRevert = '0xf901558080832dc6c08080b9014a608060405234801561001057600080fd5b5061012a806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80633c3da97d14602d575b600080fd5b6033604b565b60405180821515815260200191505060405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161460ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f6e616275636f646f73000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600190509056fea2646970667358221220deff2af94802840074e9ec6b4bc34a1dd49cd0abb7541476d8d4c50c13c864c764736f6c634300060c0033';
    const v = 27;
    const r = '0x08a9eca583b378ce767ba933ee1569df62aad5e3b40e315970628eb0d090f443';
    const s = '0x75c564c4a37644c7fc1633c48d56010ae23e2b2820e7cf91021fd8e91431863f';

    const receipt = await this.txRelay.addNode(node, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: node});
    
    expect((await this.txRelay.getNodes()).toString()).to.equal('1');
    
    expect((await this.txRelay.getGasLimit({from: node})).toString()).to.equal('0');

    const receipt2 = await this.txRelay.deployMetaTx(byteCodeRevert,v,r,s, {from: node});

    expectEvent(receipt2, 'Relayed', {from: node });

    expectEvent(receipt2, 'ContractDeployed', {from: "0x76C3049C227ccAED0c8DCAfF1170C11e0285Df23", contractDeployed:"0xd8Be627F46F79EC31e47B9632e5B88f5b89491De" });

    expect(await this.txRelay.getGasLimit({from: node})).to.be.bignumber.gt('150000000');

    const receipt3 = await this.txRelay.relayMetaTx(encodedFunction, v,r,s, {from: node});

    expectEvent(receipt3, 'TransactionRelayed', {to: "0xd8Be627F46F79EC31e47B9632e5B88f5b89491De", output:'0x08c379a0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000096e616275636f646f730000000000000000000000000000000000000000000000'});
    
  });

});