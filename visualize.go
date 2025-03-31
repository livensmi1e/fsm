package fsm

import (
	"fmt"
	"strings"
)

func DrawMermaid(m *machine) string {
	var sb strings.Builder
	sb.WriteString("stateDiagram-v2\n")
	indent := "    "
	sb.WriteString(indent + fmt.Sprintf("[*] --> %s\n", m.initial))
	for src, data := range m.states {
		for event, dst := range data.transitions {
			sb.WriteString(indent + fmt.Sprintf("%s --> %s : %s\n", src, dst, event))
		}
	}
	return sb.String()
}
