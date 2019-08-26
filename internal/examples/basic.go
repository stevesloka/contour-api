package examples

import (
	ingressroutev1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func Basic() *ingressroutev1.IngressRoute {
	return &ingressroutev1.IngressRoute{
		TempObjectMeta: temp.TempObjectMeta{
			Name:      "BasicIngressRoute",
			Namespace: "projectcontour-examples",
		},
		Spec: ingressroutev1.IngressRouteSpec{
			VirtualHost: &ingressroutev1.VirtualHost{
				Fqdn: "projectcontour.io",
			},
			Includes: []ingressroutev1.Include{{
				Name:      "wwwsite",
				Namespace: "teama",
				Conditions: []map[string]string{{
					"Prefix": "/",
				}},
			}},
			Routes: []ingressroutev1.Route{{
				Match: "/",
				Services: []ingressroutev1.Service{{
					Name: "webapp",
					Port: 80,
				}},
			}},
		},
		Status: ingressroutev1.Status{
			CurrentStatus: "valid",
			Description:   "valid IngressRoute",
		},
	}
}
