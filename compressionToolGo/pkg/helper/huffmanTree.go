package helper

type HuffmanNode interface {
	Leaf() bool
	Weight() int
}

type HuffmanLeafNode struct {
	element rune
	weight  int
}

func InitLeafNode(r rune, w int) *HuffmanLeafNode{
	return &HuffmanLeafNode{element: r, weight: w}
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
	leftEdge int
	rightEdge int
}

func InitInternalNode(w int, l HuffmanNode, r HuffmanNode ) *HuffmanInternalNode{
	return &HuffmanInternalNode{weight: w, left: l, right: r}
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

func (ln * HuffmanInternalNode) LeftEdge() int{
	return ln.leftEdge
}

func (ln * HuffmanInternalNode) RightEdge() int{
	return ln.rightEdge
}

func (ln * HuffmanInternalNode) SetLeftEdge(val int){
	ln.leftEdge = val
}

func (ln * HuffmanInternalNode) SetRightEdge(val int){
	ln.rightEdge = val
}