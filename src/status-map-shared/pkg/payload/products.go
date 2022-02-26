package payload

import "time"

type Products struct {
	Items []Product
}

type Product struct {
	Name       string `json:"name"`
	Pattern    string `json:"pattern"`
	Components []struct {
		Name         string   `json:"name"`
		Kind         string   `json:"kind"`
		Dependencies []string `json:"dependencies,omitempty"`
		Implements   []string `json:"implements,omitempty"`
	} `json:"components,omitempty"`
	Dependencies []struct {
		Name    string `json:"name"`
		Service string `json:"service"`
	} `json:"dependencies,omitempty"`
	Attributes []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"attributes"`
	Tombstone struct {
		Created time.Time `json:"created"`
		Expires time.Time `json:"expires"`
	} `json:"tombstone,omitempty"`
}
