package examples

import (
	HTTPProxy "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) wildcardPath() Example {
	return Example{
		Name: "Delegation",
		Description: "A simple HTTPProxy which shows how an Include passes down a path prefix to children HTTPProxys. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-exmaples:80 \n" +
			"# - GET projectcontour.io/app/bar/foo --> servea.teama:8080",
		DirName: "delegation-prefix-wildcard",
		HTTPProxy: []*HTTPProxy.HTTPProxy{
			{
				//ObjectMetaTemp: temp.ObjectMetaTemp{
				//	Name:      "PathWildcardHTTPProxy",
				//	Namespace: "projectcontour-examples",
				//},
				Spec: HTTPProxy.HTTPProxySpec{
					VirtualHost: &HTTPProxy.VirtualHost{
						Fqdn: "projectcontour.io",
					},
					Includes: []HTTPProxy.Include{{
						Name:      "wwwsite",
						Namespace: "teama",
						Condition: HTTPProxy.Condition{
							Prefix: "/app/*/foo",
						},
					}},
					Routes: []HTTPProxy.Route{{
						Services: []HTTPProxy.Service{{
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
				Spec: HTTPProxy.HTTPProxySpec{
					Routes: []HTTPProxy.Route{{
						Services: []HTTPProxy.Service{{
							Name: "servea",
							Port: 8080,
						}},
					}},
				},
			}},
	}
}
