package cpu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func GetStats() (*Stats, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return getCpuStats(file)
}

type Stats struct {
	User,
	Nice,
	System,
	Idle,
	Iowait,
	Irq,
	Softirq,
	Steal,
	Guest,
	GuestNice,
	Total uint64
	CPUCount,
	StatCount int
}

type cpuStat struct {
	name string
	ptr  *uint64
}

func getCpuStats(out io.Reader) (*Stats, error) {
	scanner := bufio.NewScanner(out)
	var cpu Stats

	cpuStats := []cpuStat{
		{"User", &cpu.User},
		{"Nice", &cpu.Nice},
		{"System", &cpu.System},
		{"Idle", &cpu.Idle},
		{"Iowait", &cpu.Iowait},
		{"Irq", &cpu.Irq},
		{"Softirq", &cpu.Softirq},
		{"Steal", &cpu.Steal},
		{"Guest", &cpu.Guest},
		{"GuestNice", &cpu.GuestNice},
	}

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to scan /proc/stat")
	}

	valStrs := strings.Fields(scanner.Text())[1:]
	cpu.StatCount = len(valStrs)

	for i, valStr := range valStrs {
		val, err := strconv.ParseUint(valStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to scan %s from /proc/stat", cpuStats[i].name)
		}
		*cpuStats[i].ptr = val
		cpu.Total += val
	}

	cpu.Total -= cpu.Guest
	cpu.Total -= cpu.GuestNice

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "cpu") && unicode.IsDigit(rune(line[3])) {
			cpu.CPUCount++
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("scan error for /proc/stat: %s", err)
		}
	}

	return &cpu, nil
}

func (s *Stats) GetCpuPercent() float64 {
	totalNonIdle := s.User + s.System + s.Nice + s.Irq + s.Softirq + s.Steal
	totalSystem := s.System + s.Irq + s.Softirq
	totalUsage := totalNonIdle + totalSystem

	return 100 * float64(totalSystem) / float64(totalUsage)
}
