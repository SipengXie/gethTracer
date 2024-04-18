package vm

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Try to support inter-transaction concurrency
// now we could only considering the intra-transaction concurrency
type Metadata struct {
	TxId  int `json:"-" `    // used for inter-transaction concurrency
	Index int `json:"index"` // used for intra-transaction concurrency

	// Addr, Pc, OpCode are extra information for the concrete instruction
	Addr   common.Address `json:"-"`
	Pc     uint64         `json:"pc"`
	OpCode string         `json:"opcode"`
}

var SourceMeta = Metadata{
	Index: -1,
}

func NewMetadata(txId, index int, addr common.Address, pc uint64, opcode OpCode) *Metadata {
	return &Metadata{TxId: txId, Index: index, Addr: addr, Pc: pc, OpCode: opcode.String()}
}

var metaStackPool = sync.Pool{
	New: func() interface{} {
		return &MetaStack{data: make([]Metadata, 0, 16)}
	},
}

type MetaStack struct {
	data []Metadata
}

func newMetaStack() *MetaStack {
	return metaStackPool.Get().(*MetaStack)
}

func returnMetaStack(s *MetaStack) {
	s.data = s.data[:0]
	stackPool.Put(s)
}

// Data returns the underlying Metadata array.
func (st *MetaStack) Data() []Metadata {
	return st.data
}

func (st *MetaStack) push(d *Metadata) {
	// NOTE push limit (1024) is checked in baseCheck
	st.data = append(st.data, *d)
}

func (st *MetaStack) pop() (ret Metadata) {
	ret = st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return
}

func (st *MetaStack) len() int {
	return len(st.data)
}

func (st *MetaStack) swap(n int) {
	st.data[st.len()-n], st.data[st.len()-1] = st.data[st.len()-1], st.data[st.len()-n]
}

func (st *MetaStack) dup(n int) {
	st.push(&st.data[st.len()-n])
}

func (st *MetaStack) peek() *Metadata {
	return &st.data[st.len()-1]
}

// Back returns the n'th item in stack
func (st *MetaStack) Back(n int) *Metadata {
	return &st.data[st.len()-n-1]
}

type MetaMemory struct {
	store []Metadata
}

// NewMemory returns a new memory model.
func newMetaMemory() *MetaMemory {
	return &MetaMemory{}
}

// Set sets offset + size to value
func (m *MetaMemory) Set(offset, size uint64, value Metadata) {
	// It's possible the offset is greater than 0 and size equals 0. This is because
	// the calcMemSize (common.go) could potentially return 0 when size is zero (NO-OP)
	if size > 0 {
		// length of store may never be less than offset + size.
		// The store should be resized PRIOR to setting the memory
		if offset+size > uint64(len(m.store)) {
			panic("invalid memory: store empty")
		}
		for i := uint64(0); i < size; i++ {
			m.store[offset+i] = value
		}
	}
}

// Set32 sets the 32 bytes starting at offset to the value of val, left-padded with zeroes to
// 32 bytes.
func (m *MetaMemory) Set32(offset uint64, value Metadata) {
	// length of store may never be less than offset + size.
	// The store should be resized PRIOR to setting the memory
	if offset+32 > uint64(len(m.store)) {
		panic("invalid memory: store empty")
	}
	// Fill in relevant bits
	for i := uint64(0); i < 32; i++ {
		m.store[offset+i] = value
	}
}

// Resize resizes the memory to size
func (m *MetaMemory) Resize(size uint64) {
	if uint64(m.Len()) < size {
		m.store = append(m.store, make([]Metadata, size-uint64(m.Len()))...)
	}
}

// GetCopy returns offset + size as a new slice == GetPtr
func (m *MetaMemory) GetCopy(offset, size int64) Metadata {
	return m.GetPtr(offset, size)
}

// GetPtr returns the offset + size
func (m *MetaMemory) GetPtr(offset, size int64) Metadata {
	if size == 0 {
		return Metadata{}
	}

	if len(m.store) > int(offset) {
		return m.store[offset]
	}

	return Metadata{}
}

// Len returns the length of the backing slice
func (m *MetaMemory) Len() int {
	return len(m.store)
}

// Data returns the backing slice
func (m *MetaMemory) Data() []Metadata {
	return m.store
}

// Copy copies data from the src position slice into the dst position.
// The source and destination may overlap.
// OBS: This operation assumes that any necessary memory expansion has already been performed,
// and this method may panic otherwise.
func (m *MetaMemory) Copy(dst, src, len uint64) {
	if len == 0 {
		return
	}
	copy(m.store[dst:], m.store[src:src+len])
}

// this structure is the shadow storage implementation
// only serving for SLOAD and SSTORE
// could possibly facilitate inter-transaction concurrency
type EachStorage map[common.Hash]Metadata

type MetaStorage struct {
	store map[common.Address]EachStorage
}

func newMetaStorage() *MetaStorage {
	return &MetaStorage{store: make(map[common.Address]EachStorage)}
}

func (s *MetaStorage) Get(addr common.Address, key common.Hash) Metadata {
	if _, ok := s.store[addr]; !ok {
		return SourceMeta
	}
	return s.store[addr][key]
}

func (s *MetaStorage) Set(addr common.Address, key common.Hash, value Metadata) {
	if _, ok := s.store[addr]; !ok {
		s.store[addr] = make(EachStorage)
	}
	s.store[addr][key] = value
}

type MetaAccount struct {
	store map[common.Address]Metadata
}

func newMetaAccount() *MetaAccount {
	return &MetaAccount{store: make(map[common.Address]Metadata)}
}

func (s *MetaAccount) Get(addr common.Address) Metadata {
	if v, ok := s.store[addr]; !ok {
		return SourceMeta
	} else {
		return v
	}
}

func (s *MetaAccount) Set(addr common.Address, value Metadata) {
	s.store[addr] = value
}
