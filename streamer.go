package streamer

/*
#cgo windows CFLAGS: -I./lib -I./lib/gdk -I./lib/sdk -I./lib/sdk/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo windows CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DWIN32
#cgo windows LDFLAGS: -Wl,--subsystem,windows,--kill-at

#cgo linux CFLAGS: -I./lib -I./lib/gdk -I./lib/sdk -I./lib/sdk/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#include "unitybuild.c"
#include "sampgdk.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func findNative(name string) (C.AMX_NATIVE, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	native := C.sampgdk_FindNative(cName)
	if native == nil {
		return nil, fmt.Errorf("native could not be found")
	}

	return native, nil
}

func invokeNative(native C.AMX_NATIVE, types string, args ...interface{}) int {
	return C.sampgdk_InvokeNative(native, types, args)
}
