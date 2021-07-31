package webview2

type _ICoreWebView2ScriptDialogOpeningEventArgsVtbl struct {
	_IUnknownVtbl
	GetUri         ComProc
	GetKind        ComProc
	GetMessage     ComProc
	Accept         ComProc
	GetDefaultText ComProc
	GetResultText  ComProc
	PutResultText  ComProc
	GetDeferral    ComProc
}

type ICoreWebView2ScriptDialogOpeningEventArgs struct {
	vtbl *_ICoreWebView2ScriptDialogOpeningEventArgsVtbl
}
