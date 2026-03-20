package main

import (
        "fmt"
        "os"
        "strconv"

        "github.com/WAZULU503/Name-james-dev-tool/internal/ports"
        "github.com/WAZULU503/Name-james-dev-tool/internal/ui"
)

func main() {
        if len(os.Args) < 2 {
                printUsage()
                return
        }

        cmd := os.Args[1]

        if cmd == "dev" || cmd == "p" || cmd == "port" {
                if len(os.Args) < 3 {
                        printUsage()
                        return
                }
        }

        sub := ""
        if len(os.Args) >= 3 {
                sub = os.Args[2]
        }

        switch cmd {

        case "dev":
                switch sub {
                case "ls":
                        items, err := ports.GetPorts()
                        if err != nil || len(items) == 0 {
                                fmt.Println("No active ports.")
                                return
                        }
                        ui.PrintTable(items)

                case "clean":
                        err := ports.FreePorts()
                        if err != nil {
                                fmt.Println(err)
                                return
                        }

                        // VERIFY AFTER CLEAN
                        items, err := ports.GetPorts()
                        if err != nil || len(items) == 0 {
                                fmt.Println("✔ environment cleared (verified)")
                                return
                        }

                        fmt.Println("✖ environment not fully cleared")
                        ui.PrintTable(items)

                default:
                        printUsage()
                }

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
                        printUsage()
                }

        default:
                printUsage()
        }
}

func printUsage() {
        fmt.Println("jdt")
        fmt.Println("Usage:")
        fmt.Println("  jdt p ls       - List active ports")
        fmt.Println("  jdt p k <port> - Kill port")
        fmt.Println("  jdt p f        - Free common ports")
        fmt.Println("  jdt dev ls     - List dev ports")
        fmt.Println("  jdt dev clean  - Clean dev environment")
}
