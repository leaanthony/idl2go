package webview2

type _ICoreWebView2WebResourceResponseVtbl struct {
	_IUnknownVtbl
	GetContent      ComProc
	PutContent      ComProc
	GetHeaders      ComProc
	GetStatusCode   ComProc
	PutStatusCode   ComProc
	GetReasonPhrase ComProc
	PutReasonPhrase ComProc
}

type ICoreWebView2WebResourceResponse struct {
	vtbl *_ICoreWebView2WebResourceResponseVtbl
}
