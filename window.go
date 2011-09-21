package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkx.h>
*/
import "C"

type Window struct {
	Drawable
}

func (w *Window) g() *C.GdkWindow {
	return (*C.GdkWindow)(w.GetPtr())
}

func (w *Window) AsWindow() *Window {
		return w
}
