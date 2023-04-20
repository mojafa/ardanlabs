package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// type Reply struct {
// 	Name     string
// 	NumRepos int `json:"public_repos"`
// 	//Login        string
// }

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	res, err := http.Get(url)

	if err != nil {
		return "", 0, err
	}
	if res.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v-%s", url, res.Status)
	}
	fmt.Printf("Content-Type: %s\n", res.Header.Get("Content-Type"))

	//declare anonymous struct

	var r struct {
		Name     string
		NumRepos int `json:"public_repos"`
		//Login        string
	}
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.NumRepos, nil
}

func main() {
	fmt.Println(githubInfo("mojafa"))
	// res, err := http.Get("https://api.github.com/users/tebeka")

	// if err != nil {
	// 	log.Fatalf("error : %s", err)
	// 	// log.Printf("error: %s", err)
	// 	// os.Exit(1)
	// }
	// if res.StatusCode != http.StatusOK {
	// 	log.Fatalf("error : %s", res.Status)
	// }
	// fmt.Printf("Content-Type: %s\n", res.Header.Get("Content-Type"))
	// if _, err := io.Copy(os.Stdout, res.Body); err != nil {
	// 	log.Fatalf("error: can't copy  %s", err)

	// }
	// var r Reply
	// dec := json.NewDecoder(res.Body)
	// if err := dec.Decode(&r); err != nil {
	// 	log.Fatalf("error: can't decode  %s", err)

	// }
	// fmt.Printf("%#v\n", r)
}

//Marshalling or serialization and unmarshaling or deserialization in Go

/*
JSON <-> Go

true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8,...
array <-> []any ([]interface{}) //use any or as it was known before the empty interface
object <-> map[string]any, struct



encoding/json API
JSON -> io.Reader -> Go :json.Decoder
JSON -> []byte -> Go : json.Unmarshal
go -> io.Writer -> JSON :json.Encoder
go -> []byte ->JSON :json.Marshal
*/
