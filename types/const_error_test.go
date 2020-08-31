package types

import (
	"errors"
	"fmt"
	"testing"
)

func TestConstError(t *testing.T) {
	const someError = ConstError("some error")
	f := func() error {
		return fmt.Errorf("in func: %w", someError)
	}
	err := f()

	if want, got := "in func: some error", err.Error(); want != got {
		t.Errorf("err.Error()\n\twant: %q\ngot:  %q", want, got)
	}
	if !errors.Is(err, someError) {
		t.Errorf("err should be a someError, but it doesn't")
	}
}
