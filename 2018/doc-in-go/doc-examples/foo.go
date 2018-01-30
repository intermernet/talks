// Copyright 2018 Mike Hughes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Package foo provides functionality to bar and baz foos.

package foo

import "errors"

// ErrUnsupportedType is returned if attempting to set a Foo to anything other than a bar or a baz.
var ErrUnsupportedType = errors.New("unsupported type for Foo")

// Foo can be either a bar or a baz.
type Foo struct {
	IsA string
}

// Bar sets a Foo to type bar. // HL
func Bar(f Foo) {
	f.IsA = "bar"
}

// Baz sets a Foo to type baz.
func Baz(f Foo) {
	f.IsA = "baz"
}

// Set changes a Foo to either a bar or a baz.
// It returns an error if an unupported type is specified.
func (f Foo) Set(s string) error {
	if s != "bar" || s != "baz" {
		return ErrUnsupportedType
	}
	f.IsA = s
	return nil
}
