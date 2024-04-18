package vm

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/panjf2000/ants"
)

// Addr = []byte{0x19}
// input is uint256 size
func (evm *EVM) newPool(input []byte) error {
	size := new(big.Int).SetBytes(getData(input, 0, 32))
	pool, err := ants.NewPool(int(size.Int64()))
	if err != nil {
		return err
	}
	evm.pool = pool
	evm.wg = new(sync.WaitGroup)
	return nil
}

// Addr = []byte{0x20}
// input is uint256 how we call a function, let's have a test
func (evm *EVM) submitInterpreterRun(caller ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
	evm.wg.Add(1)
	ret := []byte{}
	err := error(nil)
	err = evm.pool.Submit(func() {
		// here we require stateDB to be concurrency-friendly
		code := evm.StateDB.GetCode(addr)
		addrCopy := addr
		contract := NewContract(caller, AccountRef(addrCopy), value, gas, -1)
		contract.SetCallCode(&addrCopy, evm.StateDB.GetCodeHash(addrCopy), code)
		ret, err = evm.interpreter.Run(contract, input, false) // may share a memory, which means the interpreter may be modified...
		gas = contract.Gas
		evm.wg.Done()
	})
	return ret, gas, err
}

// Addr = []byte{0x21}
// just waiting
func (evm *EVM) waitWg() {
	evm.wg.Wait()
}
