//Package model :
// Describe how a permission is stored
package model

import (
	"errors"
	"iam/core/constant"
)

//Permission :
type Permission struct {
	IDRole   uint64
	IDDomain uint64
	IDObject uint64
	Action   constant.Action
}

//NewPermission :
// Constructor
func NewPermission(r Role, d Item, o Item, a constant.Action) (*Permission, error) {
	switch {
	case d.Type != ITEM_TYPE_DOMAIN:
		return nil, errors.New("domain item is not a domain")
	case o.Type != ITEM_TYPE_OBJ:
		return nil, errors.New("object item is not a object")
	}

	res := Permission{
		IDRole:   r.ID,
		IDDomain: d.ID,
		IDObject: o.ID,
		Action:   a,
	}

	return &res, nil
}
