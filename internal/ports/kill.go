package ports

import (
	"fmt"
	"os/exec"
	"strconv"
)

func KillPort(port int) error {

	ports, err := GetPorts()

	if err != nil {
		return err
	}

	for _, p := range ports {

		if p.Port == port {

			fmt.Printf("\nKill %s on port %d (project: %s)? [y/N]: ",
				p.Service, p.Port, p.Project)

			var input string

			fmt.Scanln(&input)

			if input != "y" && input != "Y" {

				fmt.Println("Aborted")

				return nil
			}

			cmd := exec.Command("kill", "-9", strconv.Itoa(p.PID))

			return cmd.Run()
		}
	}

	return fmt.Errorf("no process found on port %d", port)
}
