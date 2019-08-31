package examples

import (
	HTTPLoadBalancerv1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func (e *Example) header() Example {
	return Example{
		Name: "Header",
		Description: "An HTTPLoadBalancer which shows how a Condition on a route routes requests with specific headers. \n" +
			"# - GET projectcontour.io/ --> webapp-others.projectcontour-examples:80 \n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/ --> webapp-en.projectcontour-examples:8080\n",
		DirName: "header",
		HTTPLoadBalancer: []*HTTPLoadBalancerv1.HTTPLoadBalancer{
			{
				ObjectMetaTemp: temp.ObjectMetaTemp{
					Name:      "HeaderHTTPLoadBalancer",
					Namespace: "projectcontour-examples",
				},
				Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
					VirtualHost: &HTTPLoadBalancerv1.VirtualHost{
						Fqdn: "projectcontour.io",
					},
					Routes: []HTTPLoadBalancerv1.Route{
						{
						Condition: HTTPLoadBalancerv1.Condition{
							HeadersMatch: map[string][]string{
								"Content-Language": { "en" },
							},
						},
						Services: []HTTPLoadBalancerv1.Service{{
							Name: "webapp-en",
							Port: 80,
						}},
					},
					{
						Condition: HTTPLoadBalancerv1.Condition{
							Prefix: "/",
						},
						Services: []HTTPLoadBalancerv1.Service{{
							Name: "webapp-others",
							Port: 80,
						}},
					}},
				},
			}},
	}
}
