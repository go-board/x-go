package xruntime

import (
	"unsafe"
)

type TypeFlag uint8
type NameOffset int32
type TypeOffset int32

type Type struct {
	Size       uintptr
	PtrData    uintptr // Size of memory prefix holding all pointers
	Hash       uint32
	TypeFlag   TypeFlag
	Align      uint8
	FieldAlign uint8
	Kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
	// Gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in Kind, Gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	Gcdata    *byte
	Str       NameOffset
	PtrToThis TypeOffset
}

type Name struct {
	Bytes *byte
}

type Method struct {
	Name NameOffset
	Type TypeOffset
}

type InterfaceType struct {
	Type    Type
	PkgPath Name
	Methods []Method
}

type VTable struct {
	IType *InterfaceType
	Type  *Type
	Hash  uint32 // copy of _type.Hash. Used for type switches.
	_     [4]byte
	Fun   [1]uintptr // variable sized. Fun[0]==0 means _type does not implement IType.
}

type Interface struct {
	VTable *VTable
	Data   unsafe.Pointer
}

type EmptyInterface struct {
	Type *Type
	Data unsafe.Pointer
}

func InterfaceData(v interface{}) *Interface {
	return (*Interface)(unsafe.Pointer(&v))
}

func EmptyInterfaceData(v interface{}) *EmptyInterface {
	return (*EmptyInterface)(unsafe.Pointer(&v))
}
