package webview2

type _ICoreWebView2DeferralVtbl struct {
	_IUnknownVtbl
	Complete ComProc
}

type ICoreWebView2Deferral struct {
	vtbl *_ICoreWebView2DeferralVtbl
}
