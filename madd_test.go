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
	dst.Memset(0)

	src1 := NewGPUBuf(s)
	defer src1.Free()
	src1.Memset(2)

	src2 := NewGPUBuf(s)
	defer src2.Free()
	src2.Memset(3)

	kernel.Kmadd2(dst.UPtr(), src1.UPtr(), 1, src2.UPtr(), 2, N, kernel.Make1DConf(N))

	have := dst.HostCopy().Data()[N-1]
	if have != 8 {
		t.Errorf("have: %v, want: %v", have, 8.0)
	}
}
