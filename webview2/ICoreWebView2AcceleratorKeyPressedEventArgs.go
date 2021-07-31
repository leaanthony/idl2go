package webview2

type _ICoreWebView2AcceleratorKeyPressedEventArgsVtbl struct {
	_IUnknownVtbl
	GetKeyEventKind      ComProc
	GetVirtualKey        ComProc
	GetKeyEventLParam    ComProc
	GetPhysicalKeyStatus ComProc
	GetHandled           ComProc
	PutHandled           ComProc
}

type ICoreWebView2AcceleratorKeyPressedEventArgs struct {
	vtbl *_ICoreWebView2AcceleratorKeyPressedEventArgsVtbl
}
