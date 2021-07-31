package webview2

type _ICoreWebView2PermissionRequestedEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri             ComProc
	GetPermissionKind  ComProc
	GetIsUserInitiated ComProc
	GetState           ComProc
	PutState           ComProc
	GetDeferral        ComProc
}

type ICoreWebView2PermissionRequestedEventArgs struct {
	vtbl *_ICoreWebView2PermissionRequestedEventArgsVtbl
}
