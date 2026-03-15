package ports

import (
	"bufio"
	"bytes"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var projectCache = map[int]string{}

func GetPorts() ([]PortInfo, error) {

	cmd := exec.Command("lsof", "-nP", "-iTCP", "-sTCP:LISTEN", "-F")

	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return parseLsof(out), nil
}

func parseLsof(data []byte) []PortInfo {

	var results []PortInfo
	seen := map[int]bool{}

	scanner := bufio.NewScanner(bytes.NewReader(data))

	var pid int
	var cmd string

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		switch line[0] {

		case 'p':
			pid, _ = strconv.Atoi(line[1:])

		case 'c':
			cmd = line[1:]

		case 'n':

			addr := line[1:]

			if strings.Contains(addr, ":") {

				pStr := addr[strings.LastIndex(addr, ":")+1:]

				port, _ := strconv.Atoi(pStr)

				if seen[port] {
					continue
				}

				seen[port] = true

				project := detectProject(pid)

				service, guess := detectService(port)

				results = append(results, PortInfo{
					Port:    port,
					PID:     pid,
					Process: cmd,
					Service: service,
					Project: project,
					Guess:   guess,
				})
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})

	return results
}

func detectProject(pid int) string {

	if name, ok := projectCache[pid]; ok {
		return name
	}

	cmd := exec.Command("lsof", "-p", strconv.Itoa(pid), "-a", "-d", "cwd", "-Fn")

	out, err := cmd.Output()

	if err != nil {
		return "unknown"
	}

	lines := strings.Split(string(out), "\n")

	for _, l := range lines {

		if strings.HasPrefix(l, "n") {

			path := l[1:]

			name := filepath.Base(path)

			projectCache[pid] = name

			return name
		}
	}

	return "unknown"
}

func detectService(port int) (string, bool) {

	if name, ok := serviceMap[port]; ok {
		return name + "?", true
	}

	return "Unknown", false
}
