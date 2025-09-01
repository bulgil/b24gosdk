package user

import "github.com/bulgil/b24gosdk"

type User struct {
	ID       b24gosdk.B24int `json:"ID"`
	Name     string          `json:"NAME"`
	LastName string          `json:"LAST_NAME"`
}
