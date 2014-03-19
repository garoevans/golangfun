package main

import (
	"flag"
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	key := flag.String("key", "", "A key to get")
	val := flag.String("val", "", "Coupled with 'key' sets a new val")
	flag.Parse()

	if len(*key) == 0 {
		flag.PrintDefaults()
	} else {
		file, e := ioutil.ReadFile("./config.json")
		if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

		var f interface{}
		json.Unmarshal(file, &f)

		if len(*val) > 0 {
			f.(map[string]interface{})[*key] = *val
			newF, e := json.Marshal(&f)
			if e != nil {
					fmt.Printf("JSON error: %v\n", e)
					os.Exit(1)
			}
			ioutil.WriteFile("./config.json", newF, 0644)
		}

		fmt.Printf("%s: %s\n", *key, f.(map[string]interface{})[*key])
	}
}
