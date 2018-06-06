package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func test1() {
	var out = `
    a: First!
    f: Second
    b:
      c: 
        f: Third
    `
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(out), &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m["b"].(map[interface{}]interface{})["c"].(map[interface{}]interface{})["f"])
}

type Webapps struct {
	// Domain_list []map[interface{}]interface{}
	Domain_list interface{}
}
type WebAppFmt struct {
	Webapps map[int]*Webapps
}

func test2() {
	out, err := ioutil.ReadFile("test2.yaml")

	m := WebAppFmt{}
	err = yaml.Unmarshal([]byte(out), &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", m.Webapps[1].Domain_list)
	fmt.Printf("%+v\n", m.Webapps[2].Domain_list)

	m.Webapps[1].Domain_list = []interface{}{}

	yamlBytes, err := yaml.Marshal(m)

	fmt.Printf("%+v", string(yamlBytes))
}

func main() {
	test2()
}
