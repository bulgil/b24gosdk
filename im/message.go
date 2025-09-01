package im

import "github.com/bulgil/b24gosdk"

type Message struct {
	DialogID b24gosdk.B24int `json:"DIALOG_ID"`
	Message  string          `json:"MESSAGE"`
}
