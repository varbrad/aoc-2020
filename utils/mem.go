package utils

import (
	"fmt"
	"runtime"

	"github.com/dustin/go-humanize"
)

// LogMemoryUsage helper to visualize memory usage
func LogMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v", humanize.Bytes(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v", humanize.Bytes(m.TotalAlloc))
	fmt.Printf("\tSys = %v\n", humanize.Bytes(m.Sys))
}
