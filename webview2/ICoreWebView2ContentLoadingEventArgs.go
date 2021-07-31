package webview2

type _ICoreWebView2ContentLoadingEventArgsVtbl struct {
	_IUnknownVtbl
	GetIsErrorPage  ComProc
	GetNavigationId ComProc
}

type ICoreWebView2ContentLoadingEventArgs struct {
	vtbl *_ICoreWebView2ContentLoadingEventArgsVtbl
}
