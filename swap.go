package util9

import (
	"io"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Swap struct {
	Memory            uint64
	PageSize          uint64
	KernelPageTotal   uint64
	UserPageUsed      uint64
	UserPageTotal     uint64
	SwapPageUsed		uint64
	SwapPageTotal	uint64
	KernelMallocUsed  uint64
	KernelMallocTotal uint64
	KernelDrawUsed    uint64
	KernelDrawTotal   uint64
}

func (sw Swap) String() string {
	s := fmt.Sprintf("%d memory\n", sw.Memory)
	s += fmt.Sprintf("%d pagesize\n", sw.PageSize)
	s += fmt.Sprintf("%d kernel\n", sw.KernelPageTotal)
	s += fmt.Sprintf("%d/%d user\n", sw.UserPageUsed, sw.UserPageTotal)
	s += fmt.Sprintf("%d/%d swap\n", sw.SwapPageUsed, sw.SwapPageTotal)
	s += fmt.Sprintf("%d/%d kernel malloc\n", sw.KernelMallocUsed, sw.KernelMallocTotal)
	s += fmt.Sprintf("%d/%d kernel draw\n", sw.KernelDrawUsed, sw.KernelDrawTotal)
	return s
}

var mu sync.Mutex
var devswap io.ReaderAt
var devswaperror error

func init() {
	ds, err := os.Open("/dev/swap")
	if err != nil {
		devswaperror = err
		return
	}

	devswap = ds
}

func ReadSwap() (*Swap, error) {
	if devswap == nil {
		return nil, devswaperror
	}

	buf := make([]byte, 256)

	n, err := devswap.ReadAt(buf, 0)
	if n <= 0 && err != nil {
		return nil, err
	}

	buf = buf[:n]

	spl := strings.Split(string(buf), "\n")

	swp := &Swap{}

	for i, l := range spl {
		splnum := strings.Split(l, " ")
		if len(splnum) < 1 {
			continue
		}

		num := splnum[0]

		switch i {
		case 0:
			swp.Memory, _ = strconv.ParseUint(num, 10, 64)
		case 1:
			swp.PageSize, _ = strconv.ParseUint(num, 10, 64)
		case 2:
			swp.KernelPageTotal, _ = strconv.ParseUint(num, 10, 64)
		case 3:
			splused := strings.Split(num, "/")
			if len(splused) < 2 {
				break
			}
			
			swp.UserPageUsed, _ = strconv.ParseUint(splused[0], 10, 64)
			swp.UserPageTotal, _ = strconv.ParseUint(splused[1], 10, 64)
		case 4:
			splused := strings.Split(num, "/")
		fmt.Println(splused)
			if len(splused) < 2 {
				break
			}
			
			swp.SwapPageUsed, _ = strconv.ParseUint(splused[0], 10, 64)
			swp.SwapPageTotal, _ = strconv.ParseUint(splused[1], 10, 64)
		case 5:
			splused := strings.Split(num, "/")
		fmt.Println(splused)
			if len(splused) < 2 {
				break
			}
			
			swp.KernelMallocUsed, _ = strconv.ParseUint(splused[0], 10, 64)
			swp.KernelMallocTotal, _ = strconv.ParseUint(splused[1], 10, 64)
		case 6:
			splused := strings.Split(num, "/")
		fmt.Println(splused)
			if len(splused) < 2 {
				break
			}
			
			swp.KernelDrawUsed, _ = strconv.ParseUint(splused[0], 10, 64)
			swp.KernelDrawTotal, _ = strconv.ParseUint(splused[1], 10, 64)
		}
	}

	return swp, nil
}
