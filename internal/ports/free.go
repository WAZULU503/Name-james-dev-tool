package ports

import (
	"fmt"
	"os/exec"
	"strconv"
)

func FreePorts() error {

	ports, err := GetPorts()

	if err != nil {
		return err
	}

	if len(ports) == 0 {

		fmt.Println("No active dev ports")

		return nil
	}

	fmt.Println("Active ports:")

	for _, p := range ports {

		fmt.Printf("%d %s %s\n", p.Port, p.Service, p.Project)
	}

	fmt.Print("\nKill all? [y/N]: ")

	var input string

	fmt.Scanln(&input)

	if input != "y" && input != "Y" {
		return nil
	}

	for _, p := range ports {

		exec.Command("kill", "-9", strconv.Itoa(p.PID)).Run()
	}

	fmt.Println("Environment cleared")

	return nil
}
