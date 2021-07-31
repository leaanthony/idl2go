package webview2

import "golang.org/x/sys/windows"

// ComProc stores a COM procedure.
type ComProc uintptr

// NewComProc creates a new COM proc from a Go function.
func NewComProc(fn interface{}) ComProc {
	return ComProc(windows.NewCallback(fn))
}

// IUnknown
type _IUnknown struct {
	vtbl *_IUnknownVtbl
}

type _IUnknownVtbl struct {
	QueryInterface ComProc
	AddRef         ComProc
	Release        ComProc
}

type _IUnknownImpl interface {
	QueryInterface(refiid, object uintptr) uintptr
	AddRef() uintptr
	Release() uintptr
}
