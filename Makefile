include $(GOROOT)/src/Make.inc

TARG = github.com/ziutek/gdk
CGOFILES = common.go event.go drawable.go window.go

include $(GOROOT)/src/Make.pkg
