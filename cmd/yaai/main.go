package main

// #include <stdio.h>
// #include <stdlib.h>
// #include <dlfcn.h>
import "C"

import (
	"fmt"

	"unsafe"

	"github.com/alecthomas/kingpin"
)

var (
	path *string = kingpin.Flag("path", "path to shared object").
		Required().
		String()
)

func main() {
	kingpin.Parse()

	fmt.Println(*path)
	getSharedThing(*path)
}

func getSharedThing(soPath string) error {
	path := C.CString(soPath)
	defer C.free(unsafe.Pointer(path))

	handle := C.dlopen(path, C.RTLD_LAZY)
	if handle == nil {
		return fmt.Errorf("could not open dynamic library %s", soPath)
	}

	sym := C.CString("SharedThing")
	defer C.free(unsafe.Pointer(sym))
	SharedThing := C.dlsym(handle, sym)
	if SharedThing == nil {
		return fmt.Errorf("couldn't find symbol 'SharedThing' in dynamic library %s", soPath)
	}

	return nil
}
