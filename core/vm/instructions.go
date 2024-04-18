// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

func opAdd(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaY := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, ADD)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Add(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaY)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaY)
	return nil, nil
}

func opSub(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaY := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SUB)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Sub(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaY)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaY)
	return nil, nil
}

func opMul(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaY := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, MUL)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Mul(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaY)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaY)
	return nil, nil
}

func opDiv(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaY := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, DIV)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Div(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaY)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaY)
	return nil, nil
}

func opSdiv(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaY := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SDIV)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.SDiv(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaY)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaY)
	return nil, nil
}

func opMod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaY := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, MOD)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Mod(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaY)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaY)
	return nil, nil
}

func opSmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaY := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SMOD)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.SMod(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaY)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaY)
	return nil, nil
}

func opExp(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, EXP)

	base, exponent := scope.Stack.pop(), scope.Stack.peek()
	exponent.Exp(&base, exponent)

	metaBase := scope.metaStack.pop()
	metaExponent := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaBase, metaExponent}, *newMetaRes)
	return nil, nil
}

func opSignExtend(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SIGNEXTEND)

	back, num := scope.Stack.pop(), scope.Stack.peek()
	num.ExtendSign(num, &back)

	metaBack := scope.metaStack.pop()
	metaNum := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaBack, metaNum}, *newMetaRes)
	return nil, nil
}

func opNot(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, NOT)

	x := scope.Stack.peek()
	x.Not(x)

	metaX := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX}, *newMetaRes)
	return nil, nil
}

func opLt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, LT)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Lt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opGt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, GT)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Gt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opSlt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SLT)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Slt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opSgt(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SGT)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Sgt(y) {
		y.SetOne()
	} else {
		y.Clear()
	}

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opEq(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, EQ)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	if x.Eq(y) {
		y.SetOne()
	} else {
		y.Clear()
	}

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opIszero(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, ISZERO)

	x := scope.Stack.peek()
	if x.IsZero() {
		x.SetOne()
	} else {
		x.Clear()
	}

	metaX := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX}, *newMetaRes)
	return nil, nil
}

func opAnd(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, AND)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.And(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opOr(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, OR)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Or(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opXor(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, XOR)

	x, y := scope.Stack.pop(), scope.Stack.peek()
	y.Xor(&x, y)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY}, *newMetaRes)
	return nil, nil
}

func opByte(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, BYTE)

	th, val := scope.Stack.pop(), scope.Stack.peek()
	val.Byte(&th)

	metaTh := scope.metaStack.pop()
	metaVal := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaTh, metaVal}, *newMetaRes)
	return nil, nil
}

func opAddmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, ADDMOD)

	x, y, z := scope.Stack.pop(), scope.Stack.pop(), scope.Stack.peek()
	if z.IsZero() {
		z.Clear()
	} else {
		z.AddMod(&x, &y, z)
	}

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	metaZ := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY, metaZ}, *newMetaRes)
	return nil, nil
}

func opMulmod(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, MULMOD)

	x, y, z := scope.Stack.pop(), scope.Stack.pop(), scope.Stack.peek()
	z.MulMod(&x, &y, z)

	metaX := scope.metaStack.pop()
	metaY := scope.metaStack.pop()
	metaZ := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaX, metaY, metaZ}, *newMetaRes)
	return nil, nil
}

// opSHL implements Shift Left
// The SHL instruction (shift left) pops 2 values from the stack, first arg1 and then arg2,
// and pushes on the stack arg2 shifted to the left by arg1 number of bits.
func opSHL(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// Note, second operand is left in the stack; accumulate result into it, and no need to push it afterwards
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SHL)

	shift, value := scope.Stack.pop(), scope.Stack.peek()
	if shift.LtUint64(256) {
		value.Lsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}

	metaShift := scope.metaStack.pop()
	metaValue := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaShift, metaValue}, *newMetaRes)
	return nil, nil
}

// opSHR implements Logical Shift Right
// The SHR instruction (logical shift right) pops 2 values from the stack, first arg1 and then arg2,
// and pushes on the stack arg2 shifted to the right by arg1 number of bits with zero fill.
func opSHR(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// Note, second operand is left in the stack; accumulate result into it, and no need to push it afterwards
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SHR)

	shift, value := scope.Stack.pop(), scope.Stack.peek()
	if shift.LtUint64(256) {
		value.Rsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}

	metaShift := scope.metaStack.pop()
	metaValue := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaShift, metaValue}, *newMetaRes)
	return nil, nil
}

// opSAR implements Arithmetic Shift Right
// The SAR instruction (arithmetic shift right) pops 2 values from the stack, first arg1 and then arg2,
// and pushes on the stack arg2 shifted to the right by arg1 number of bits with sign extension.
func opSAR(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SAR)

	shift, value := scope.Stack.pop(), scope.Stack.peek()
	if shift.GtUint64(256) {
		if value.Sign() >= 0 {
			value.Clear()
		} else {
			// Max negative shift: all bits set
			value.SetAllOne()
		}

		metaShift := scope.metaStack.pop()
		metaValue := scope.metaStack.pop()
		scope.metaStack.push(newMetaRes)

		interpreter.evm.Graph.AddDependency([]Metadata{metaShift, metaValue}, *newMetaRes)
		return nil, nil
	}
	n := uint(shift.Uint64())
	value.SRsh(value, n)

	metaShift := scope.metaStack.pop()
	metaValue := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaShift, metaValue}, *newMetaRes)
	return nil, nil
}

