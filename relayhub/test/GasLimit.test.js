// test/GasLimit.test.js
// Load dependencies
const { BN, expectEvent, expectRevert } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');

// Load compiled artifacts
const GasLimit = artifacts.require('GasLimit');

// Start test block
contract('GasLimit', function (accounts) {
  const [ node, admin, permissioningContract, newAccountIngress, notAdmin, notPermissioningContract, node2, node3 ] = accounts;

  beforeEach(async function () {
    // Deploy a new gasLimit contract for each test
    this.gasLimit = await GasLimit.new(10,permissioningContract, {from: admin});
  });

  // Test case
  it('add new Node', async function () {
    const receipt = await this.gasLimit.addNode(node, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: node});
    
    expect((await this.gasLimit.getNodes()).toString()).to.equal('1');
  });

  // Test case
  it('add new Node from a not Permissioning account', async function () {
    await expectRevert(this.gasLimit.addNode(node, {from: notPermissioningContract}),"Caller is not Account Contract");
    
  });

  // Test case
  it('delete a Node', async function () {
    const receipt = await this.gasLimit.addNode(node, {from: permissioningContract});
    expectEvent(receipt, 'NodeAdded', { newNode: node});

    const receipt2 = await this.gasLimit.addNode(node2, {from: permissioningContract});
    expectEvent(receipt2, 'NodeAdded', { newNode: node2});

    const receipt3 = await this.gasLimit.addNode(node3, {from: permissioningContract});
    expectEvent(receipt3, 'NodeAdded', { newNode: node3});
    
    expect((await this.gasLimit.getNodes()).toString()).to.equal('3');

    const receipt4 = await this.gasLimit.deleteNode(node2, {from: permissioningContract});
    expectEvent(receipt4, 'NodeDeleted', { oldNode: node2});
    
    expect((await this.gasLimit.getNodes()).toString()).to.equal('2');

    const receipt5 = await this.gasLimit.deleteNode(node3, {from: permissioningContract});
    expectEvent(receipt5, 'NodeDeleted', { oldNode: node3});
    
    expect((await this.gasLimit.getNodes()).toString()).to.equal('1');
  });

  // Test case
  it('set gasUsedLastBlocks', async function () {
    const receipt = await this.gasLimit.setGasUsedLastBlocks(20000000, {from: admin});
    
    expect((await this.gasLimit.getGasUsedLastBlocks()).toString()).to.equal('20000000');
  });

  // Test case
  it('set new block frequency', async function () {
    const receipt = await this.gasLimit.setBlocksFrequency(5, {from: admin});
    
    expectEvent(receipt, 'BlockFrequencyChanged', { admin: admin, blocksFrequency: new BN(5)});
  });

  // Test case
  it('set max gas block Limit', async function () {
    const receipt = await this.gasLimit.setMaxGasBlockLimit(1000000, {from: admin});
    
    expectEvent(receipt, 'MaxGasBlockLimitChanged', { admin: admin, maxGasBlockLimit: new BN(1000000)});
  });

  // Test case
  it('set Gas used by RelayHub', async function () {
    const receipt = await this.gasLimit.setGasUsedRelayHub(100000, {from: admin});
    
    expectEvent(receipt, 'GasUsedRelayHubChanged', { admin: admin, gasUsedRelayHub: new BN(100000)});
  });

  // Test case
  it('set account ingress', async function () {
    const receipt = await this.gasLimit.setAccounIngress(newAccountIngress, {from: admin});
    
    expectEvent(receipt, 'AccountIngressChanged', { admin: admin, newAddress: newAccountIngress});
  });

  // Test case
  it('set account ingress from a not Admin account', async function () {
    await expectRevert(this.gasLimit.setAccounIngress(newAccountIngress, {from: notAdmin}),"Caller is not Admin");
  });

});