package examples

import (
	HTTPProxy "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) conditions_routeinclude() Example {
	return Example{
		Name: "Delegation",
		Description: "An HTTPProxy which shows how an Include passes down a path prefix to children HTTPProxys as well as routes define Conditions as well. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-examples:80 \n" +
			"# - GET -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' projectcontour.io/blog/engineering --> serve-mobile.teama:8080\n" +
			"# - GET -H projectcontour.io/blog/engineering --> serve-www.teama:8080",
		DirName: "delegation-routeinclude",
		HTTPProxy: []*HTTPProxy.HTTPProxy{{
			//ObjectMetaTemp: temp.ObjectMetaTemp{
			//	Name:      "DelegateHTTPProxy",
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
						Prefix: "/blog",
					},
				}},
				Routes: []HTTPProxy.Route{{
					Services: []HTTPProxy.Service{{
						Name: "webapp",
						Port: 80,
					}},
				}},
			},
		}, {
			//ObjectMetaTemp: temp.ObjectMetaTemp{
			//	Name:      "wwwsite",
			//	Namespace: "teama",
			//},
			Spec: HTTPProxy.HTTPProxySpec{
				Routes: []HTTPProxy.Route{
					{
						Condition: &HTTPProxy.Condition{
							Prefix: "/engineering",
							HeadersContain: map[string][]string{
								"user-agent": {"Android", "iPhone"},
							},
						},
						Services: []HTTPProxy.Service{{
							Name: "serve-mobile",
							Port: 8080,
						}},
					},
					{
						Condition: &HTTPProxy.Condition{
							Prefix: "/engineering",
						},
						Services: []HTTPProxy.Service{{
							Name: "serve-www",
							Port: 8080,
						}},
					}},
			},
		}},
	}
}
