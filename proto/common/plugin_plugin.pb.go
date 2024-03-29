//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v4.25.2
// source: common/plugin.proto

package common

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

const CommonPluginAPIVersion = 1

//export common_api_version
func _common_api_version() uint64 {
	return CommonPluginAPIVersion
}

var common Common

func RegisterCommon(p Common) {
	common = p
}

//export common_prepare
func _common_prepare(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	req := new(Request)
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := common.Prepare(context.Background(), req)
	if err != nil {
		ptr, size = wasm.ByteToPtr([]byte(err.Error()))
		return (uint64(ptr) << uint64(32)) | uint64(size) |
			// Indicate that this is the error string by setting the 32-th bit, assuming that
			// no data exceeds 31-bit size (2 GiB).
			(1 << 31)
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export common_run
func _common_run(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	req := new(Request)
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := common.Run(context.Background(), req)
	if err != nil {
		ptr, size = wasm.ByteToPtr([]byte(err.Error()))
		return (uint64(ptr) << uint64(32)) | uint64(size) |
			// Indicate that this is the error string by setting the 32-th bit, assuming that
			// no data exceeds 31-bit size (2 GiB).
			(1 << 31)
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export common_check
func _common_check(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	req := new(Request)
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := common.Check(context.Background(), req)
	if err != nil {
		ptr, size = wasm.ByteToPtr([]byte(err.Error()))
		return (uint64(ptr) << uint64(32)) | uint64(size) |
			// Indicate that this is the error string by setting the 32-th bit, assuming that
			// no data exceeds 31-bit size (2 GiB).
			(1 << 31)
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

//export common_rollback
func _common_rollback(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	req := new(Request)
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := common.Rollback(context.Background(), req)
	if err != nil {
		ptr, size = wasm.ByteToPtr([]byte(err.Error()))
		return (uint64(ptr) << uint64(32)) | uint64(size) |
			// Indicate that this is the error string by setting the 32-th bit, assuming that
			// no data exceeds 31-bit size (2 GiB).
			(1 << 31)
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
