package disk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stats struct {
	Name   string
	Reads  uint64
	Writes uint64
}

func GetStats() ([]Stats, error) {
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return getDiskStats(file)
}

func getDiskStats(out io.Reader) ([]Stats, error) {
	scanner := bufio.NewScanner(out)

	var diskStats []Stats

	for scanner.Scan() {
		index := strings.Fields(scanner.Text())
		if len(index) < 14 {
			continue
		}
		name := index[2]
		reads, err := strconv.ParseUint(index[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse %s reads of: ", name)
		}
		writes, err := strconv.ParseUint(index[7], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse %s writes of: ", name)
		}
		diskStats = append(diskStats, Stats{
			Name:   name,
			Reads:  reads,
			Writes: writes,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning /proc/diskstats %s", err)
	}
	return diskStats, nil
}
