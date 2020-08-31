package types

// Null is a type which instance occupys zero memory.
type Null struct{}

// NewNull creates a instance occupys zero memory.
func NewNull() Null {
	return Null{}
}
