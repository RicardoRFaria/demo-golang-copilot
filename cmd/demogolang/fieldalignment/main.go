package main

import (
	"demo-golang/model"
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("Size in bytes: " + strconv.FormatUint(uint64(unsafe.Sizeof(model.EmployeeStateUnaligned{})), 10)) // Output: 40
	fmt.Println("Size in bytes: " + strconv.FormatUint(uint64(unsafe.Sizeof(model.EmployeeStateAligned{})), 10))   // Output: 32

	sizeInBytes := uint64(unsafe.Sizeof(model.EmployeeStateUnaligned{})) * 15000
	sizeInMegabytes := bytesToMegabytes(sizeInBytes)
	fmt.Println("Size in megabytes with 15k of UNALIGNED: " + strconv.FormatFloat(sizeInMegabytes, 'f', 2, 64))

	sizeInBytes = uint64(unsafe.Sizeof(model.EmployeeStateAligned{})) * 15000
	sizeInMegabytes = bytesToMegabytes(sizeInBytes)
	fmt.Println("Size in megabytes with 15k of ALIGNED: " + strconv.FormatFloat(sizeInMegabytes, 'f', 2, 64))

}

func bytesToMegabytes(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024
}
