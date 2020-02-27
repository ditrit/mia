// Package constant :
// Contains constant variables
package constant

//nolint: golint, stylecheck
const (
	NAME_MAX_LEN  int    = 255
	ROOT_DOMAINS  string = "__RootDomain__"
	ROOT_SUBJECTS string = "__RootSubjects__"
	ROOT_OBJECTS  string = "__RootObjects__"
)

//Action :
//Enum declaration
type Action int

//nolint: golint, stylecheck
const (
	ACTION_EXECUTE Action = 1 + iota
)
