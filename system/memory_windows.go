package system

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetDiskFreeSpaceExW = modkernel32.NewProc("GetDiskFreeSpaceExW")
)

type memoryStatusEx struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}

func getMemoryInfo() (*MemoryInfo, error) {
	var memInfo memoryStatusEx
	memInfo.Length = uint32(unsafe.Sizeof(memInfo))

	ret, _, err := modkernel32.NewProc("GlobalMemoryStatusEx").Call(uintptr(unsafe.Pointer(&memInfo)))
	if ret == 0 {
		return nil, fmt.Errorf("failed to get memory information: %w", err)
	}

	if memInfo.TotalPhys == 0 {
		return nil, fmt.Errorf("failed to get memory information")
	}

	return &MemoryInfo{
		Total:     memInfo.TotalPhys,
		Free:      memInfo.AvailPhys,
		Available: memInfo.AvailPhys,
	}, nil
}

func getDiskInfo(path string) (*DiskInfo, error) {
	var freeBytesAvailable, totalBytes, totalFreeBytes int64

	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	ret, _, err := procGetDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(pathPtr)),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&totalFreeBytes)),
	)

	if ret == 0 {
		return nil, fmt.Errorf("failed to get disk information: %v", err)
	}

	return &DiskInfo{
		Total: uint64(totalBytes),
		Free:  uint64(totalFreeBytes),
		Used:  uint64(totalBytes - totalFreeBytes),
	}, nil
}
