package vs

import (
	"unsafe"
)

type HostBuf struct {
	size Size
	data []float32
}

func NewHostBuf(s Size) *HostBuf {
	return &HostBuf{s, make([]float32, s.Len())}
}

func (h *HostBuf) Size() Size {
	return h.size
}

func (h *HostBuf) ptr() unsafe.Pointer {
	return unsafe.Pointer(&h.data[0])
}

func (h *HostBuf) Data() []float32 {
	return h.data
}

//func Reshape(array []float32, N Size) [][][][]float32 {
//	if N.Len() != len(array) {
//		panic(fmt.Errorf("reshape: size %v does not match len %v", N, len(array)))
//	}
//	sliced := make([][][][]float32, N[0])
//	for i := range sliced {
//		sliced[i] = make([][][]float32, N[1])
//	}
//	for i := range sliced {
//		for j := range sliced[i] {
//			sliced[i][j] = make([][]float32, N[2])
//		}
//	}
//	for i := range sliced {
//		for j := range sliced[i] {
//			for k := range sliced[i][j] {
//				sliced[i][j][k] = array[((i*N[1]+j)*N[2]+k)*N[3]+0 : ((i*N[1]+j)*N[2]+k)*N[3]+N[3]]
//			}
//		}
//	}
//	return sliced
//}
