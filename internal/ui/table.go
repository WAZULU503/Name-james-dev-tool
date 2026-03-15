package ui

import (
	"fmt"
	"strings"

	"github.com/YOURUSERNAME/james-dev-tool/internal/ports"
)

func PrintTable(items []ports.PortInfo) {

	if len(items) == 0 {

		fmt.Println("No common dev ports detected.")

		return
	}

	fmt.Printf("%-6s %-18s %-25s %-8s\n",
		"PORT", "SERVICE", "PROJECT", "PID")

	fmt.Println(strings.Repeat("-", 60))

	for _, p := range items {

		fmt.Printf("%-6d %-18s %-25s %-8d\n",
			p.Port,
			p.Service,
			p.Project,
			p.PID,
		)
	}
}
