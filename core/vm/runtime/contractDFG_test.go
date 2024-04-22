package runtime

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
)

func TestExecution(t *testing.T) {
	state := state.NewFakeState()

	deployCode := common.Hex2Bytes("608060405234801561001057600080fd5b506108a3806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633a9bbfcd146100515780634c803feb146100815780636b83dd2e146100b1578063b5463014146100e1575b600080fd5b61006b60048036038101906100669190610614565b610111565b6040516100789190610675565b60405180910390f35b61009b60048036038101906100969190610614565b610335565b6040516100a89190610675565b60405180910390f35b6100cb60048036038101906100c69190610614565b610496565b6040516100d89190610675565b60405180910390f35b6100fb60048036038101906100f69190610614565b6104f4565b6040516101089190610675565b60405180910390f35b6000806001836101219190610690565b67ffffffffffffffff811115610160577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561018e5781602001602082028036833780820191505090505b50905060005b8381116102eb57600181116101ee57808282815181106101dd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181815250506102d8565b816002826101fc9190610771565b81518110610233577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151826001836102499190610771565b81518110610280577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516102929190610690565b8282815181106102cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181815250505b80806102e3906107af565b915050610194565b50808381518110610325577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151915050919050565b6000808214156103485760009050610491565b600182141561035a5760019050610491565b3073ffffffffffffffffffffffffffffffffffffffff16634c803feb6002846103839190610771565b6040518263ffffffff1660e01b815260040161039f9190610675565b60206040518083038186803b1580156103b757600080fd5b505afa1580156103cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ef919061063d565b3073ffffffffffffffffffffffffffffffffffffffff16634c803feb6001856104189190610771565b6040518263ffffffff1660e01b81526004016104349190610675565b60206040518083038186803b15801561044c57600080fd5b505afa158015610460573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610484919061063d565b61048e9190610690565b90505b919050565b6000808214156104a957600090506104ef565b600060019050600191506000600290505b838110156104ec57600083836104d09190610690565b90508392508093505080806104e4906107af565b9150506104ba565b50505b919050565b60008082141561050757600090506105e5565b600060028361051691906106e6565b90506000600190505b81811161053257600181901b905061051f565b600181901c90506001925060006001905060005b60008311156105e057818261055b9190610717565b85866105679190610717565b6105719190610690565b9050600083871611156105ab5784600261058b9190610717565b826105969190610690565b826105a19190610717565b91508094506105d4565b848260026105b99190610717565b6105c39190610771565b856105ce9190610717565b94508091505b600183901c9250610546565b505050505b919050565b6000813590506105f981610856565b92915050565b60008151905061060e81610856565b92915050565b60006020828403121561062657600080fd5b6000610634848285016105ea565b91505092915050565b60006020828403121561064f57600080fd5b600061065d848285016105ff565b91505092915050565b61066f816107a5565b82525050565b600060208201905061068a6000830184610666565b92915050565b600061069b826107a5565b91506106a6836107a5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156106db576106da6107f8565b5b828201905092915050565b60006106f1826107a5565b91506106fc836107a5565b92508261070c5761070b610827565b5b828204905092915050565b6000610722826107a5565b915061072d836107a5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610766576107656107f8565b5b828202905092915050565b600061077c826107a5565b9150610787836107a5565b92508282101561079a576107996107f8565b5b828203905092915050565b6000819050919050565b60006107ba826107a5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156107ed576107ec6107f8565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b61085f816107a5565b811461086a57600080fd5b5056fea26469706673582212205aa624f01aeacae044ff9989fb2c19d7c1b42a8c4a0a0c427dbdb95f6e696b1764736f6c63430008040033")
	user := common.BytesToAddress([]byte("user"))
	state.CreateAccount(user)
	state.SetBalance(user, big.NewInt(1000000000000000000))

	cfg := new(Config)
	setDefaults(cfg)
	cfg.FakeState = state
	cfg.Origin = user
	evm := NewEnv(cfg)
	userRef := vm.AccountRef(user)

	_, addr, _, err := evm.Create(userRef, deployCode, cfg.GasLimit, big.NewInt(0), -1)
	if err != nil {
		t.Fatalf("Failed to deploy contract: %v", err)
	}

	// t.Log("Contract deployed at address", addr.Hex())
	// t.Log("Contract code", common.Bytes2Hex(code))
	// graph := evm.Graph
	// vm.VisualizeGraph(graph)

	// fib1input := common.Hex2Bytes("4c803feb0000000000000000000000000000000000000000000000000000000000000003")
	// fib2input := common.Hex2Bytes("3a9bbfcd0000000000000000000000000000000000000000000000000000000000000003")
	// fib3input := common.Hex2Bytes("6b83dd2e0000000000000000000000000000000000000000000000000000000000000003")
	fib4input := common.Hex2Bytes("b54630140000000000000000000000000000000000000000000000000000000000000003")

	evm = NewEnv(cfg)
	res, _, err := evm.Call(userRef, addr, fib4input, cfg.GasLimit, big.NewInt(0), -1)
	if err != nil {
		t.Fatalf("Failed to call contract: %v", err)
	}
	t.Log("Contract returned", res)

	graph := evm.Graph
	vm.VisualizeGraph(graph)
}