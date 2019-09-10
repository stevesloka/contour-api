package examples

import (
	HTTPProxy "github.com/heptio/contour/apis/projectcontour/v1alpha1"
)

func (e *Example) delegate() Example {
	return Example{
		Name: "Delegation",
		Description: "A simple HTTPProxy which shows how an Include passes down a path prefix to children HTTPProxys. \n" +
			"# - GET projectcontour.io/ --> webapp.default:80 \n" +
			"# - GET projectcontour.io/blog --> servea.teama:8080",
		DirName: "delegation-prefix",
		HTTPProxy: []*HTTPProxy.HTTPProxy{{
			//ObjectMetaTemp: temp.ObjectMetaTemp{
			//	Name:      "delegate-http-loadbalancer",
			//	Namespace: "default",
			//},
			Spec: HTTPProxy.HTTPProxySpec{
				VirtualHost: &HTTPProxy.VirtualHost{
					Fqdn: "kuard.local",
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
			Spec: HTTPProxy.HTTPProxySpec{
				Routes: []HTTPProxy.Route{{
					Services: []HTTPProxy.Service{{
						Name: "nginx",
						Port: 80,
					}},
				}},
			},
		}},
	}
}
