package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "status", "--porcelain", "-b")
	out, err := cmd.Output()
	if err == nil {
		stats := strings.Split(string(out), "\n")
		var brname string
		var br, add, unt bool
		for _, s := range stats {
			if s == "" {
				continue
			}
			switch {
			case s[:3] == "## ":
				br = true
			case strings.HasPrefix(s, "A"):
				add = true
			case strings.HasPrefix(s, "?"):
				unt = true

			}
			if br {
				brname = strings.Split(s, " ")[1]
				brname = strings.Split(brname, ".")[0]
				br = false
			}
		}
		fmt.Printf("%s %t %t\n", brname, add, unt)
	}
}
