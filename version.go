package levigo

/*
#cgo LDFLAGS: -lrocksdb
#include "rocksdb/c.h"
*/
import "C"

func GetLevelDBMajorVersion() int {
	return int(C.rocksdb_major_version())
}

func GetLevelDBMinorVersion() int {
	return int(C.rocksdb_minor_version())
}
