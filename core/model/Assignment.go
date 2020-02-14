//Package model :
// Describing how an assignment
package model

import (
	"errors"
)

//Assignment :
type Assignment struct {
	IDRole    uint64
	IDSubject uint64
	IDDomain  uint64
}

//NewAssignment :
func NewAssignment(r Role, s Item, d Item) (*Assignment, error) {
	switch {
	case s.Type != ITEM_TYPE_SUBJ:
		return nil, errors.New("subject item is not a subject")
	case d.Type != ITEM_TYPE_DOMAIN:
		return nil, errors.New("domain item is not a domain")
	}

	res := Assignment{
		IDRole:    r.ID,
		IDSubject: s.ID,
		IDDomain:  d.ID,
	}

	return &res, nil
}
