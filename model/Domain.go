//Package model :
//Describing the domain structure
package model

//Domain :
// Domain only contains a name and is associated to a ID generated by sqlite
type Domain struct {
	ID   uint64
	Name string
}

//NewDomain :
// Constructor
func NewDomain(name string) *Domain {
	res := new(Domain)
	res.Name = name

	return res
}

//TODO : find a way to create safely the empty domain common of everyone
