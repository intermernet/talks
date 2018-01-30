// Copyright 2018 Mike Hughes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Package notes provides examples of custom godoc note types
package notes

// True returns true
//
// BUG(mike): Should return true
func True() bool {
	return false
}

// Divide divides one number by another
//
// TODO(mike): Check divisor to prevent divide by zero
func Divide(a, b int) int {
	return a / b
}

// Add adds one number to another
//
// MILESTONE(v1): Finally we can Add!
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts one number from another
//
// NOTE(mike): v1.1 Now with subtractive abilities
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies one number with another
//
// NOTE(mike): v1.2 Rock, shocking the mike of the many times, times the times tables
func Multiply(a, b int) int {
	return a * b
}
