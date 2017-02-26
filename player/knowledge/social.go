package player

type Node struct {
	ID        int
	Neighbors []int
}

type Social struct {
	graph []Node `json:"graph"`
}
