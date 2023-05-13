package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := Sha1Sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)

	sig, err = Sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

// if file names end with .gz
// $ cat http.log.gs |gunzip | sha1sum
// else
// $ cat http.log.gs | sha1sum
func Sha1Sum(filename string) (string, error) {
	//idiom: acquire a resource, check for error, defer release

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close() //defers are called like in LIFO order

	var r io.Reader = file
	if strings.HasSuffix(filename, ".gz") {
		gzFile, err := gzip.NewReader(file)
		// 		file, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gzFile.Close()
		r = gzFile
	}

	w := sha1.New()
	// io.CopyN(os.Stdout, r, 100)
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil

}

// defer file.CLose() to close
// defer func(){
// 	r.Close(ctx)
// }()

// a := 1
// {
// 	// a := 2 //shadows outer
// 	a= 2 //changes outer
// 	fmt.Println("inner",a) //affects only inner a
// }
// fmt.Println("outer", a)
// :=
