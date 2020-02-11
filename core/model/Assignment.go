//Package model :
// Describing how an assignment
package model

//Assignment :
type Assignment struct {
	IDRole    uint64
	IDSubject uint64
	IDDomain  uint64
}

//NewAssignment :
func NewAssignment(r Role, s Subject, d Domain) Assignment {
	res := Assignment{
		IDRole:    r.ID,
		IDSubject: s.ID,
		IDDomain:  d.ID,
	}

	return res
}
