package model

type EmployeeStateUnaligned struct {
	Id        int64  // 8 bytes
	IsActive  bool   // 1 byte with 7 padding bytes
	Age       int64  // 8 bytes
	Name      string // 16 bytes
	IsMarried bool   // 1 byte
}

type EmployeeStateAligned struct {
	Id        int64  // 8 bytes
	Age       int64  // 8 bytes
	Name      string // 16 bytes
	IsActive  bool   // 1 byte
	IsMarried bool   // 1 byte
}
