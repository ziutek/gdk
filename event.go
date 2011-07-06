package gdk

/*
#include <gdk/gdk.h>

#cgo pkg-config: gtk+-2.0
*/
import "C"

import (
	"github.com/ziutek/glib"
)


type Event C.GdkEvent

func (e *Event) Type() glib.Type {
	return glib.Type(C.gdk_event_get_type())
}
