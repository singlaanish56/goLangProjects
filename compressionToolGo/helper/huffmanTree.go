package helper

type HuffmanNode interface {
	Leaf() bool
	Weight() int
}

type HuffmanLeafNode struct {
	element rune
	weight  int
}

func (ln *HuffmanLeafNode) Leaf() bool {
	return true
}

func (ln *HuffmanLeafNode) Weight() int {
	return ln.weight
}

func (ln *HuffmanLeafNode) Element() rune {
	return ln.element
}

type HuffmanInternalNode struct {
	weight int
	left   HuffmanNode
	right  HuffmanNode
}

func (ln *HuffmanInternalNode) Leaf() bool {
	return false
}

func (ln *HuffmanInternalNode) Weight() int {
	return ln.weight
}

func (ln *HuffmanInternalNode) Left() HuffmanNode {
	return ln.left
}

func (ln *HuffmanInternalNode) Right() HuffmanNode {
	return ln.right
}
