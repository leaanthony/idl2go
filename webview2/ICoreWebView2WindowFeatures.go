package webview2

type _ICoreWebView2WindowFeaturesVtbl struct {
	_IUnknownVtbl
	GetHasPosition             ComProc
	GetHasSize                 ComProc
	GetLeft                    ComProc
	GetTop                     ComProc
	GetHeight                  ComProc
	GetWidth                   ComProc
	GetShouldDisplayMenuBar    ComProc
	GetShouldDisplayStatus     ComProc
	GetShouldDisplayToolbar    ComProc
	GetShouldDisplayScrollBars ComProc
}

type ICoreWebView2WindowFeatures struct {
	vtbl *_ICoreWebView2WindowFeaturesVtbl
}
