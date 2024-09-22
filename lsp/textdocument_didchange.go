package lsp

type TextDocumentDidChangeNotification struct {
	Notification
	Params DidChangeTextDocumentsParams `json:"params"`
}

type DidChangeTextDocumentsParams struct {
	TextDocument   TextDocumentIdentifier           `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
