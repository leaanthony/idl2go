package webview2

type _ICoreWebView2MoveFocusRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetReason  ComProc
	GetHandled ComProc
	PutHandled ComProc
}

type ICoreWebView2MoveFocusRequestedEventArgs struct {
	vtbl *_ICoreWebView2MoveFocusRequestedEventArgsVtbl
}
