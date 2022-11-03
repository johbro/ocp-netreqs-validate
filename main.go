package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"gopkg.in/yaml.v3"
)

type Site struct {
	Name     string   `yaml:"name"`
	Validate string   `yaml:"validate"`
	Urls     []string `yaml:"urls"`
}

type YMLSites struct {
	Sites []Site `yaml:"sites"`
}


func main() {
	fmt.Println("Checking required access for OpenShift 4")
	checkemoji := '\U00002705'
	failemoji := '\U0000274C'
	yfile, err := ioutil.ReadFile("sites.yaml")

	if err != nil {

		log.Fatal(err)
	}

	var siteList YMLSites
	err2 := yaml.Unmarshal(yfile, &siteList)

	if err2 != nil {

		log.Fatal(err2)
	}

	//fmt.Printf("%s\n", siteList)
	for k, site := range siteList.Sites {
		//fmt.Printf("key: %s, value: %s \n", k, site.Name)
		Unused(k)
		if site.Validate == "all" {
			fmt.Printf("validating ALL sites for: %s \n", site.Name)
			for i, url := range site.Urls {
				Unused(i)
				resp, err3 := http.Get(url)
				if err3 != nil {
					fmt.Printf("  %c %s is not accessible. Check network\n", failemoji, url)
				} else {
					fmt.Printf("  %c %s GOOD! (got %d response)\n", checkemoji, url, resp.StatusCode)
				}
			}
		} else {
			//as long as one site is online
			fmt.Printf("validating ANY sites for: %s \n", site.Name)

			var successCount int
			var failureCount int

			for i := 1; i < len(site.Urls); i++ {
				resp, err3 := http.Get(site.Urls[i])
				if err3 != nil {
					failureCount += 1
				} else {
					successCount += 1
				}
			}
			if successCount >= 1 {
				fmt.Printf("  %c %s [%d/%d] GOOD!\n", checkemoji, site.Name, successCount, len(site.Urls))
			} else {
				fmt.Printf("  %c %s is not accessible. Check network\n", checkemoji, site.Name)
			}

		}
	}

}
func Unused(x ...interface{}) {}

