package day2

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return v.cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

// 返回一个与原切片数据相同但独立的副本
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
