package main

import (
	"SyMon/cpu"
	"fmt"
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	// Adding Termui for better visualization
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to init termui: %v", err)
	}
	defer ui.Close()

	// Variable that takes in a pointer the the Stats struct from the CPU package
	var prevCpuStats *cpu.Stats

	g1 := widgets.NewGauge()
	g1.Title = "CPU Usage:"
	g1.SetRect(0, 3, 50, 6)
	g1.BarColor = ui.ColorGreen
	g1.LabelStyle = ui.NewStyle(ui.ColorWhite)
	g1.BorderStyle.Fg = ui.ColorWhite

	ui.Render(g1)
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(2 * time.Second).C

MainLoop:
	for {
		select {
		case <-ticker:
			// Gets cpu statistics from /cpu/cpu.go
			currentCpuStats, err := cpu.GetStats()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				return
			}

			var cpuUsage float64
			// Calls the CalcCpuPercent function if prevCpuStats are empty and sets prevCpuStats to currentCpuStats for comparison
			if prevCpuStats != nil {
				cpuUsage = cpu.CalcCpuPercent(prevCpuStats, currentCpuStats)
			}
			prevCpuStats = currentCpuStats

			// Update the CPU gauge
			g1.Percent = int(cpuUsage)
			ui.Render(g1)

		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				break MainLoop
			}
		}
	}
}
