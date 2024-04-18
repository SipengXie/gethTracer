package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type fakeAccountData struct {
	Nonce    uint64      `json:"nonce,omitempty"`
	Balance  *big.Int    `json:"balance,omitempty"`
	Root     common.Hash `json:"root,omitempty"` // MPT root of the storage trie
	CodeHash common.Hash `json:"code_hash,omitempty"`
}

type fakeAccountObject struct {
	Address      common.Address              `json:"address,omitempty"`
	ByteCode     []byte                      `json:"byte_code,omitempty"`
	Data         fakeAccountData             `json:"data,omitempty"`
	CacheStorage map[common.Hash]common.Hash `json:"cache_storage,omitempty"` // 用于缓存存储的变量
	IsAlive      bool                        `json:"is_alive,omitempty"`
}

func newFakeAccountObject(address common.Address, data fakeAccountData) *fakeAccountObject {
	if data.Balance == nil {
		data.Balance = new(big.Int)
	}
	if (data.CodeHash == common.Hash{}) {
		data.CodeHash = types.EmptyCodeHash
	}
	return &fakeAccountObject{
		Address:      address,
		Data:         data,
		CacheStorage: make(map[common.Hash]common.Hash),
		IsAlive:      true,
	}
}

func (object *fakeAccountObject) GetBalance() *big.Int {
	return object.Data.Balance
}

func (object *fakeAccountObject) SubBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	object.Data.Balance = new(big.Int).Sub(object.Data.Balance, amount)
}

func (object *fakeAccountObject) AddBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	object.Data.Balance = new(big.Int).Add(object.Data.Balance, amount)
}

func (object *fakeAccountObject) SetBalance(amount *big.Int) {
	object.Data.Balance = amount
}

func (object *fakeAccountObject) GetNonce() uint64 {
	return object.Data.Nonce
}

func (object *fakeAccountObject) SetNonce(nonce uint64) {
	object.Data.Nonce = nonce
}

func (object *fakeAccountObject) CodeHash() common.Hash {
	return object.Data.CodeHash
}

func (object *fakeAccountObject) Code() []byte {
	return object.ByteCode
}

func (object *fakeAccountObject) SetCode(codeHash common.Hash, code []byte) {
	object.Data.CodeHash = codeHash
	object.ByteCode = code
}

func (object *fakeAccountObject) GetStorageState(key common.Hash) (common.Hash, bool) {
	value, exist := object.CacheStorage[key]
	if exist {
		// fmt.Println("exist cache ", " key: ", key, " value: ", value)
		return value, true
	}
	return common.Hash{}, false
}

func (object *fakeAccountObject) SetStorageState(key, value common.Hash) {
	object.CacheStorage[key] = value
}

func (object *fakeAccountObject) Empty() bool {
	return object.Data.Nonce == 0 && object.Data.Balance.Sign() == 0 && (object.Data.CodeHash == types.EmptyCodeHash)
}
