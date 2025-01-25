package rbs

import (
	"bufio"
	"strings"
)

type Operation struct {
	Name   string
	Inputs []string
	Output string
}

type Wire struct {
	Inputs string
	Output string
}

type Circuit struct {
	Operations []Operation
	Directions string
	Wiring     Wire
	Inputs     []string
}

func parseOp(line string) Operation {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return Operation{}
	}
	domainStr := strings.Trim(fields[1], "<>")
	inputs := strings.Split(domainStr, ",")
	output := strings.TrimSpace(fields[2])

	return Operation{
		Name:   fields[0],
		Inputs: inputs,
		Output: output,
	}
}

func ParseCircuit(input string) Circuit {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var circuit Circuit
	var currentSection string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Skip empty lines and separator lines
		if line == "" || strings.Contains(line, "---") {
			continue
		}

		// Parse header line
		if strings.Contains(line, "Name") {
			continue
		}

		// Check for section headers
		if strings.Contains(line, "Directions") {
			currentSection = "directions"
		} else if strings.Contains(line, "Wiring") {
			currentSection = "wiring"
		} else if strings.Contains(line, "Inputs") {
			currentSection = "inputs"
		}

		// Parse content based on current section
		ws := strings.Split(afterDash(line), "~")
		switch currentSection {
		case "":
			op := parseOp(line)
			if op.Name != "" {
				circuit.Operations = append(circuit.Operations, op)
			}
		case "directions":
			circuit.Directions = strings.Split(afterDash(line), "~")[0]
		case "wiring":
			circuit.Wiring = Wire{
				Inputs: ws[0],
			}
			if len(ws) == 2 {
				circuit.Wiring.Output = ws[1]
			}
		case "inputs":
			inputsStr := strings.TrimSpace(afterDash(line))
			inputsStr = strings.TrimLeft(inputsStr, "(")
			circuit.Inputs = strings.Fields(inputsStr)
		}
	}

	return circuit
}

func afterDash(s string) string {
	parts := strings.Split(s, "-")
	return parts[len(parts)-1]
}
