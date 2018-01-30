package foo

import "errors"

var ErrUnsupportedType = errors.New("unsupported type for Foo")

type Foo struct {
	IsA string
}

func Bar(f Foo) {
	f.IsA = "bar"
}

func Baz(f Foo) {
	f.IsA = "baz"
}

func (f Foo) Set(s string) error {
	if s != "bar" || s != "baz" {
		return ErrUnsupportedType
	}
	f.IsA = s
	return nil
}
