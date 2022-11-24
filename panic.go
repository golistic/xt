// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"fmt"
	"io"
	"testing"
)

// Panics checks if function f panics.
func Panics(t *testing.T, f func()) {
	TestHelper(t)

	panics(t, nil, "", f)
}

// PanicsEq checks if function f panics and whether the panic message is equal to exp.
func PanicsEq(t *testing.T, exp string, f func()) {
	TestHelper(t)

	panics(t, nil, exp, f)
}

func panics(t *testing.T, out io.Writer, exp string, f func()) {
	TestHelper(t)

	defer func() {
		TestHelper(t)

		r := recover()
		if r == nil {
			fatal(t, out, fmt.Sprintf("\u001B[31;1mexpected panic\u001B[0m"))
		}
		if exp != "" {
			Eq(t, exp, fmt.Sprintf("%v", r), "expected panic message ^^^^")
		}
	}()

	f()
}
