package examples

import (
	HTTPLoadBalancerv1 "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) headerDelegation() Example {
	return Example{
		Name: "HeaderDelegation",
		Description: "An HTTPLoadBalancer which shows how an Include passes down a specific header to children HTTPLoadBalancers. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-exmaples:80 \n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/ --> servea.teama:8080\n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/blog --> servea-blog.teama-blog:8080",
		DirName: "delegation-header",
		HTTPLoadBalancer: []*HTTPLoadBalancerv1.HTTPLoadBalancer{
			{
				//ObjectMetaTemp: temp.ObjectMetaTemp{
				//	Name:      "DelegateHeaderHTTPLoadBalancer",
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
							HeadersMatch: map[string][]string{
								"Content-Language": { "en" },
							},
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
					Includes: []HTTPLoadBalancerv1.Include{{
						Name:      "wwwsite-blog",
						Namespace: "teama-blog",
						Condition: HTTPLoadBalancerv1.Condition{
							Prefix: "/blog",
						},
					}},
					Routes: []HTTPLoadBalancerv1.Route{{
						Services: []HTTPLoadBalancerv1.Service{{
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
				Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
					Routes: []HTTPLoadBalancerv1.Route{{
						Services: []HTTPLoadBalancerv1.Service{{
							Name: "servea-blog",
							Port: 8080,
						}},
					}},
				},
			}},
	}
}
