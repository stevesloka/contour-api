package examples

import (
	HTTPProxy "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) header() Example {
	return Example{
		Name: "Header",
		Description: "An HTTPProxy which shows how a Condition on a route routes requests with specific headers. \n" +
			"# - GET projectcontour.io/ --> webapp-others.projectcontour-examples:80 \n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/ --> webapp-en.projectcontour-examples:8080\n",
		DirName: "header",
		HTTPProxy: []*HTTPProxy.HTTPProxy{
			{
				//ObjectMetaTemp: temp.ObjectMetaTemp{
				//	Name:      "HeaderHTTPProxy",
				//	Namespace: "projectcontour-examples",
				//},
				Spec: HTTPProxy.HTTPProxySpec{
					VirtualHost: &HTTPProxy.VirtualHost{
						Fqdn: "projectcontour.io",
					},
					Routes: []HTTPProxy.Route{
						{
						Condition: &HTTPProxy.Condition{
							HeadersMatch: map[string][]string{
								"Content-Language": { "en" },
							},
						},
						Services: []HTTPProxy.Service{{
							Name: "webapp-en",
							Port: 80,
						}},
					},
					{
						Condition: &HTTPProxy.Condition{
							Prefix: "/",
						},
						Services: []HTTPProxy.Service{{
							Name: "webapp-others",
							Port: 80,
						}},
					}},
				},
			}},
	}
}
