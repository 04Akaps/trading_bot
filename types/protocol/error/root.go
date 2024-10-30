package error

import (
	"errors"
	"strings"
)

type ErrString string

const (
	NOT_SUPPORTED_EXCNAHGER = ErrString("Not Support Exchanger")
)

func (e ErrString) E(msg ...string) error {
	fullMessage := strings.Join(append([]string{string(e)}, msg...), " ")
	return errors.New(fullMessage)
}
