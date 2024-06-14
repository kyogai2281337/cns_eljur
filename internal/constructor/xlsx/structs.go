package xlsx

type BlockSize struct {
	X int
	Y int
}
type Block struct {
	Size BlockSize
}

func NewBlock(bs BlockSize) Block {
	return Block{
		Size: bs,
	}
}

func NewBlockSize(x, y int) BlockSize {
	return BlockSize{
		X: x,
		Y: y + 1,
	}
}
