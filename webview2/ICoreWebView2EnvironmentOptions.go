package webview2

type _ICoreWebView2EnvironmentOptionsVtbl struct {
	_IUnknownVtbl
	GetAdditionalBrowserArguments             ComProc
	PutAdditionalBrowserArguments             ComProc
	GetLanguage                               ComProc
	PutLanguage                               ComProc
	GetTargetCompatibleBrowserVersion         ComProc
	PutTargetCompatibleBrowserVersion         ComProc
	GetAllowSingleSignOnUsingOSPrimaryAccount ComProc
	PutAllowSingleSignOnUsingOSPrimaryAccount ComProc
}

type ICoreWebView2EnvironmentOptions struct {
	vtbl *_ICoreWebView2EnvironmentOptionsVtbl
}
