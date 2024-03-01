package memory

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stats struct {
	MemTotal,
	MemUsed,
	MemFree,
	MemAvailable,
	Buffers,
	Cached,
	Active,
	Inactive,
	SwapCache,
	SwapFree,
	SwapTotal,
	SwapUsed,
	Slab,
	PageTotal,
	Commited_AS,
	VmAllocUsed float64
	Enabed bool
}

func GetStats() (*Stats, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return getMemStats(file)
}

func ConvertMemVal(val float64) (float64, string) {
	if val < 1024*1024 {
		return val / 1024, "KB"
	} else {
		return val / (1024 * 1024), "MB"
	}
}

func getMemStats(out io.Reader) (*Stats, error) {
	scanner := bufio.NewScanner(out)
	var mem Stats

	memStats := map[string]*float64{
		"MemTotal":     &mem.MemTotal,
		"MemFree":      &mem.MemFree,
		"MemAvailable": &mem.MemAvailable,
		"Buffers":      &mem.Buffers,
		"Cached":       &mem.Cached,
		"Active":       &mem.Active,
		"Inactive":     &mem.Inactive,
		"SwapCache":    &mem.SwapCache,
		"SwapFree":     &mem.SwapFree,
		"SwapTotal":    &mem.SwapTotal,
		"Slab":         &mem.Slab,
		"PageTotal":    &mem.PageTotal,
		"Commited_AS":  &mem.Commited_AS,
		"VmAllocUsed":  &mem.VmAllocUsed,
	}

	for scanner.Scan() {
		res := scanner.Text()
		i := strings.IndexRune(res, ':')
		if i < 0 {
			continue
		}
		index := res[:i]
		if ptr := memStats[index]; ptr != nil {
			val := strings.TrimSpace(strings.TrimRight(res[i+1:], "kB"))
			if v, err := strconv.ParseFloat(val, 64); err == nil {
				*ptr = v * 1024
			}
			if index == "MemAvailable" {
				mem.Enabed = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning /proc/meminfo: %s", err)
	}

	mem.SwapUsed = mem.SwapTotal - mem.SwapFree

	if mem.Enabed {
		mem.MemUsed = mem.MemTotal - mem.MemAvailable
	} else {
		mem.MemUsed = mem.MemTotal - mem.MemFree - mem.Buffers - mem.Cached
	}

	return &mem, nil
}
