package examples

import (
	HTTPLoadBalancerv1 "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

// Example represents a single example containing a set of HTTPLoadBalancer objects
type Example struct {
	Name         string
	Description  string
	DirName      string
	HTTPLoadBalancer []*HTTPLoadBalancerv1.HTTPLoadBalancer
}

// Get returns the set of example HTTPLoadBalancer objects
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
