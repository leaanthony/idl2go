package webview2

type _ICoreWebView2HttpResponseHeadersVtbl struct {
	_IUnknownVtbl
	AppendHeader ComProc
	Contains     ComProc
	GetHeader    ComProc
	GetHeaders   ComProc
	GetIterator  ComProc
}

type ICoreWebView2HttpResponseHeaders struct {
	vtbl *_ICoreWebView2HttpResponseHeadersVtbl
}
