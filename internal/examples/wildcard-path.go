package examples

import (
	ingressroutev1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func (e *Example) wildcardPath() Example {
	return Example{
		Name: "Delegation",
		Description: "A simple IngressRoute which shows how an Include passes down a path prefix to children IngressRoutes. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-exmaples:80 \n" +
			"# - GET projectcontour.io/app/bar/foo --> servea.teama:8080",
		DirName: "delegation-prefix-wildcard",
		IngressRoute: []*ingressroutev1.IngressRoute{
			{
				ObjectMetaTemp: temp.ObjectMetaTemp{
					Name:      "PathWildcardIngressRoute",
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
							Prefix: "/app/*/foo",
						}},
					}},
					Routes: []ingressroutev1.Route{{
						Services: []ingressroutev1.Service{{
							Name: "webapp",
							Port: 80,
						}},
					}},
				},
			},
			{
				ObjectMetaTemp: temp.ObjectMetaTemp{
					Name:      "wwwsite",
					Namespace: "teama",
				},
				Spec: ingressroutev1.IngressRouteSpec{
					Routes: []ingressroutev1.Route{{
						Services: []ingressroutev1.Service{{
							Name: "servea",
							Port: 8080,
						}},
					}},
				},
			}},
	}
}
