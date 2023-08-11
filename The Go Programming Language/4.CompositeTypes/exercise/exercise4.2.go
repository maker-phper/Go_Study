package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var str, sha string
	str = os.Args[1]
	sha = os.Args[2]

	if sha == "384" {
		fmt.Println(sha512.Sum384([]byte(str)))
	} else if sha == "512" {
		fmt.Println(sha512.Sum512([]byte(str)))
	} else {
		fmt.Println(sha256.Sum256([]byte(str)))
	}
}

//https://github.com/torbiak/gopl/blob/master/ex4.2/sha.go
// gopl/ex4.2/sha.go
// if debug it in Goland, add -w 512 in running configurations options "program arguments"
var width = flag.Int("w", 256, "hash width (256 or 512)")

func goplSha() {
	flag.Parse()
	var function func(b []byte) []byte
	switch *width {
	case 256:
		function = func(b []byte) []byte {
			h := sha256.Sum256(b)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width specified")
	}
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", function(b))
}
