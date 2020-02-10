//Package model :
//This file contains the necessary code to link a subject to his parent subject
package model

//SubjectLink :
// Describe a link in a graph to hierarchize subjects
type SubjectLink struct {
	IDSubjectParent uint64
	IDSubjectChild  uint64
}