func opKeccak256(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, KECCAK256)

	offset, size := scope.Stack.pop(), scope.Stack.peek()
	data := scope.Memory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	if interpreter.hasher == nil {
		interpreter.hasher = crypto.NewKeccakState()
	} else {
		interpreter.hasher.Reset()
	}
	interpreter.hasher.Write(data)
	interpreter.hasher.Read(interpreter.hasherBuf[:])
	evm := interpreter.evm
	if evm.Config.EnablePreimageRecording {
		evm.StateDB.AddPreimage(interpreter.hasherBuf, data)
	}
	size.SetBytes(interpreter.hasherBuf[:])

	metaOffset := scope.metaStack.pop()
	metaSize := scope.metaStack.pop()
	metaData := scope.metaMemory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaOffset, metaSize, metaData}, *newMetaRes)
	return nil, nil
}

func opOrigin(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, ORIGIN)

	scope.Stack.push(new(uint256.Int).SetBytes(interpreter.evm.Origin.Bytes()))

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

func opBalance(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, BALANCE)

	slot := scope.Stack.peek()
	address := common.Address(slot.Bytes20())
	balance := interpreter.evm.StateDB.GetBalance(address)
	slot.SetFromBig(balance)

	metaSlot := scope.metaStack.pop()
	metaBalance := scope.metaBalance.Get(address)
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaSlot, metaBalance}, *newMetaRes)
	return nil, nil
}

// 只要从scope.Contract里面拿数据，都要对sourceIndex产生依赖
func opAddress(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, ADDRESS)

	scope.Stack.push(new(uint256.Int).SetBytes(scope.Contract.Address().Bytes()))

	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, nil
}

// 只要从scope.Contract里面拿数据，都要对sourceIndex产生依赖
func opCaller(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALLER)

	scope.Stack.push(new(uint256.Int).SetBytes(scope.Contract.Caller().Bytes()))

	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, nil
}

// 只要从scope.Contract里面拿数据，都要对sourceIndex产生依赖
func opCallValue(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALLVALUE)

	v, _ := uint256.FromBig(scope.Contract.value)
	scope.Stack.push(v)

	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, nil
}

// 只要从scope.Contract里面拿数据，都要对sourceIndex产生依赖
func opCallDataLoad(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALLDATALOAD)

	// 从栈顶拿一个64位数据，然后作为offset去取数据，数据长度32字节
	x := scope.Stack.peek()
	// 这个get data可能是一个opcode产生的，因为有合约调合约的存在
	if offset, overflow := x.Uint64WithOverflow(); !overflow {
		data := getData(scope.Contract.Input, offset, 32)
		x.SetBytes(data)
	} else {
		// 否则相当于压进去一个0
		x.Clear()
	}

	metaX := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{metaX, source}, *newMetaRes)
	return nil, nil
}

// 只要从scope.Contract里面拿数据，都要对sourceIndex产生依赖
func opCallDataSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALLDATASIZE)

	scope.Stack.push(new(uint256.Int).SetUint64(uint64(len(scope.Contract.Input))))

	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, nil
}

func opCallDataCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALLDATACOPY)

	var (
		memOffset  = scope.Stack.pop()
		dataOffset = scope.Stack.pop()
		length     = scope.Stack.pop()
	)
	dataOffset64, overflow := dataOffset.Uint64WithOverflow()
	if overflow {
		dataOffset64 = 0xffffffffffffffff
	}
	// These values are checked for overflow during gas cost calculation
	memOffset64 := memOffset.Uint64()
	length64 := length.Uint64()
	scope.Memory.Set(memOffset64, length64, getData(scope.Contract.Input, dataOffset64, length64))

	metaMemOffset := scope.metaStack.pop()
	metaDataOffset := scope.metaStack.pop()
	metaLength := scope.metaStack.pop()
	scope.metaMemory.Set(memOffset64, length64, *newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{metaMemOffset, metaDataOffset, metaLength, source}, *newMetaRes)
	return nil, nil
}

// 从interpreter.returnData里拿数据，也需要依赖它的SourceIndex
func opReturnDataSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, RETURNDATASIZE)

	scope.Stack.push(new(uint256.Int).SetUint64(uint64(len(interpreter.returnData))))

	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[interpreter.sourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, nil
}

func opReturnDataCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALLDATACOPY)

	var (
		memOffset  = scope.Stack.pop()
		dataOffset = scope.Stack.pop()
		length     = scope.Stack.pop()
	)

	offset64, overflow := dataOffset.Uint64WithOverflow()
	if overflow {
		return nil, ErrReturnDataOutOfBounds
	}
	// we can reuse dataOffset now (aliasing it for clarity)
	var end = dataOffset
	end.Add(&dataOffset, &length)
	end64, overflow := end.Uint64WithOverflow()
	if overflow || uint64(len(interpreter.returnData)) < end64 {
		return nil, ErrReturnDataOutOfBounds
	}
	scope.Memory.Set(memOffset.Uint64(), length.Uint64(), interpreter.returnData[offset64:end64])

	metaMemOffset := scope.metaStack.pop()
	metaDataOffset := scope.metaStack.pop()
	metaLength := scope.metaStack.pop()
	scope.metaMemory.Set(memOffset.Uint64(), length.Uint64(), *newMetaRes)

	source := interpreter.evm.Graph.Vertexes[interpreter.sourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{metaMemOffset, metaDataOffset, metaLength, source}, *newMetaRes)
	return nil, nil
}

func opExtCodeSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, EXTCODESIZE)

	slot := scope.Stack.peek()
	codeSize := uint64(interpreter.evm.StateDB.GetCodeSize(slot.Bytes20()))
	slot.SetUint64(codeSize)

	metaSlot := scope.metaStack.pop()
	metaCode := scope.metaCode.Get(slot.Bytes20())
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaSlot, metaCode}, *newMetaRes)
	return nil, nil
}

// 从scope contract拿数据了
func opCodeSize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CODESIZE)

	l := new(uint256.Int)
	l.SetUint64(uint64(len(scope.Contract.Code)))
	scope.Stack.push(l)

	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, nil
}

// 从scope contract拿数据了
func opCodeCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CODECOPY)

	var (
		memOffset  = scope.Stack.pop()
		codeOffset = scope.Stack.pop()
		length     = scope.Stack.pop()
	)
	uint64CodeOffset, overflow := codeOffset.Uint64WithOverflow()
	if overflow {
		uint64CodeOffset = 0xffffffffffffffff
	}
	codeCopy := getData(scope.Contract.Code, uint64CodeOffset, length.Uint64())
	scope.Memory.Set(memOffset.Uint64(), length.Uint64(), codeCopy)

	metaMemOffset := scope.metaStack.pop()
	metaDataOffset := scope.metaStack.pop()
	metaLength := scope.metaStack.pop()
	scope.metaMemory.Set(memOffset.Uint64(), length.Uint64(), *newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{metaMemOffset, metaDataOffset, metaLength, source}, *newMetaRes)
	return nil, nil
}

// 从statedb拿数据了
func opExtCodeCopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, EXTCODECOPY)

	var (
		stack      = scope.Stack
		a          = stack.pop()
		memOffset  = stack.pop()
		codeOffset = stack.pop()
		length     = stack.pop()
	)
	uint64CodeOffset, overflow := codeOffset.Uint64WithOverflow()
	if overflow {
		uint64CodeOffset = 0xffffffffffffffff
	}
	addr := common.Address(a.Bytes20())
	code := interpreter.evm.StateDB.GetCode(addr)
	codeCopy := getData(code, uint64CodeOffset, length.Uint64())
	scope.Memory.Set(memOffset.Uint64(), length.Uint64(), codeCopy)

	metaA := scope.metaStack.pop()
	metaMemOffset := scope.metaStack.pop()
	metaDataOffset := scope.metaStack.pop()
	metaLength := scope.metaStack.pop()
	metaCode := scope.metaCode.Get(addr)
	scope.metaMemory.Set(memOffset.Uint64(), length.Uint64(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaA, metaMemOffset, metaDataOffset, metaLength, metaCode}, *newMetaRes)

	return nil, nil
}

// opExtCodeHash returns the code hash of a specified account.
// There are several cases when the function is called, while we can relay everything
// to `state.GetCodeHash` function to ensure the correctness.
//
//  1. Caller tries to get the code hash of a normal contract account, state
//     should return the relative code hash and set it as the result.
//
//  2. Caller tries to get the code hash of a non-existent account, state should
//     return common.Hash{} and zero will be set as the result.
//
//  3. Caller tries to get the code hash for an account without contract code, state
//     should return emptyCodeHash(0xc5d246...) as the result.
//
//  4. Caller tries to get the code hash of a precompiled account, the result should be
//     zero or emptyCodeHash.
//
// It is worth noting that in order to avoid unnecessary create and clean, all precompile
// accounts on mainnet have been transferred 1 wei, so the return here should be
// emptyCodeHash. If the precompile account is not transferred any amount on a private or
// customized chain, the return value will be zero.
//
//  5. Caller tries to get the code hash for an account which is marked as self-destructed
//     in the current transaction, the code hash of this account should be returned.
//
//  6. Caller tries to get the code hash for an account which is marked as deleted, this
//     account should be regarded as a non-existent account and zero should be returned.

func opExtCodeHash(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, EXTCODEHASH)

	slot := scope.Stack.peek()
	address := common.Address(slot.Bytes20())
	if interpreter.evm.StateDB.Empty(address) {
		slot.Clear()
	} else {
		slot.SetBytes(interpreter.evm.StateDB.GetCodeHash(address).Bytes())
	}

	metaSlot := scope.metaStack.pop()
	metaCode := scope.metaCode.Get(address)
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaSlot, metaCode}, *newMetaRes)
	return nil, nil
}

