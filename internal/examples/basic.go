package examples

import (
	HTTPLoadBalancerv1 "github.com/heptio/contour/apis/projectcontour/v1alpha1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (e *Example) basic() Example {
	return Example{
		Name: "Basic HTTPLoadBalancer",
		Description: "A basic HTTPLoadBalancer which routes a request to `projectcontour.io/` to the backend `webapp` in the namespace `projectcontour-examples`\n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-examples:80",
		DirName: "basic",
		HTTPLoadBalancer: []*HTTPLoadBalancerv1.HTTPLoadBalancer{{
			//ObjectMeta: metav1.ObjectMeta{
			//	Name:      "basic-httploadbalancer",
			//	Namespace: "default",
			//},
			Spec: HTTPLoadBalancerv1.HTTPLoadBalancerSpec{
				VirtualHost: &HTTPLoadBalancerv1.VirtualHost{
					Fqdn: "kuard.local",
				},
				Routes: []HTTPLoadBalancerv1.Route{{
					Services: []HTTPLoadBalancerv1.Service{{
						Name: "kuard",
						Port: 80,
					}},
				}},
			},
		}},
	}
}
