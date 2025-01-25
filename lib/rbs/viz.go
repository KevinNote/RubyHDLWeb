package rbs

import (
	"fmt"
)

func VizToDot(circuit Circuit) string {
	dot := "digraph circuit {\n"
	dot += "\tnode [shape=box];\n"

	// Add operation nodes
	for i, op := range circuit.Operations {
		dot += fmt.Sprintf("\top%d [style=filled,fillcolor=khaki1,label=\"%s\"];\n", i, op.Name)
		// Add edges from inputs to operation
		for _, input := range op.Inputs {
			dot += fmt.Sprintf("\t\"%s\" -> op%d;\n", input, i)
		}
		// Add edge from operation to output
		dot += fmt.Sprintf("\top%d -> \"%s\";\n", i, op.Output)
	}

	// Add input nodes with different shape
	dot += "\tnode [shape=ellipse];\n"
	for _, input := range circuit.Inputs {
		dot += fmt.Sprintf("\t\"%s\" [style=filled,fillcolor=lightblue];\n", input)
	}

	dot += "}\n"
	return dot
}
