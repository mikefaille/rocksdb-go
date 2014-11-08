package levigo

/*
#cgo LDFLAGS: -lrocksdb
#include "rocksdb/c.h"
*/
import "C"

func GetRocksDBMajorVersion() int {
	return int(C.rocksdb_major_version())
}

func GetRocksDBMinorVersion() int {
	return int(C.rocksdb_minor_version())
}
