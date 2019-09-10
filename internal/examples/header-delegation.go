package examples

import (
	HTTPProxy "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) headerDelegation() Example {
	return Example{
		Name: "HeaderDelegation",
		Description: "An HTTPProxy which shows how an Include passes down a specific header to children HTTPProxys. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-exmaples:80 \n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/ --> servea.teama:8080\n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/blog --> servea-blog.teama-blog:8080",
		DirName: "delegation-header",
		HTTPProxy: []*HTTPProxy.HTTPProxy{
			{
				//ObjectMetaTemp: temp.ObjectMetaTemp{
				//	Name:      "DelegateHeaderHTTPProxy",
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
							HeadersMatch: map[string][]string{
								"Content-Language": { "en" },
							},
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
					Includes: []HTTPProxy.Include{{
						Name:      "wwwsite-blog",
						Namespace: "teama-blog",
						Condition: HTTPProxy.Condition{
							Prefix: "/blog",
						},
					}},
					Routes: []HTTPProxy.Route{{
						Services: []HTTPProxy.Service{{
							Name: "servea",
							Port: 8080,
						}},
					}},
				},
			},
			{
				//ObjectMetaTemp: temp.ObjectMetaTemp{
				//	Name:      "wwwsite-blog",
				//	Namespace: "teama-blog",
				//},
				Spec: HTTPProxy.HTTPProxySpec{
					Routes: []HTTPProxy.Route{{
						Services: []HTTPProxy.Service{{
							Name: "servea-blog",
							Port: 8080,
						}},
					}},
				},
			}},
	}
}
