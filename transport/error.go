package transport

import "fmt"

type B24Error struct {
	Err            string
	ErrDescription string
}

func (e B24Error) Error() string {
	err := "<nil>"
	errDescription := "<nil>"

	if e.Err != "" {
		err = e.Err
	}

	if e.ErrDescription != "" {
		errDescription = e.ErrDescription
	}

	return fmt.Sprintf("%s - %s", err, errDescription)
}
