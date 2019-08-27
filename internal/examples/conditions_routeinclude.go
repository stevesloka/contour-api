package examples

import (
	ingressroutev1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func (e *Example) conditions_routeinclude() Example {
	return Example{
		Name: "Delegation",
		Description: "An IngressRoute which shows how an Include passes down a path prefix to children IngressRoutes as well as routes define Conditions as well. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-examples:80 \n" +
			"# - GET -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' projectcontour.io/blog/engineering --> serve-mobile.teama:8080\n" +
			"# - GET -H projectcontour.io/blog/engineering --> serve-www.teama:8080",
		DirName: "delegation-routeinclude",
		IngressRoute: []*ingressroutev1.IngressRoute{{
			ObjectMetaTemp: temp.ObjectMetaTemp{
				Name:      "DelegateIngressRoute",
				Namespace: "projectcontour-examples",
			},
			Spec: ingressroutev1.IngressRouteSpec{
				VirtualHost: &ingressroutev1.VirtualHost{
					Fqdn: "projectcontour.io",
				},
				Includes: []ingressroutev1.Include{{
					Name:      "wwwsite",
					Namespace: "teama",
					Condition: []ingressroutev1.Condition{{
						Prefix: "/blog",
					}},
				}},
				Routes: []ingressroutev1.Route{{
					Services: []ingressroutev1.Service{{
						Name: "webapp",
						Port: 80,
					}},
				}},
			},
		}, {
			ObjectMetaTemp: temp.ObjectMetaTemp{
				Name:      "wwwsite",
				Namespace: "teama",
			},
			Spec: ingressroutev1.IngressRouteSpec{
				Routes: []ingressroutev1.Route{
					{
						Condition: ingressroutev1.Condition{
							Prefix: "/engineering",
							HeadersContain: map[string][]string{
								"user-agent": {"Android", "iPhone"},
							},
						},
						Services: []ingressroutev1.Service{{
							Name: "serve-mobile",
							Port: 8080,
						}},
					},
					{
						Condition: ingressroutev1.Condition{
							Prefix: "/engineering",
						},
						Services: []ingressroutev1.Service{{
							Name: "serve-www",
							Port: 8080,
						}},
					}},
			},
		}},
	}
}
