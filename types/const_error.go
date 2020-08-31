package types

// ConstError is used for a const error.
// ```
// const SomeError = goutil.ConstError("some error")
// ```
type ConstError string

func (e ConstError) Error() string {
	return string(e)
}
