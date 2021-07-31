package webview2

type _ICoreWebView2HttpHeadersCollectionIteratorVtbl struct {
	_IUnknownVtbl
	GetCurrentHeader    ComProc
	GetHasCurrentHeader ComProc
	MoveNext            ComProc
}

type ICoreWebView2HttpHeadersCollectionIterator struct {
	vtbl *_ICoreWebView2HttpHeadersCollectionIteratorVtbl
}
