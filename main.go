package main

import (
	"SyMon/cpu"
	"SyMon/disk"
	"SyMon/memory"
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		// Gets cpu statistics from /cpu/cpu.go
		cpu, err := cpu.GetStats()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		// Calls the GetCpuPercent function
		cpuUsage := cpu.GetCpuPercent()

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

		// Clears the console to refresh output
		fmt.Print("\033[H\033[2J")

		for _, stat := range diskStats {
			fmt.Printf("Disk: %s, Reads: %d Writes: %d\n", stat.Name, stat.Reads, stat.Writes)
		}

		fmt.Printf("memory total: %d %s\n", convertedMemTotal, unit)
		fmt.Printf("memory used: %d %s\n", convertedMemUsed, unit)
		fmt.Printf("memory free: %d %s\n", convertedMemFree, unit)
		fmt.Printf("cpu usage: %f%%\n", cpuUsage)

		time.Sleep(2 * time.Second)
	}
}
