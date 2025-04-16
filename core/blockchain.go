package core

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesis := NewBlock(nil, []byte{}, 0)
	return &Blockchain{Blocks: []*Block{genesis}}
}

func (bc *Blockchain) AddBlock(txs []Transaction) {
	last := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(txs, last.Hash, last.Header.Height+1)
	bc.Blocks = append(bc.Blocks, newBlock)
}
