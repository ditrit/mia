//Package model :
//This file contains the necessary code to link a domain to his parent domain
package model

//DomainLink :
// Describe a link in a graph to hierarchize domains
type DomainLink struct {
	IDDomainParent uint64
	IDDomainChild  uint64
}
