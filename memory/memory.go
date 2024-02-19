package memory

//import (
// "fmt"
//"os"
// "io"
//"bufio"
//"strconv"
//"strings"
//)

type Stats struct {
	MemToal,
	MemFree,
	MemAvailable,
	Buffers,
	Cached,
	Active,
	Inactive,
	SwapCache,
	SwapFree,
	SwapTotal,
	Slab,
	PageTotal,
	Commited_AS,
	VmAllocUsed uint64
}
