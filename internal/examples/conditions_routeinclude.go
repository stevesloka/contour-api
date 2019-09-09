package examples

import (
	HTTPLoadBalancerv1 "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) conditions_routeinclude() Example {
	return Example{
		Name: "Delegation",
		Description: "An HTTPLoadBalancer which shows how an Include passes down a path prefix to children HTTPLoadBalancers as well as routes define Conditions as well. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-examples:80 \n" +
			"# - GET -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' projectcontour.io/blog/engineering --> serve-mobile.teama:8080\n" +
			"# - GET -H projectcontour.io/blog/engineering --> serve-www.teama:8080",
		DirName: "delegation-routeinclude",
		HTTPLoadBalancer: []*HTTPLoadBalancerv1.HTTPLoadBalancer{{
			//ObjectMetaTemp: temp.ObjectMetaTemp{
			//	Name:      "DelegateHTTPLoadBalancer",
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
						Prefix: "/blog",
					},
				}},
				Routes: []HTTPLoadBalancerv1.Route{{
					Services: []HTTPLoadBalancerv1.Service{{
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
			Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
				Routes: []HTTPLoadBalancerv1.Route{
					{
						Condition: &HTTPLoadBalancerv1.Condition{
							Prefix: "/engineering",
							HeadersContain: map[string][]string{
								"user-agent": {"Android", "iPhone"},
							},
						},
						Services: []HTTPLoadBalancerv1.Service{{
							Name: "serve-mobile",
							Port: 8080,
						}},
					},
					{
						Condition: &HTTPLoadBalancerv1.Condition{
							Prefix: "/engineering",
						},
						Services: []HTTPLoadBalancerv1.Service{{
							Name: "serve-www",
							Port: 8080,
						}},
					}},
			},
		}},
	}
}
