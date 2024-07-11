package main

// InitializeTree takes a slice of n species names as input.
// It returns a rooted binary tree with with 2n-1 total nodes, where
// the leaves are the first n and have the associated species names.
func InitializeTree(speciesNames []string) Tree {
	var t Tree

	numLeaves := len(speciesNames)

	if numLeaves == 0 {
		panic("Error: no speices given")
	}

	t = make([]Node, 2*numLeaves-1)

	for i := range t {
		//decleare a node
		var vx Node

		t[i] = &vx
	}

	return t
}
