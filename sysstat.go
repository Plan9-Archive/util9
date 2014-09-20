package util9

import (
	"io"
	"os"
	"strconv"
	"strings"
)

type Sysstat struct {
	Cpu             uint64
	ContextSwitches uint64
	Interrupts      uint64
	SystemCalls     uint64
	PageFaults      uint64
	TLBFaults       uint64
	TLBPurges       uint64
	LoadAverage     uint64
	IdleTime        uint64
	InterruptTime   uint64
}

var devsysstat io.ReaderAt
var devsysstaterror error

func init() {
	dstat, err := os.Open("/dev/sysstat")
	if err != nil {
		devsysstaterror = err
		return
	}

	devsysstat = dstat
}

func ReadSysstat() ([]Sysstat, error) {
	if devsysstat == nil {
		return nil, devsysstaterror
	}

	buf := make([]byte, 8192)

	n, err := devsysstat.ReadAt(buf, 0)
	if n <= 0 && err != nil {
		return nil, err
	}

	buf = buf[:n]

	spl := strings.Split(string(buf), "\n")

	var stats []Sysstat

	for _, l := range spl {
		splnum := strings.Fields(l)
		if len(splnum) != 10 {
			continue
		}

		stat := Sysstat{}

		stat.Cpu = atoi(splnum[0])
		stat.ContextSwitches = atoi(splnum[1])
		stat.Interrupts = atoi(splnum[2])
		stat.SystemCalls = atoi(splnum[3])
		stat.PageFaults = atoi(splnum[4])
		stat.TLBFaults = atoi(splnum[5])
		stat.TLBPurges = atoi(splnum[6])
		stat.LoadAverage = atoi(splnum[7])
		stat.IdleTime = atoi(splnum[8])
		stat.InterruptTime = atoi(splnum[9])

		stats = append(stats, stat)
	}

	return stats, nil
}

func atoi(n string) uint64 {
	v, _ := strconv.ParseUint(n, 10, 64)
	return v
}
