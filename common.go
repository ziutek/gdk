package gdk

/*
#include <gdk/gdk.h>

#cgo pkg-config: gtk+-2.0
*/
import "C"

func ThreadsEnter() {
	C.gdk_threads_enter()
}

func ThreadsLeave() {
	C.gdk_threads_leave()
}

func init() {
	C.gdk_threads_init()
}
