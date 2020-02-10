//Package model :
//This file contains the necessary code to link a object to his parent object
package model

//ObjectLink :
// Describe a link in a graph to hierarchize objects
type ObjectLink struct {
	IDObjectParent uint64
	IDObjectChild  uint64
}
