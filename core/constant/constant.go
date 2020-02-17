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
	ACTION_EXECUTE Action = iota
)

//Effect :
//Enum declaration
type Effect int

//nolint: golint, stylecheck
const (
	EFFECT_ALLOW Effect = iota
	EFFECT_DENY
)
