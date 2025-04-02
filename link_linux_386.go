//go:build !windows && !darwin
// +build !windows,!darwin

// 11 december 2015

package ui

// #cgo LDFLAGS: -lm -ldl
// #cgo pkg-config: gtk+-3.0
import "C"
