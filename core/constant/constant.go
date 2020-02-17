// Package constant :
// Contains constant variables
package constant

//nolint: golint, stylecheck
const (
	NAME_MAX_LEN int = 255
)

//Action :
//Enum declaration
type Action int

//nolint: golint, stylecheck
const (
	Execute Action = iota
)

//Effect :
//Enum declaration
type Effect int

//nolint: golint, stylecheck
const (
	ACTION_ALLOW Effect = iota
	ACTION_DENY
)
