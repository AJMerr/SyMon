package main

import (
	"SyMon/cpu"
	"SyMon/disk"
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

	diskStats, err := disk.GetStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	for _, stat := range diskStats {
		fmt.Printf("Disk: %s, Reads: %d Writes: %d\n", stat.Name, stat.Reads, stat.Writes)
	}

	convertedMemTotal, unit := memory.ConvertMemVal(mem.MemTotal)
	fmt.Printf("memory total: %d %s\n", convertedMemTotal, unit)
	fmt.Printf("memory used: %d bytes\n", mem.MemUsed)
	fmt.Printf("cpu total: %d bytes\n", cpu.Total)
}
