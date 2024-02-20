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
	kB := mem.MemTotal / 1024
	mB := kB / 1024
	fmt.Printf("memory total: %d bytes\n", mB)
	fmt.Printf("memory used: %d bytes\n", kB)
	fmt.Printf("cpu total: %d bytes\n", cpu.Total)
}
