package vs

import (
	"runtime"
	"testing"

	"github.com/barnex/cuda5/cu"
)

func InitTest() {
	runtime.LockOSThread()
	cu.Init(0)
	dev := cu.Device(*flagGPUID)
	cudaCtx = cu.CtxCreate(cu.CTX_SCHED_YIELD, dev)
	cudaCtx.SetCurrent()
}

func TestGPUBuf_Free(t *testing.T) {
	InitTest()
	for i := 0; i < 1000; i++ {
		b := NewGPUBuf(Size{1, 1024, 1024, 3})
		b.Free()
		b.Free()
	}
}

func TestGPUBuf_Copy(t *testing.T) {
	InitTest()

	b := NewGPUBuf(Size{1, 1024, 1024, 3})
	defer b.Free()

	h := b.HostCopy()
	for _, v := range h.Data() {
		if v != 0 {
			t.Fatalf("have: %v, want: 0", v)
		}
	}

	for i := range h.Data() {
		h.Data()[i] = float32(i)
	}

	b.Upload(h)

	g := b.HostCopy()
	for i, v := range g.Data() {
		if v != float32(i) {
			t.Fatalf("have: %v, want: %v", v, i)
		}
	}

}