func opGasprice(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, GASPRICE)

	v, _ := uint256.FromBig(interpreter.evm.GasPrice)
	scope.Stack.push(v)

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

func opBlockhash(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, BLOCKHASH)

	num := scope.Stack.peek()

	num64, overflow := num.Uint64WithOverflow()
	if overflow {
		num.Clear()
		return nil, nil
	}
	var upper, lower uint64
	upper = interpreter.evm.Context.BlockNumber.Uint64()
	if upper < 257 {
		lower = 0
	} else {
		lower = upper - 256
	}
	if num64 >= lower && num64 < upper {
		num.SetBytes(interpreter.evm.Context.GetHash(num64).Bytes())
	} else {
		num.Clear()
	}

	metaNum := scope.metaStack.pop()
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaNum}, *newMetaRes)
	return nil, nil
}

func opCoinbase(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, COINBASE)

	scope.Stack.push(new(uint256.Int).SetBytes(interpreter.evm.Context.Coinbase.Bytes()))

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

func opTimestamp(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, TIMESTAMP)

	scope.Stack.push(new(uint256.Int).SetUint64(interpreter.evm.Context.Time))

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

func opNumber(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, NUMBER)

	v, _ := uint256.FromBig(interpreter.evm.Context.BlockNumber)
	scope.Stack.push(v)

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

func opDifficulty(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, DIFFICULTY)

	v, _ := uint256.FromBig(interpreter.evm.Context.Difficulty)
	scope.Stack.push(v)

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

func opRandom(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, RANDOM)

	v := new(uint256.Int).SetBytes(interpreter.evm.Context.Random.Bytes())
	scope.Stack.push(v)

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

func opGasLimit(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, GASLIMIT)

	scope.Stack.push(new(uint256.Int).SetUint64(interpreter.evm.Context.GasLimit))

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{SourceMeta}, *newMetaRes)
	return nil, nil
}

// 我想这个指令会产生一个汇点，也许有了数据流图之后它可以被抛弃
// !! 后续考虑这个的优化
// !! 标记数据的同时应该还需要标记需要什么样的数据……可能需要规定opCode之间的数据传递模式
func opPop(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, GASLIMIT)

	scope.Stack.pop()

	meta := scope.metaStack.pop()

	interpreter.evm.Graph.AddDependency([]Metadata{meta}, *newMetaRes)
	return nil, nil
}

func opMload(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, MLOAD)

	v := scope.Stack.peek()
	offset := int64(v.Uint64())
	v.SetBytes(scope.Memory.GetPtr(offset, 32))

	metaV := scope.metaStack.pop()
	metaMem := scope.metaMemory.GetPtr(offset, 32)
	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaV, metaMem}, *newMetaRes)
	return nil, nil
}

func opMstore(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, MSTORE)

	// pop value of the stack
	mStart, val := scope.Stack.pop(), scope.Stack.pop()
	scope.Memory.Set32(mStart.Uint64(), &val)

	metaStart, metaVal := scope.metaStack.pop(), scope.metaStack.pop()
	scope.metaMemory.Set32(mStart.Uint64(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaStart, metaVal}, *newMetaRes)
	return nil, nil
}

func opMstore8(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, MSTORE8)

	off, val := scope.Stack.pop(), scope.Stack.pop()
	scope.Memory.store[off.Uint64()] = byte(val.Uint64())

	metaOff, metaVal := scope.metaStack.pop(), scope.metaStack.pop()
	scope.metaMemory.store[off.Uint64()] = *newMetaRes

	interpreter.evm.Graph.AddDependency([]Metadata{metaOff, metaVal}, *newMetaRes)
	return nil, nil
}

func opSload(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SLOAD)

	loc := scope.Stack.peek()
	hash := common.Hash(loc.Bytes32())
	val := interpreter.evm.StateDB.GetState(scope.Contract.Address(), hash)
	loc.SetBytes(val.Bytes())

	metaLoc := scope.metaStack.pop()
	metaVal := scope.metaStorage.Get(scope.Contract.Address(), hash)
	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{metaLoc, metaVal, source}, *newMetaRes)
	return nil, nil
}

func opSstore(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SSTORE)

	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	loc := scope.Stack.pop()
	val := scope.Stack.pop()
	interpreter.evm.StateDB.SetState(scope.Contract.Address(), loc.Bytes32(), val.Bytes32())

	metaLoc := scope.metaStack.pop()
	metaVal := scope.metaStack.pop()
	scope.metaStorage.Set(scope.Contract.Address(), loc.Bytes32(), *newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{metaLoc, metaVal, source}, *newMetaRes)
	return nil, nil
}

func opJump(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, JUMP)

	if interpreter.evm.abort.Load() {
		return nil, errStopToken
	}
	pos := scope.Stack.pop()
	if !scope.Contract.validJumpdest(&pos) {
		return nil, ErrInvalidJump
	}
	*pc = pos.Uint64() - 1 // pc will be increased by the interpreter loop

	metaPos := scope.metaStack.pop()

	interpreter.evm.Graph.AddDependency([]Metadata{metaPos}, *newMetaRes)
	return nil, nil
}

