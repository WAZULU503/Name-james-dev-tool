package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/YOURUSERNAME/james-dev-tool/internal/ports"
	"github.com/YOURUSERNAME/james-dev-tool/internal/ui"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  jdt p ls")
		fmt.Println("  jdt p k <port>")
		fmt.Println("  jdt p f")
		return
	}

	cmd := os.Args[1]
	sub := os.Args[2]

	switch cmd {

	case "p", "port":

		switch sub {

		case "ls", "list":

			items, err := ports.GetPorts()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			ui.PrintTable(items)

		case "k", "kill":

			if len(os.Args) < 4 {
				fmt.Println("Specify port")
				return
			}

			port, _ := strconv.Atoi(os.Args[3])
			err := ports.KillPort(port)

			if err != nil {
				fmt.Println(err)
			}

		case "f", "free":

			err := ports.FreePorts()

			if err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Println("Unknown port command")
		}

	default:
		fmt.Println("Unknown command")
	}
}
