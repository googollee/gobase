package types

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		d    time.Duration
		want string
	}{
		{time.Second, "1s"},
		{5 * time.Minute, "5m0s"},
	}

	for _, test := range tests {
		t.Run(test.want, func(t *testing.T) {
			duration := Duration(test.d)

			got, err := duration.MarshalText()
			if err != nil {
				t.Fatal("duration.MarshalText() != nil, error:", err)
			}
			if want, got := test.want, string(got); want != got {
				t.Errorf("duration.MarshalText() returns:\n\twant: %q\n\tgot:  %q", want, got)
			}

			if err := duration.UnmarshalText(got); err != nil {
				t.Fatal("duration.UnmarshalText() != nil, error:", err)
			}
			if want, got := test.d, duration.Duration(); want != got {
				t.Errorf("duration.UnmarshalText() returns:\n\twant: %q\n\tgot:  %q", want, got)
			}
		})
	}
}

func TestEmptyDuration(t *testing.T) {
	tests := []struct {
		text []byte
		want Duration
	}{
		{nil, 0},
		{[]byte{}, 0},
		{[]byte(""), 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%#v", test.text), func(t *testing.T) {
			var got Duration
			if err := got.UnmarshalText(test.text); err != nil {
				t.Fatal("duration.UnmarshalText() != nil, error:", err)
			}
			if want, got := test.want, got; want != got {
				t.Errorf("duration.UnmarshalText() returns:\n\twant: %q\n\tgot:  %q", want, got)
			}
		})
	}

}
