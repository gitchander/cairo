package cairo

import (
	"errors"
	"fmt"
)

func newError(m string) error {
	return errors.New(fmt.Sprintf("cairo: %s", m))
}

func NewErrorFromStatus(status Status) error {

	if status != STATUS_SUCCESS {
		return newError(StatusToString(status))
	}

	return nil
}
