package examples

import (
	ingressroutev1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func (e *Example) header() Example {
	return Example{
		Name: "Header",
		Description: "An IngressRoute which shows how a Condition on a route routes requests with specific headers. \n" +
			"# - GET projectcontour.io/ --> webapp-others.projectcontour-examples:80 \n" +
			"# - GET -H 'Content-Language: en' projectcontour.io/ --> webapp-en.projectcontour-examples:8080\n",
		DirName: "header",
		IngressRoute: []*ingressroutev1.IngressRoute{
			{
				ObjectMetaTemp: temp.ObjectMetaTemp{
					Name:      "HeaderIngressRoute",
					Namespace: "projectcontour-examples",
				},
				Spec: ingressroutev1.IngressRouteSpec{
					VirtualHost: &ingressroutev1.VirtualHost{
						Fqdn: "projectcontour.io",
					},
					Routes: []ingressroutev1.Route{
						{
						Condition: ingressroutev1.Condition{
							HeadersMatch: map[string][]string{
								"Content-Language": { "en" },
							},
						},
						Services: []ingressroutev1.Service{{
							Name: "webapp-en",
							Port: 80,
						}},
					},
					{
						Condition: ingressroutev1.Condition{
							Prefix: "/",
						},
						Services: []ingressroutev1.Service{{
							Name: "webapp-others",
							Port: 80,
						}},
					}},
				},
			}},
	}
}
