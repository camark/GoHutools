//go:build !windows

package system

import (
	"fmt"
	"syscall"
)

func getMemoryInfo() (*MemoryInfo, error) {
	var info syscall.Sysinfo_t
	if err := syscall.Sysinfo(&info); err != nil {
		return nil, fmt.Errorf("failed to get memory information: %w", err)
	}

	// Convert to bytes (values are in mem_unit units)
	total := info.Totalram * uint64(info.Unit)
	free := info.Freeram * uint64(info.Unit)
	available := free + info.Bufferram*uint64(info.Unit)

	return &MemoryInfo{
		Total:     total,
		Free:      free,
		Available: available,
	}, nil
}

func getDiskInfo(path string) (*DiskInfo, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return nil, fmt.Errorf("failed to get disk information: %w", err)
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	available := stat.Bavail * uint64(stat.Bsize)

	return &DiskInfo{
		Total: total,
		Free:  available,
		Used:  total - free,
	}, nil
}
