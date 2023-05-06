package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/mojafa")
	if err != nil {
		log.Fatalf("Error: %s", err)
		// log.Printf(error: %s, err)
		// os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s", resp.Status)
	}
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: can't copy - %s", err)
	// }

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Printf("%#v\n", r)
}

type Reply struct {
	Name         string
	Public_Repos int
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
