package main

import (
	"SyMon/cpu"
	"SyMon/disk"
	"SyMon/memory"
	"fmt"
	"log"
	"os"
	"strconv"
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

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(2 * time.Second).C

MainLoop:
	for {
		select {
		case <-ticker:
			// Gets cpu statistics from CPU module
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

			// Gets memory statistics from Memory module
			mem, err := memory.GetStats()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				return
			}
			memTotal, unit := memory.ConvertMemVal(mem.MemTotal)
			memUsed, unit := memory.ConvertMemVal(mem.MemUsed)
			memFree, unit := memory.ConvertMemVal(mem.MemFree)

			memTotalToString := strconv.Itoa(int(memTotal))
			memUsedToString := strconv.Itoa(int(memUsed))
			memFreeToString := strconv.Itoa(int(memFree))

			p1 := widgets.NewParagraph()
			p1.Title = "Total Memory:"
			p1.BorderStyle.Fg = ui.ColorWhite
			p1.Text = memTotalToString + unit

			p2 := widgets.NewParagraph()
			p2.Title = "Memory Used:"
			p2.BorderStyle.Fg = ui.ColorWhite
			p2.Text = memUsedToString + unit

			p3 := widgets.NewParagraph()
			p3.Title = "Memory Free:"
			p3.BorderStyle.Fg = ui.ColorWhite
			p3.Text = memFreeToString + unit

			// Get stats from Disk module.
			diskStats, err := disk.GetStats()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}

			// Setting up TermUI for Disk Stats
			p4 := widgets.NewParagraph()
			p4.BorderStyle.Fg = ui.ColorWhite

			p5 := widgets.NewParagraph()
			p5.BorderStyle.Fg = ui.ColorWhite

			for _, stat := range diskStats {
				dName := stat.Name
				dReads := stat.Reads
				dWrites := stat.Writes

				dReadsToString := strconv.Itoa(int(dReads))
				dWritesToString := strconv.Itoa(int(dWrites))

				p4.Title = dName
				p4.Text = dReadsToString

				p5.Title = dName
				p5.Text = dWritesToString
			}

			grid := ui.NewGrid()
			termWidth, termHeight := ui.TerminalDimensions()
			grid.SetRect(0, 0, termWidth, termHeight)

			grid.Set(
				ui.NewRow(.5/4, g1),
				ui.NewRow(1.0/5,
					ui.NewCol(1.0/5, p1),
					ui.NewCol(1.0/5, p2),
					ui.NewCol(1.0/5, p3)),
				ui.NewRow(1.0/5,
					ui.NewCol(1.0/5, p4),
					ui.NewCol(1.0/5, p5)),
			)
			ui.Render(grid)

		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				break MainLoop
			}
		}
	}
}
