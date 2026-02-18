package main

// An Op is a single regular expression operator.
//
//go:generate stringer -type Op -trimprefix Op
type Op uint8

const (
	OpLiteral Op = iota
	OpConcat
	OpAlternate
)
