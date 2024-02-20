package main

import (
	"SyMon/cpu"
	"SyMon/memory"
	"fmt"
	"os"
)

func main() {
	cpu, err := cpu.GetStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	mem, err := memory.GetStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	fmt.Printf("memory total: %d bytes\n", mem.MemTotal)
	fmt.Printf("memory used: %d bytes\n", mem.MemUsed)
	fmt.Printf("cpu total: %d bytes\n", cpu.Total)
}
