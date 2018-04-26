package vs

import (
	"reflect"
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
	defer func() {
		if p := recover(); p == nil {
			t.Error("expected panic")
		}
	}()
	b := NewGPUBuf(Size{1, 1024, 1024, 3})
	b.Free()
	b.Free() // panic: double-free
}

func TestGPUBuf_DoubleFree(t *testing.T) {
	InitTest()
	for i := 0; i < 1000; i++ {
		b := NewGPUBuf(Size{1, 1024, 1024, 3})
		b.Free()
	}
}

func TestGPUBuf_Copy(t *testing.T) {
	InitTest()

	h := NewHostBuf(Size{1, 512, 1024, 3})
	for i := range h.Data() {
		h.Data()[i] = float32(i)
	}

	b := NewGPUBuf(h.Size())
	defer b.Free()
	b.Upload(h)

	b2 := NewGPUBuf(b.Size())
	b2.CopyFrom(b)

	g := b2.HostCopy()
	for i, v := range g.Data() {
		if v != float32(i) {
			t.Fatalf("have: %v, want: %v", v, i)
		}
	}
}

func TestGPUBuf_Memset(t *testing.T) {
	InitTest()

	b := NewGPUBuf(Size{1, 1, 1, 4})
	defer b.Free()

	b.Memset(42)
	h := b.HostCopy().Data()
	want := []float32{42, 42, 42, 42}
	if !reflect.DeepEqual(h, want) {
		t.Errorf("have: %v, want: %v", h, want)
	}
}

func TestGPUBuf_CopyBadSize(t *testing.T) {
	InitTest()

	defer func() {
		if p := recover(); p == nil {
			t.Error("expected panic")
		}
	}()

	h := NewHostBuf(Size{1, 512, 1024, 3})
	for i := range h.Data() {
		h.Data()[i] = float32(i)
	}

	b := NewGPUBuf(Size{2, 512, 1024, 3})
	defer b.Free()
	b.Upload(h)
}
