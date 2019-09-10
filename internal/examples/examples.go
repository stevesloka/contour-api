package examples

import (
	HTTPProxy "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

// Example represents a single example containing a set of HTTPProxy objects
type Example struct {
	Name         string
	Description  string
	DirName      string
	HTTPProxy []*HTTPProxy.HTTPProxy
}

// Get returns the set of example HTTPProxy objects
func (e *Example) Get() []Example {
	ex := []Example{
		e.basic(),
		e.delegate(),
		e.headerDelegation(),
		e.wildcardPath(),
		e.conditions_routeinclude(),
		e.header(),
	}
	return ex
}
