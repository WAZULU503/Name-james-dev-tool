package ports

type PortInfo struct {
	Port    int
	PID     int
	Process string
	Service string
	Project string
	Guess   bool
}

var serviceMap = map[int]string{
	3000: "Next.js",
	5173: "Vite",
	8000: "Dev Server",
	3306: "MySQL",
	6379: "Redis",
}
