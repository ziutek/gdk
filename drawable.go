package gdk

/*
#include <gdk/gdk.h>
#include <gdk/gdkx.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Drawable struct {
	glib.Object
}

func (d *Drawable) g() *C.GdkDrawable {
	return (*C.GdkDrawable)(d.GetPtr())
}

func (d *Drawable) AsDrawable() *Drawable {
		return d
}

func (d *Drawable) GetXid() uint {
	return uint(C.gdk_x11_drawable_get_xid(d.g()))
}
