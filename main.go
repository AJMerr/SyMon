package main

import (
	"SyMon/cpu"
	"fmt"
	"os"
)

func main() {
	cpu, err := cpu.GetStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Printf("CPU Total: %d bytes\n", cpu.Total)
}
