//Package model :
//Describing the role structure
package model

import (
	"errors"
	"iam/core/constant"
)

//Role :
// Role only contains a name and is associated to a ID generated by sqlite
type Role struct {
	ID   uint64
	Name string
}

//NewRole :
// Constructor
func NewRole(name string) (*Role, error) {
	if len(name) > constant.NAME_MAX_LEN {
		return nil, errors.New("the name cannot be longer that 255 characters")
	} else if len(name) == 0 {
		return nil, errors.New("the name cannot be empty")
	}

	res := new(Role)
	res.Name = name

	return res, nil
}
