package brain

type Node struct {
	num       float32
	linksN    []*Node
	linksO    []*Output
	fired     bool
	infs      int // Influences
	weights   []float32
	lastLayer bool
}

func newNode(numLinks int, isLast bool) Node {
	n := Node{}

	n.lastLayer = isLast

	if isLast {
		n.linksO = make([]*Output, numLinks)
	} else {
		n.linksN = make([]*Node, numLinks)
	}

	n.weights = make([]float32, numLinks)

	return n
}

// Calculate output
func (node *Node) push() {
	if node.num > 0 {
		node.fired = true
	}
	if node.lastLayer {
		for i := 0; i < len(node.linksO); i++ {
			node.linksO[i].num += node.weights[i] * node.num
			node.linksO[i].infs++
		}
	} else {
		for i := 0; i < len(node.linksN); i++ {
			node.linksN[i].num += node.weights[i] * node.num
			node.linksN[i].infs++
		}
	}
}

// Reset to default
func (node *Node) set() {
	node.num = 0
	node.infs = 0
	node.fired = false
}

func (node *Node) randomize() {
	for i := 0; i < len(node.weights); i++ {
		node.weights[i] = randWeight()
	}
}
