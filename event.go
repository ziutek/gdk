package gdk

/*
#include <gdk/gdk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)


type Event C.GdkEvent

func (e *Event) Type() glib.Type {
	return glib.Type(C.gdk_event_get_type())
}

