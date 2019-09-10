package examples

import (
	HTTPProxy "github.com/heptio/contour/apis/projectcontour/v1alpha1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (e *Example) basic() Example {
	return Example{
		Name: "Basic HTTPProxy",
		Description: "A basic HTTPProxy which routes a request to `projectcontour.io/` to the backend `webapp` in the namespace `projectcontour-examples`\n" +
			"# - GET projectcontour.io/ --> webapp.projectcontour-examples:80",
		DirName: "basic",
		HTTPProxy: []*HTTPProxy.HTTPProxy{{
			//ObjectMeta: metav1.ObjectMeta{
			//	Name:      "basic-HTTPProxy",
			//	Namespace: "default",
			//},
			Spec: HTTPProxy.HTTPProxySpec{
				VirtualHost: &HTTPProxy.VirtualHost{
					Fqdn: "kuard.local",
				},
				Routes: []HTTPProxy.Route{{
					Services: []HTTPProxy.Service{{
						Name: "kuard",
						Port: 80,
					}},
				}},
			},
		}},
	}
}
