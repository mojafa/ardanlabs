package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println(githubInfo("mojafa"))
}

// in orer to help JSON dcoder we can define our own types
// type Reply struct {
// 	Name         string
// 	Public_Repos int
// }

// githubInfo returns name and no of public repos for login
func githubInfo(login string) (string, int, error) {
	// TODO
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
		// log.Fatalf("Error: %s", err)
		// log.Printf(error: %s, err)
		// os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		// log.Fatalf("Error: %s", resp.Status)
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: can't copy - %s", err)
	// }

	// var r Reply  - replaced with anonymous struct

	var r struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}

	dec := json.NewDecoder(resp.Body)
	//go passes by values which gets copy of r. &r is a pointer to r
	if err := dec.Decode(&r); err != nil {
		// log.Fatalf("error: can't decode - %s", err)
		return "", 0, err
	}
	// fmt.Printf("%#v\n", r)
	return r.Name, r.NumRepos, nil
}

// curl -i - H 'User-Agent: go' https://api.github.com/users/mojafa

//serializations and deserializations
/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, int64, uint64
array <-> []any([]interface{})
object <-> map[string]any, struct (map[string]interface{})
eg time type in go, like timstamp, in json yuop need to store as int or string
JSON encoder and decoder let you do custom types
*/

//encoding/json API
/*
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
JSOn <- io.Writer <- Go: json.Encoder
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/

//go has pointers which are used for references
// pointers in go are used for references

//in go we have a slice

// eg []byte, []int, []int32, ...

//in JSON you can mix types eg ["hi", 7, null]
//in go you cabn't mix types
// but you can use []any
//any (interfce{})
//[]any
//any is 1.18+

// go has auto tagging in sturcts
// eg
// type Reply struct {
// NumRepos int can be  used to rep Public_repos dound in the Github resp, but you could do
// NumRepos int `json:"public-repos`"
// }
