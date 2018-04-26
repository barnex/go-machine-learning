package vs

import (
	"testing"

	"github.com/barnex/vectorstream/kernel"
)

func TestMAdd(t *testing.T) {
	InitTest()

	N := 1024
	s := Size{1, 1, 1, N}
	dst := NewGPUBuf(s)
	defer dst.Free()
	src1 := NewGPUBuf(s)
	defer src1.Free()
	src2 := NewGPUBuf(s)
	defer src2.Free()

	kernel.Kmadd2(dst.UPtr(), src1.UPtr(), 1, src2.UPtr(), 1, N, kernel.Make1DConf(N))
}
