package types

import "time"

// Duration is time.Duration with time.ParseDuration as UnmarshalText().
type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		// ignore null value
		return nil
	}

	v, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = Duration(v)
	return nil
}

func (d *Duration) MarshalText() ([]byte, error) {
	str := d.Duration().String()
	return []byte(str), nil
}

func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}
