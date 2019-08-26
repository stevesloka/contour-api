package examples

import (
	ingressroutev1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func (e *Example) basic() Example {
	return Example{
		Name: "Basic IngressRoute",
		Description: "A basic IngressRoute which routes a request to `projectcontour.io/` to the backend `webapp` in the namespace `projectcontour-examples`\n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-examples:80",
		DirName: "basic",
		IngressRoute: []*ingressroutev1.IngressRoute{{
			ObjectMetaTemp: temp.ObjectMetaTemp{
				Name:      "BasicIngressRoute",
				Namespace: "projectcontour-examples",
			},
			Spec: ingressroutev1.IngressRouteSpec{
				VirtualHost: &ingressroutev1.VirtualHost{
					Fqdn: "projectcontour.io",
				},
				Routes: []ingressroutev1.Route{{
					Match: "/",
					Services: []ingressroutev1.Service{{
						Name: "webapp",
						Port: 80,
					}},
				}},
			},
		}},
	}
}
