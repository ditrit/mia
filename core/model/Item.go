// Package model :
//Item is an struct that can be a Subject, a Domain, or an object
package model

import (
	"errors"
	"iam/core/constant"
)

//ItemType :
// identifier type to differ a subject, a domain and an object
type ItemType uint16

// identifiers
// nolint: golint, stylecheck
const (
	ITEM_TYPE_SUBJ ItemType = iota
	ITEM_TYPE_DOMAIN
	ITEM_TYPE_OBJ
)

//Item :
// THE item
type Item struct {
	Type ItemType
	ID   uint64 `gorm:"unique;not null"`
	Name string
}

//NewSubject :
// Constructor
func NewSubject(name string) (*Item, error) {
	switch {
	case len(name) > constant.NAME_MAX_LEN:
		return nil, errors.New("the name cannot be longer that 255 characters")
	case len(name) == 0:
		return nil, errors.New("the name cannot be empty")
	}

	res := new(Item)
	res.Type = ITEM_TYPE_SUBJ
	res.Name = name

	return res, nil
}

//NewDomain :
// Constructor
func NewDomain(name string) (*Item, error) {
	switch {
	case len(name) > constant.NAME_MAX_LEN:
		return nil, errors.New("the name cannot be longer that 255 characters")
	case len(name) == 0:
		return nil, errors.New("the name cannot be empty")
	}

	res := new(Item)
	res.Type = ITEM_TYPE_DOMAIN
	res.Name = name

	return res, nil
}

//GetRootDomain :
// Should not be called in normal workflow
// Used only in database initialization
func GetRootDomain() *Item {
	res := new(Item)
	res.Type = ITEM_TYPE_DOMAIN
	res.Name = ""

	return res
}

//NewObject :
// Constructor
func NewObject(name string) (*Item, error) {
	switch {
	case len(name) > constant.NAME_MAX_LEN:
		return nil, errors.New("the name cannot be longer that 255 characters")
	case len(name) == 0:
		return nil, errors.New("the name cannot be empty")
	}

	res := new(Item)
	res.Type = ITEM_TYPE_OBJ
	res.Name = name

	return res, nil
}
