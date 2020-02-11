//Package model :
// Describe how a permission is stored
package model

import (
	"iam/core"
)

//Permission :
type Permission struct {
	IDRole   uint64
	IDDomain uint64
	IDObject uint64
	Action   core.Action
}

//NewPermission :
// Constructor
func NewPermission(r Role, d Domain, o Object, a core.Action) Permission {
	res := Permission{
		IDRole:   r.ID,
		IDDomain: d.ID,
		IDObject: o.ID,
		Action:   a,
	}

	return res
}
