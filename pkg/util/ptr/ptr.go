// Copyright (c) ZStack.io, Inc.

// Package ptr provides utility functions for working with pointers.
package ptr

// Of returns a pointer to the given value.
// This is useful when you need to pass a pointer to a literal value.
//
// Example usage:
//
//	param := CreateVmInstanceParam{
//	    CpuNum:     ptr.Of(4),
//	    MemorySize: ptr.Of(int64(8589934592)),
//	    Name:       ptr.Of("my-vm"),
//	    Platform:   ptr.Of("Linux"),
//	    Tags:       ptr.Of([]string{"tag1", "tag2"}),
//	}
func Of[T any](v T) *T {
	return &v
}

// ValueOr returns the value pointed to by ptr, or defaultVal if ptr is nil.
// This is the inverse of Of - useful when reading optional fields.
//
// Example usage:
//
//	cpuNum := ptr.ValueOr(param.CpuNum, 1)
func ValueOr[T any](ptr *T, defaultVal T) T {
	if ptr == nil {
		return defaultVal
	}
	return *ptr
}

// Value returns the value pointed to by ptr, or the zero value if ptr is nil.
func Value[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}
