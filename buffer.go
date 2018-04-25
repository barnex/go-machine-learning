package vs

import (
	"flag"

	"github.com/barnex/cuda5/cu"
)

var flagGPUID = flag.Int("gpu", 0, "Index of GPU to use")

var cudaCtx cu.Context

type Buffer struct {
	Size
	ptr cu.DevicePtr
}

func NewBuffer(s Size) *Buffer {
	ptr := cu.MemAlloc(int64(s.Len()) * sizeOfFloat32)
	return &Buffer{s, ptr}
}

func (b *Buffer) Free() {
	if b.ptr == 0 {
		return
	}
	cu.MemFree(b.ptr)
	b.ptr = 0
}

const sizeOfFloat32 = 4
