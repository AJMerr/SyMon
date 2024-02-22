package main

import (
	"SyMon/cpu"
	"SyMon/disk"
	"SyMon/memory"
	"fmt"
	"os"
)

func main() {
	// Gets cpu statistics from /cpu/cpu.go
	cpu, err := cpu.GetStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	// Gets memory statistics from /memory/memory.go
	mem, err := memory.GetStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	// Converts the bytes to KB or MB respectively for each memory stat
	convertedMemTotal, unit := memory.ConvertMemVal(mem.MemTotal)
	convertedMemUsed, unit := memory.ConvertMemVal(mem.MemUsed)
	convertedMemFree, unit := memory.ConvertMemVal(mem.MemFree)

	// Gets disk statistics from /disk/disk.go
	diskStats, err := disk.GetStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	for _, stat := range diskStats {
		fmt.Printf("Disk: %s, Reads: %d Writes: %d\n", stat.Name, stat.Reads, stat.Writes)
	}

	fmt.Printf("memory total: %d %s\n", convertedMemTotal, unit)
	fmt.Printf("memory used: %d %s\n", convertedMemUsed, unit)
	fmt.Printf("memory free: %d %s\n", convertedMemFree, unit)
	fmt.Printf("cpu total: %d\n", cpu.Total)
}
