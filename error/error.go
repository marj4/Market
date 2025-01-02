package error

import "fmt"

const ()

func Wrap(msg string, err error) error {
	return fmt.Errorf(msg, err)
}
