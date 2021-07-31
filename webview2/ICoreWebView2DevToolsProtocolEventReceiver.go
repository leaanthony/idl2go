package webview2

type _ICoreWebView2DevToolsProtocolEventReceiverVtbl struct {
	_IUnknownVtbl
	AddDevToolsProtocolEventReceived    ComProc
	RemoveDevToolsProtocolEventReceived ComProc
}

type ICoreWebView2DevToolsProtocolEventReceiver struct {
	vtbl *_ICoreWebView2DevToolsProtocolEventReceiverVtbl
}
