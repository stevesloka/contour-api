package examples

import (
	ingressroutev1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
)

// Example represents a single example containing a set of IngressRoute objects
type Example struct {
	Name         string
	Description  string
	DirName      string
	IngressRoute []*ingressroutev1.IngressRoute
}

// Get returns the set of example IngressRoute objects
func (e *Example) Get() []Example {
	ex := []Example{
		e.basic(),
		e.delegate(),
		e.headerDelegation(),
		e.wildcardPath(),
		e.conditions_routeinclude(),
	}
	return ex
}
