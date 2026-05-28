package brain

type Input struct {
	num     float32
	links   []*Node
	weights []float32
}

func newInput(numLinks int) Input {
	i := Input{}

	i.weights = make([]float32, numLinks)
	i.links = make([]*Node, numLinks)

	return i
}

// Calculate output
func (input *Input) push() {
	for i := 0; i < len(input.links); i++ {
		input.links[i].num += input.num * input.weights[i]
		input.links[i].infs++
	}
}

// Reset to default
func (input *Input) set() {
	input.num = 0
}

func (input *Input) randomize() {
	for i := 0; i < len(input.weights); i++ {
		input.weights[i] = randWeight()
	}
}