func opJumpi(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, JUMPI)

	if interpreter.evm.abort.Load() {
		return nil, errStopToken
	}
	pos, cond := scope.Stack.pop(), scope.Stack.pop()
	if !cond.IsZero() {
		if !scope.Contract.validJumpdest(&pos) {
			return nil, ErrInvalidJump
		}
		*pc = pos.Uint64() - 1 // pc will be increased by the interpreter loop
	}
	metaPos := scope.metaStack.pop()
	metaCond := scope.metaStack.pop()

	interpreter.evm.Graph.AddDependency([]Metadata{metaPos, metaCond}, *newMetaRes)
	return nil, nil
}

func opJumpdest(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {

	return nil, nil
}

func opPc(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, PC)

	scope.Stack.push(new(uint256.Int).SetUint64(*pc))

	scope.metaStack.push(newMetaRes)

	// pc_last_modify 一定是当前index - 1
	pc_last_modify := interpreter.evm.Graph.Vertexes[newMetaRes.Index-1]
	interpreter.evm.Graph.AddDependency([]Metadata{pc_last_modify}, *newMetaRes)
	return nil, nil
}

func opMsize(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, MSIZE)

	scope.Stack.push(new(uint256.Int).SetUint64(uint64(scope.Memory.Len())))

	scope.metaStack.push(newMetaRes)

	memoryLenLastModify := interpreter.evm.Graph.Vertexes[scope.memory_len_last_modify]
	interpreter.evm.Graph.AddDependency([]Metadata{memoryLenLastModify}, *newMetaRes)
	return nil, nil
}

// 从scope.Contract里面拿数据了
func opGas(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, GAS)

	scope.Stack.push(new(uint256.Int).SetUint64(scope.Contract.Gas))

	scope.metaStack.push(newMetaRes)

	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, nil
}

// 从他拿数据的地方构建依赖，但要更新interpreter和contract
func opCreate(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CREATE)
	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	var (
		value        = scope.Stack.pop()
		offset, size = scope.Stack.pop(), scope.Stack.pop()
		input        = scope.Memory.GetCopy(int64(offset.Uint64()), int64(size.Uint64()))
		gas          = scope.Contract.Gas
	)
	if interpreter.evm.chainRules.IsEIP150 {
		gas -= gas / 64
	}
	// reuse size int for stackvalue
	stackvalue := size

	scope.Contract.UseGas(gas)
	//TODO: use uint256.Int instead of converting with toBig()
	var bigVal = big0
	if !value.IsZero() {
		bigVal = value.ToBig()
	}

	res, addr, returnGas, suberr := interpreter.evm.Create(scope.Contract, input, gas, bigVal, *scope.opCodeCounter)
	// Push item on the stack based on the returned error. If the ruleset is
	// homestead we must check for CodeStoreOutOfGasError (homestead only
	// rule) and treat as an error, if the ruleset is frontier we must
	// ignore this error and pretend the operation was successful.
	if interpreter.evm.chainRules.IsHomestead && suberr == ErrCodeStoreOutOfGas {
		stackvalue.Clear()
	} else if suberr != nil && suberr != ErrCodeStoreOutOfGas {
		stackvalue.Clear()
	} else {
		stackvalue.SetBytes(addr.Bytes())
	}
	scope.Stack.push(&stackvalue)

	metaValue := scope.metaStack.pop()
	metaOffset := scope.metaStack.pop()
	metaSize := scope.metaStack.pop()
	metaInput := scope.metaMemory.GetCopy(int64(offset.Uint64()), int64(size.Uint64()))

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaValue, metaOffset, metaSize, metaInput}, *newMetaRes)

	// 改了scope.Contract中的东西，sourceIndex就会改变
	scope.Contract.Gas += returnGas
	scope.Contract.SourceIndex = newMetaRes.Index

	if suberr == ErrExecutionReverted {
		// 改了Interpreter.returnData，sourceIndex就会改变
		interpreter.returnData = res // set REVERT data to return data buffer
		interpreter.sourceIndex = newMetaRes.Index
		return res, nil
	}
	// 改了Interpreter.returnData，sourceIndex就会改变
	interpreter.returnData = nil // clear dirty return data buffer
	interpreter.sourceIndex = newMetaRes.Index
	return nil, nil
}

