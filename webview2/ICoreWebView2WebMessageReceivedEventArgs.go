package webview2

type _ICoreWebView2WebMessageReceivedEventArgsVtbl struct {
	_IUnknownVtbl
	GetSource                ComProc
	GetWebMessageAsJson      ComProc
	TryGetWebMessageAsString ComProc
}

type ICoreWebView2WebMessageReceivedEventArgs struct {
	vtbl *_ICoreWebView2WebMessageReceivedEventArgsVtbl
}
