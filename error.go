package cairo

import (
	"errors"
	"fmt"
)

func newCairoError(message string) error {
	return errors.New(fmt.Sprintf("cairo: %s", message))
}
