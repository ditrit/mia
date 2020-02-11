// Package core :
// Contains constant variables
package core

//nolint: golint, stylecheck
const (
	NAME_MAX_LEN int = 255
)

//Action :
//Enum declaration
type Action int

//nolint: golint, stylecheck
const (
	ACTION_ALLOW Action = iota
	ACTION_DENY
)
