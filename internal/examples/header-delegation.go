package examples

import (
	ingressroutev1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func (e *Example) headerDelegation() Example {
	return Example{
		Name: "HeaderDelegation",
		Description: "An IngressRoute which shows how an Include passes down a specific header to children IngressRoutes. \n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-exmaples:80 \n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/ --> servea.teama:8080\n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/blog --> servea-blog.teama-blog:8080",
		DirName: "delegation-header",
		IngressRoute: []*ingressroutev1.IngressRoute{
			{
				ObjectMetaTemp: temp.ObjectMetaTemp{
					Name:      "DelegateHeaderIngressRoute",
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
							HeadersMatch: map[string][]string{
								"Content-Language": { "en" },
							},
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
					Includes: []ingressroutev1.Include{{
						Name:      "wwwsite-blog",
						Namespace: "teama-blog",
						Condition: []ingressroutev1.Condition{{
							Prefix: "/blog",
						}},
					}},
					Routes: []ingressroutev1.Route{{
						Services: []ingressroutev1.Service{{
							Name: "servea",
							Port: 8080,
						}},
					}},
				},
			},
			{
				ObjectMetaTemp: temp.ObjectMetaTemp{
					Name:      "wwwsite-blog",
					Namespace: "teama-blog",
				},
				Spec: ingressroutev1.IngressRouteSpec{
					Routes: []ingressroutev1.Route{{
						Services: []ingressroutev1.Service{{
							Name: "servea-blog",
							Port: 8080,
						}},
					}},
				},
			}},
	}
}