func opCreate2(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CREATE2)

	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	var (
		endowment    = scope.Stack.pop()
		offset, size = scope.Stack.pop(), scope.Stack.pop()
		salt         = scope.Stack.pop()
		input        = scope.Memory.GetCopy(int64(offset.Uint64()), int64(size.Uint64()))
		gas          = scope.Contract.Gas
	)
	// Apply EIP150
	gas -= gas / 64
	scope.Contract.UseGas(gas)
	// reuse size int for stackvalue
	stackvalue := size
	//TODO: use uint256.Int instead of converting with toBig()
	bigEndowment := big0
	if !endowment.IsZero() {
		bigEndowment = endowment.ToBig()
	}
	res, addr, returnGas, suberr := interpreter.evm.Create2(scope.Contract, input, gas,
		bigEndowment, &salt, *scope.opCodeCounter)
	// Push item on the stack based on the returned error.
	if suberr != nil {
		stackvalue.Clear()
	} else {
		stackvalue.SetBytes(addr.Bytes())
	}
	scope.Stack.push(&stackvalue)

	metaEndowment := scope.metaStack.pop()
	metaOffset := scope.metaStack.pop()
	metaSize := scope.metaStack.pop()
	metaSalt := scope.metaStack.pop()
	metaInput := scope.metaMemory.GetCopy(int64(offset.Uint64()), int64(size.Uint64()))

	scope.metaStack.push(newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaEndowment, metaOffset, metaSize, metaSalt, metaInput}, *newMetaRes)

	scope.Contract.Gas += returnGas
	scope.Contract.SourceIndex = newMetaRes.Index

	if suberr == ErrExecutionReverted {
		interpreter.returnData = res // set REVERT data to return data buffer
		interpreter.sourceIndex = newMetaRes.Index
		return res, nil
	}
	interpreter.returnData = nil // clear dirty return data buffer
	interpreter.sourceIndex = newMetaRes.Index
	return nil, nil
}

// call 相关的依赖图比较难画
// 可以考虑不直接加依赖，因为参数会被传入Call函数，依赖关系由Call函数来衍生
// 即基础OpCode会把依赖加上
// 这是否意味着Call相关从Contract上下文里获取的数据，都需要补充meta数据？

func opCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALL)

	stack := scope.Stack
	// Pop gas. The actual gas in interpreter.evm.callGasTemp.
	// We can use this as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, value, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get the arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	if interpreter.readOnly && !value.IsZero() {
		return nil, ErrWriteProtection
	}
	var bigVal = big0
	//TODO: use uint256.Int instead of converting with toBig()
	// By using big0 here, we save an alloc for the most common case (non-ether-transferring contract calls),
	// but it would make more sense to extend the usage of uint256.Int
	if !value.IsZero() {
		gas += params.CallStipend
		bigVal = value.ToBig()
	}

	ret, returnGas, err := interpreter.evm.Call(scope.Contract, toAddr, args, gas, bigVal, *scope.opCodeCounter)

	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}

	metaTemp := scope.metaStack.pop()
	metaAddr := scope.metaStack.pop()
	metaValue := scope.metaStack.pop()
	metaInOffset := scope.metaStack.pop()
	metaInSize := scope.metaStack.pop()
	metaRetOffset := scope.metaStack.pop()
	metaRetSize := scope.metaStack.pop()
	metaArgs := scope.metaMemory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	// two results
	scope.metaStack.push(newMetaRes)
	scope.metaMemory.Set(retOffset.Uint64(), retSize.Uint64(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaTemp, metaAddr, metaValue, metaInOffset, metaInSize, metaRetOffset, metaRetSize, metaArgs}, *newMetaRes)

	scope.Contract.Gas += returnGas
	scope.Contract.SourceIndex = newMetaRes.Index

	interpreter.returnData = ret
	interpreter.sourceIndex = newMetaRes.Index
	return ret, nil
}

func opCallCode(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, CALLCODE)

	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	stack := scope.Stack
	// We use it as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, value, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	//TODO: use uint256.Int instead of converting with toBig()
	var bigVal = big0
	if !value.IsZero() {
		gas += params.CallStipend
		bigVal = value.ToBig()
	}

	ret, returnGas, err := interpreter.evm.CallCode(scope.Contract, toAddr, args, gas, bigVal, *scope.opCodeCounter)
	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}

	metaTemp := scope.metaStack.pop()
	metaAddr := scope.metaStack.pop()
	metaValue := scope.metaStack.pop()
	metaInOffset := scope.metaStack.pop()
	metaInSize := scope.metaStack.pop()
	metaRetOffset := scope.metaStack.pop()
	metaRetSize := scope.metaStack.pop()
	metaArgs := scope.metaMemory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	// two results
	scope.metaStack.push(newMetaRes)
	scope.metaMemory.Set(retOffset.Uint64(), retSize.Uint64(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaTemp, metaAddr, metaValue, metaInOffset, metaInSize, metaRetOffset, metaRetSize, metaArgs}, *newMetaRes)

	scope.Contract.Gas += returnGas
	scope.Contract.SourceIndex = newMetaRes.Index

	interpreter.returnData = ret
	interpreter.sourceIndex = newMetaRes.Index
	return ret, nil
}

func opDelegateCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, DELEGATECALL)
	stack := scope.Stack
	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	// We use it as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	ret, returnGas, err := interpreter.evm.DelegateCall(scope.Contract, toAddr, args, gas, *scope.opCodeCounter)
	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}

	metaTemp := scope.metaStack.pop()
	metaAddr := scope.metaStack.pop()
	metaInOffset := scope.metaStack.pop()
	metaInSize := scope.metaStack.pop()
	metaRetOffset := scope.metaStack.pop()
	metaRetSize := scope.metaStack.pop()
	metaArgs := scope.metaMemory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	// two results
	scope.metaStack.push(newMetaRes)
	scope.metaMemory.Set(retOffset.Uint64(), retSize.Uint64(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaTemp, metaAddr, metaInOffset, metaInSize, metaRetOffset, metaRetSize, metaArgs}, *newMetaRes)

	scope.Contract.Gas += returnGas
	scope.Contract.SourceIndex = newMetaRes.Index

	interpreter.returnData = ret
	interpreter.sourceIndex = newMetaRes.Index
	return ret, nil
}

