package main

import "fmt"
import "net/http"

var urls = []string{
	"http://registry.redhat.io",
	"http://quay.io",
	"http://cdn.quay.io",
	"http://cdn01.quay.io",
	"http://cdn02.quay.io",
	"http://cdn03.quay.io",
	"http://sso.redhat.com",
	"http://cert-api.access.redhat.com",
	"http://api.access.redhat.com",
	"http://infogw.api.openshift.com",
	"http://console.redhat.com/api/ingress",
	"http://cloud.redhat.com/api/ingress",
	"http://mirror.openshift.com",
	"http://storage.googleapis.com/openshift-release",
	"http://quayio-production-s3.s3.amazonaws.com",
	"http://api.openshift.com",
	"http://rhcos-redirector.apps.art.xq1c.p1.openshiftapps.com",
	"http://rhcos.mirror.openshift.com",
	"http://console.redhat.com/openshift",
	"http://registry.access.redhat.com",
}



func main() {
	fmt.Println("Checking required access for OpenShift 4")
	checkemoji := '\U00002705'
	failemoji := '\U0000274C'
	for i, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%d %c %s is not accessible. Check network\n", i, failemoji, url)
		} else {
			fmt.Printf("%d %c %s GOOD! (got %d response)\n", i, checkemoji, url, resp.StatusCode)
		}
	}
}
