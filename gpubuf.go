package vs

import (
	"flag"
	"fmt"

	"github.com/barnex/cuda5/cu"
)

var flagGPUID = flag.Int("gpu", 0, "Index of GPU to use")

var cudaCtx cu.Context

type GPUBuf struct {
	Size
	ptr cu.DevicePtr
}

func NewGPUBuf(s Size) *GPUBuf {
	ptr := cu.MemAlloc(s.bytes())
	return &GPUBuf{s, ptr}
}

func (b *GPUBuf) Free() {
	if b.ptr == 0 {
		return
	}
	cu.MemFree(b.ptr)
	b.ptr = 0
}

func (b *GPUBuf) Upload(h *HostBuf) {
	checkEqualSize(b.Size, h.Size)
	cu.MemcpyHtoD(b.ptr, h.ptr(), b.Size.bytes())
}

func (b *GPUBuf) HostCopy() *HostBuf {
	h := NewHostBuf(b.Size)
	cu.MemcpyDtoH(h.ptr(), b.ptr, b.Size.bytes())
	return h
}

func checkEqualSize(a, b Size) {
	if a != b {
		panic(fmt.Sprintf("sizes not equal: %v, %v", a, b))
	}
}
