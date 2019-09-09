package examples

import (
	HTTPLoadBalancerv1 "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) delegate() Example {
	return Example{
		Name: "Delegation",
		Description: "A simple HTTPLoadBalancer which shows how an Include passes down a path prefix to children HTTPLoadBalancers. \n" +
			"# - GET projectcontour.io/ --> webapp.default:80 \n" +
			"# - GET projectcontour.io/blog --> servea.teama:8080",
		DirName: "delegation-prefix",
		HTTPLoadBalancer: []*HTTPLoadBalancerv1.HTTPLoadBalancer{{
			//ObjectMetaTemp: temp.ObjectMetaTemp{
			//	Name:      "delegate-http-loadbalancer",
			//	Namespace: "default",
			//},
			Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
				VirtualHost: &HTTPLoadBalancerv1.VirtualHost{
					Fqdn: "kuard.local",
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
						Name: "kuard",
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
				Routes: []HTTPLoadBalancerv1.Route{{
					Services: []HTTPLoadBalancerv1.Service{{
						Name: "nginx",
						Port: 80,
					}},
				}},
			},
		}},
	}
}
