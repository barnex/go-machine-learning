package kernel

import (
	"fmt"
	"github.com/mumax/3/cuda/cu"
)

// CUDA Launch parameters.
// there might be better choices for recent hardware,
// but it barely makes a difference in the end.
const (
	BlockSize    = 512
	TileX, TileY = 32, 32
	MaxGridSize  = 65535
)

// cuda launch configuration
type Config struct {
	Grid, Block cu.Dim3
}

// Make a 1D kernel launch configuration suited for N threads.
func Make1DConf(N int) *Config {
	bl := cu.Dim3{X: BlockSize, Y: 1, Z: 1}

	n2 := divUp(N, BlockSize) // N2 blocks left
	nx := divUp(n2, MaxGridSize)
	ny := divUp(n2, nx)
	gr := cu.Dim3{X: nx, Y: ny, Z: 1}

	return &Config{gr, bl}
}

// Make a 3D kernel launch configuration suited for N threads.
func Make3DConf(N [3]int) *Config {
	bl := cu.Dim3{X: TileX, Y: TileY, Z: 1}

	nx := divUp(N[X], TileX)
	ny := divUp(N[Y], TileY)
	gr := cu.Dim3{X: nx, Y: ny, Z: N[Z]}

	return &Config{gr, bl}
}

// integer minimum
func iMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Integer division rounded up.
func divUp(x, y int) int {
	return ((x - 1) / y) + 1
}

const (
	X = 0
	Y = 1
	Z = 2
)

func checkSize(a interface {
	Size() [3]int
}, b ...interface {
	Size() [3]int
}) {
	sa := a.Size()
	for _, b := range b {
		if b.Size() != sa {
			panic(fmt.Sprintf("size mismatch: %v != %v", sa, b.Size()))
		}
	}
}