func opStaticCall(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, STATICCALL)
	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	stack := scope.Stack
	// We use it as a temporary value
	temp := stack.pop()
	gas := interpreter.evm.callGasTemp
	// Pop other call parameters.
	addr, inOffset, inSize, retOffset, retSize := stack.pop(), stack.pop(), stack.pop(), stack.pop(), stack.pop()
	toAddr := common.Address(addr.Bytes20())
	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	ret, returnGas, err := interpreter.evm.StaticCall(scope.Contract, toAddr, args, gas, *scope.opCodeCounter)
	if err != nil {
		temp.Clear()
	} else {
		temp.SetOne()
	}
	stack.push(&temp)
	if err == nil || err == ErrExecutionReverted {
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	}

	metaTemp := scope.metaStack.pop()
	metaAddr := scope.metaStack.pop()
	metaInOffset := scope.metaStack.pop()
	metaInSize := scope.metaStack.pop()
	metaRetOffset := scope.metaStack.pop()
	metaRetSize := scope.metaStack.pop()
	metaArgs := scope.metaMemory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

	scope.metaStack.push(newMetaRes)
	scope.metaMemory.Set(retOffset.Uint64(), retSize.Uint64(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaTemp, metaAddr, metaInOffset, metaInSize, metaRetOffset, metaRetSize, metaArgs}, *newMetaRes)

	scope.Contract.Gas += returnGas
	scope.Contract.SourceIndex = newMetaRes.Index

	interpreter.returnData = ret
	interpreter.sourceIndex = newMetaRes.Index
	return ret, nil
}

func opReturn(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, RETURN)

	offset, size := scope.Stack.pop(), scope.Stack.pop()
	ret := scope.Memory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	metaOffset := scope.metaStack.pop()
	metaSize := scope.metaStack.pop()
	metaRet := scope.metaMemory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	interpreter.evm.Graph.AddDependency([]Metadata{metaOffset, metaSize, metaRet}, *newMetaRes)
	return ret, errStopToken
}

func opRevert(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, REVERT)

	offset, size := scope.Stack.pop(), scope.Stack.pop()
	ret := scope.Memory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	metaOffset := scope.metaStack.pop()
	metaSize := scope.metaStack.pop()
	metaRet := scope.metaMemory.GetPtr(int64(offset.Uint64()), int64(size.Uint64()))

	interpreter.returnData = ret
	interpreter.sourceIndex = newMetaRes.Index

	interpreter.evm.Graph.AddDependency([]Metadata{metaOffset, metaSize, metaRet}, *newMetaRes)
	return ret, ErrExecutionReverted
}

// 没有对应的opcode，我们先跳过吧
func opUndefined(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	return nil, &ErrInvalidOpCode{opcode: OpCode(scope.Contract.Code[*pc])}
}

// 个人理解，stop需要依赖上一个index的指令（我猜是一个jump）
func opStop(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, STOP)
	source := interpreter.evm.Graph.Vertexes[newMetaRes.Index-1]
	interpreter.evm.Graph.AddDependency([]Metadata{source}, *newMetaRes)
	return nil, errStopToken
}

// 从stack、balance、contract拿数据
// 改写了balance
func opSelfdestruct(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SELFDESTRUCT)

	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	beneficiary := scope.Stack.pop()
	balance := interpreter.evm.StateDB.GetBalance(scope.Contract.Address())
	interpreter.evm.StateDB.AddBalance(beneficiary.Bytes20(), balance)
	interpreter.evm.StateDB.SelfDestruct(scope.Contract.Address())
	if tracer := interpreter.evm.Config.Tracer; tracer != nil {
		tracer.CaptureEnter(SELFDESTRUCT, scope.Contract.Address(), beneficiary.Bytes20(), []byte{}, 0, balance)
		tracer.CaptureExit([]byte{}, 0, nil)
	}

	metaBeneficiary := scope.metaStack.pop()
	metaBalance := scope.metaStorage.Get(scope.Contract.Address(), common.Hash{})
	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]

	scope.metaBalance.Set(beneficiary.Bytes20(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaBeneficiary, metaBalance, source}, *newMetaRes)
	return nil, errStopToken
}

// 从stack、balance、contract拿数据
// 改写了contractAddr\beneficairy的balance
func opSelfdestruct6780(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, SELFDESTRUCT)

	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	beneficiary := scope.Stack.pop()
	balance := interpreter.evm.StateDB.GetBalance(scope.Contract.Address())
	interpreter.evm.StateDB.SubBalance(scope.Contract.Address(), balance)
	interpreter.evm.StateDB.AddBalance(beneficiary.Bytes20(), balance)
	interpreter.evm.StateDB.Selfdestruct6780(scope.Contract.Address())
	if tracer := interpreter.evm.Config.Tracer; tracer != nil {
		tracer.CaptureEnter(SELFDESTRUCT, scope.Contract.Address(), beneficiary.Bytes20(), []byte{}, 0, balance)
		tracer.CaptureExit([]byte{}, 0, nil)
	}

	metaBeneficiary := scope.metaStack.pop()
	metaBalance := scope.metaStorage.Get(scope.Contract.Address(), common.Hash{})
	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]

	scope.metaBalance.Set(beneficiary.Bytes20(), *newMetaRes)
	scope.metaBalance.Set(scope.Contract.Address(), *newMetaRes)

	interpreter.evm.Graph.AddDependency([]Metadata{metaBeneficiary, metaBalance, source}, *newMetaRes)
	return nil, errStopToken
}

