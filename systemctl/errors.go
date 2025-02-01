package systemctl

import "fmt"

func NewError(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
