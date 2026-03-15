# jdt (James Dev Tool)

Fast CLI tool for finding and killing blocked development ports.

When working on multiple local projects, ports like **3000**, **5173**, **8000**, or **8080** often remain open after a development server stops.

`jdt` quickly shows which process owns the port and which project started it, allowing you to terminate it instantly.

---

## Install

Requires Go.

Install directly from GitHub:

go install github.com/WAZULU503/Name-james-dev-tool/cmd/jdt@latest

After installation you can run:

jdt version

---

## Commands

### List active development ports

jdt p ls

Example output:

PORT   SERVICE            PROJECT                   PID
------------------------------------------------------------
3000   Next.js?           storefront                4212
5173   Vite?              admin-dashboard           8421
8000   Dev Server?        api-server                1331


### Kill a specific port

jdt p k 3000

You will be asked for confirmation before the process is terminated.

---

### Free all development ports

jdt p f

Lists active development ports and allows clearing them all at once.

---

### Show version

jdt version

---

## What jdt does

• Detects listening ports using `lsof`  
• Maps ports to running processes  
• Identifies the project directory that started the process  
• Allows safe termination with confirmation  

Works on:

• macOS  
• Linux  

---

## Example workflow

A development port is blocked:

Error: address already in use :3000

Run:

jdt p ls

Find the process:

PORT   SERVICE            PROJECT                   PID
------------------------------------------------------------
3000   Next.js?           marketing-site            4212

Kill it:

jdt p k 3000

Port is now free.

---

## Design Philosophy

Small tools that solve real developer problems.

No frameworks  
No telemetry  
No background services  

Just a single fast Go binary.

---

## Author

Wazulu  
aka James Redmond

---

## License

MIT License
