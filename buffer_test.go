package vs

import (
	"runtime"
	"testing"

	"github.com/barnex/cuda5/cu"
)

func TestBuffer(t *testing.T) {
	runtime.LockOSThread()
	cu.Init(0)
	dev := cu.Device(*flagGPUID)
	cudaCtx = cu.CtxCreate(cu.CTX_SCHED_YIELD, dev)
	cudaCtx.SetCurrent()

	for i := 0; i < 1000; i++ {
		b := NewBuffer(Size{1, 1024, 1024, 3})
		b.Free()
	}
}
