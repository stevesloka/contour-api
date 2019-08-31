package examples

import (
	HTTPLoadBalancerv1 "github.com/projectcontour/contour-api/apis/projectcontour/v1"
	"github.com/projectcontour/contour-api/internal/temp"
)

func (e *Example) basic() Example {
	return Example{
		Name: "Basic HTTPLoadBalancer",
		Description: "A basic HTTPLoadBalancer which routes a request to `projectcontour.io/` to the backend `webapp` in the namespace `projectcontour-examples`\n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-examples:80",
		DirName: "basic",
		HTTPLoadBalancer: []*HTTPLoadBalancerv1.HTTPLoadBalancer{{
			ObjectMetaTemp: temp.ObjectMetaTemp{
				Name:      "basic-httploadbalancer",
				Namespace: "projectcontour-examples",
			},
			Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
				VirtualHost: &HTTPLoadBalancerv1.VirtualHost{
					Fqdn: "projectcontour.io",
				},
				Routes: []HTTPLoadBalancerv1.Route{{
					Services: []HTTPLoadBalancerv1.Service{{
						Name: "webapp",
						Port: 80,
					}},
				}},
			},
		}},
	}
}
