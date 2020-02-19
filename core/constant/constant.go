// Package constant :
// Contains constant variables
package constant

//nolint: golint, stylecheck
const (
	NAME_MAX_LEN int    = 255
	ROOT_DOMAIN  string = ""
)

//Action :
//Enum declaration
type Action int

//nolint: golint, stylecheck
const (
	ACTION_EXECUTE Action = 1 + iota
)

//Effect :
//Enum declaration
type Effect int

//nolint: golint, stylecheck
const (
	EFFECT_ALLOW Effect = 1 + iota
	EFFECT_DENY
)
