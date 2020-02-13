//Package model :
//This file contains the necessary code to link an item to his parent item
package model

//ItemLink :
// Describe a link in a graph to hierarchize items
type ItemLink struct {
	Type     ItemType
	IDParent uint64 `gorm:"not null"`
	IDChild  uint64 `gorm:"not null"`
}
