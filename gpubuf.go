package vs

import (
	"flag"
	"fmt"
	"math"
	"unsafe"

	"github.com/barnex/cuda5/cu"
)

var flagGPUID = flag.Int("gpu", 0, "Index of GPU to use")

var cudaCtx cu.Context

type GPUBuf struct {
	size Size
	ptr  cu.DevicePtr
}

func NewGPUBuf(s Size) *GPUBuf {
	ptr := cu.MemAlloc(s.bytes())
	return &GPUBuf{s, ptr}
}

func (b *GPUBuf) Size() Size {
	return b.size
}

func (b *GPUBuf) Ptr() cu.DevicePtr {
	return b.ptr
}

func (b *GPUBuf) UPtr() unsafe.Pointer {
	return unsafe.Pointer(b.ptr)
}

func (b *GPUBuf) Free() {
	if b.ptr == 0 {
		panic("double free")
	}
	cu.MemFree(b.ptr)
	b.ptr = 0
}

func (b *GPUBuf) CopyFrom(src *GPUBuf) {
	checkEqualSize(b.Size(), src.Size())
	cu.MemcpyDtoD(b.ptr, src.ptr, b.size.bytes())
}

func (b *GPUBuf) Upload(h *HostBuf) {
	checkEqualSize(b.Size(), h.Size())
	cu.MemcpyHtoD(b.ptr, h.ptr(), b.size.bytes())
}

func (b *GPUBuf) HostCopy() *HostBuf {
	h := NewHostBuf(b.size)
	cu.MemcpyDtoH(h.ptr(), b.ptr, b.size.bytes())
	return h
}

func (b *GPUBuf) Memset(v float32) {
	cu.MemsetD32Async(b.Ptr(), math.Float32bits(v), int64(b.size.Len()), cu.Stream(0))
	cu.Stream(0).Synchronize()
}

func checkEqualSize(a, b Size) {
	if a != b {
		panic(fmt.Sprintf("sizes not equal: %v, %v", a, b))
	}
}