// following functions are used by the instruction jump  table

// make log instruction function
func makeLog(size int) executionFunc {
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		if interpreter.readOnly {
			return nil, ErrWriteProtection
		}
		topics := make([]common.Hash, size)
		stack := scope.Stack
		mStart, mSize := stack.pop(), stack.pop()
		for i := 0; i < size; i++ {
			addr := stack.pop()
			topics[i] = addr.Bytes32()
		}

		d := scope.Memory.GetCopy(int64(mStart.Uint64()), int64(mSize.Uint64()))
		interpreter.evm.StateDB.AddLog(&types.Log{
			Address: scope.Contract.Address(),
			Topics:  topics,
			Data:    d,
			// This is a non-consensus field, but assigned here because
			// core/state doesn't know the current block number.
			BlockNumber: interpreter.evm.Context.BlockNumber.Uint64(),
		})

		return nil, nil
	}
}

// opPush1 is a specialized version of pushN，推入下一个操作码
// 从contract、pc拿数据
// 修改pc，stack
func opPush1(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, PUSH1)

	var (
		codeLen = uint64(len(scope.Contract.Code))
		integer = new(uint256.Int)
	)
	*pc += 1
	if *pc < codeLen {
		scope.Stack.push(integer.SetUint64(uint64(scope.Contract.Code[*pc])))
	} else {
		scope.Stack.push(integer.Clear())
	}

	scope.metaStack.push(newMetaRes)

	last_pc_modifer := interpreter.evm.Graph.Vertexes[newMetaRes.Index-1]
	source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]

	interpreter.evm.Graph.AddDependency([]Metadata{last_pc_modifer, source}, *newMetaRes)
	return nil, nil
}

// make push instruction function, 压入下面size个opcode
// 从contract、pc拿数据
// 修改pc，stack
func makePush(size uint64, pushByteSize int) executionFunc {
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		opcode := PUSH2 + OpCode(size-2)
		newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, opcode)

		codeLen := len(scope.Contract.Code)

		startMin := codeLen
		if int(*pc+1) < startMin {
			startMin = int(*pc + 1)
		}

		endMin := codeLen
		if startMin+pushByteSize < endMin {
			endMin = startMin + pushByteSize
		}

		integer := new(uint256.Int)
		scope.Stack.push(integer.SetBytes(common.RightPadBytes(
			scope.Contract.Code[startMin:endMin], pushByteSize)))
		*pc += size

		scope.metaStack.push(newMetaRes)

		last_pc_modifer := interpreter.evm.Graph.Vertexes[newMetaRes.Index-1]
		source := interpreter.evm.Graph.Vertexes[scope.Contract.SourceIndex]

		interpreter.evm.Graph.AddDependency([]Metadata{last_pc_modifer, source}, *newMetaRes)
		return nil, nil
	}
}

// make dup instruction function
// 把第size（？）个元素复制到栈顶
// 涉及到从stack拿数据，修改stack
// 同时这个size从哪里定的呢？我只能假定这种直接对stack的操作对前一个opcode很有依赖性，从运行结果上来看，没有，这个在编译上就确定了
func makeDup(size int64) executionFunc {
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		opcode := DUP1 + OpCode(size-1)
		newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, opcode)

		scope.Stack.dup(int(size))

		relatedMeta := scope.metaStack.data[scope.metaStack.len()-int(size)]

		//!! dup只是push一个进去，明天从这里开始DEBUG起
		scope.metaStack.push(newMetaRes)

		interpreter.evm.Graph.AddDependency([]Metadata{relatedMeta}, *newMetaRes)
		return nil, nil
	}
}

// make swap instruction function
func makeSwap(size int64) executionFunc {
	// switch n + 1 otherwise n would be swapped with n
	size++
	return func(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
		opcode := SWAP1 + OpCode(size-1)
		newMetaRes := NewMetadata(interpreter.evm.StateDB.GetTxId(), *scope.opCodeCounter, *scope.Contract.CodeAddr, *pc, opcode)

		scope.Stack.swap(int(size))

		metaTarget := scope.metaStack.data[scope.metaStack.len()-int(size)]
		metaPeek := scope.metaStack.data[scope.metaStack.len()-1]

		// 因为修改了stack，所以需要更新metaStack
		scope.metaStack.data[scope.metaStack.len()-int(size)] = *newMetaRes
		scope.metaStack.data[scope.metaStack.len()-1] = *newMetaRes

		interpreter.evm.Graph.AddDependency([]Metadata{metaTarget, metaPeek}, *newMetaRes)
		return nil, nil
	}
}
