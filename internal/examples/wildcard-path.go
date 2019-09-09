package examples

import (
	HTTPLoadBalancerv1 "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) wildcardPath() Example {
	return Example{
		Name: "Delegation",
		Description: "A simple HTTPLoadBalancer which shows how an Include passes down a path prefix to children HTTPLoadBalancers. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-exmaples:80 \n" +
			"# - GET projectcontour.io/app/bar/foo --> servea.teama:8080",
		DirName: "delegation-prefix-wildcard",
		HTTPLoadBalancer: []*HTTPLoadBalancerv1.HTTPLoadBalancer{
			{
				//ObjectMetaTemp: temp.ObjectMetaTemp{
				//	Name:      "PathWildcardHTTPLoadBalancer",
				//	Namespace: "projectcontour-examples",
				//},
				Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
					VirtualHost: &HTTPLoadBalancerv1.VirtualHost{
						Fqdn: "projectcontour.io",
					},
					Includes: []HTTPLoadBalancerv1.Include{{
						Name:      "wwwsite",
						Namespace: "teama",
						Condition: HTTPLoadBalancerv1.Condition{
							Prefix: "/app/*/foo",
						},
					}},
					Routes: []HTTPLoadBalancerv1.Route{{
						Services: []HTTPLoadBalancerv1.Service{{
							Name: "webapp",
							Port: 80,
						}},
					}},
				},
			},
			{
				//ObjectMetaTemp: temp.ObjectMetaTemp{
				//	Name:      "wwwsite",
				//	Namespace: "teama",
				//},
				Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
					Routes: []HTTPLoadBalancerv1.Route{{
						Services: []HTTPLoadBalancerv1.Service{{
							Name: "servea",
							Port: 8080,
						}},
					}},
				},
			}},
	}
}
