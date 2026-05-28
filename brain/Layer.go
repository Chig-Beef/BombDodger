package brain

type Layer struct {
	nodes []Node
}

func newLayer(depth int, nextDepth int, isLast bool) Layer {
	l := Layer{}

	l.nodes = make([]Node, depth)
	for i := 0; i < depth; i++ {
		l.nodes[i] = newNode(nextDepth, isLast)
	}

	return l
}

// Reset to default
func (layer *Layer) set() {
	for i := 0; i < len(layer.nodes); i++ {
		layer.nodes[i].set()
	}
}

// Calculate output
func (layer *Layer) push() {
	for i := 0; i < len(layer.nodes); i++ {
		layer.nodes[i].push()
	}
}

func (layer *Layer) randomize() {
	for i := 0; i < len(layer.nodes); i++ {
		layer.nodes[i].randomize()
	}
}
