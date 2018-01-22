package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-mimetypes"
	"log"
	"os"
	"strings"
)

func main() {

	var lookup = flag.String("lookup", "extension", "Valid options are to lookup by: extension; mimetype")

	flag.Parse()

	for _, input := range flag.Args() {

		switch *lookup {

		case "extension":
			t := mimetypes.TypesByExtension(input)
			fmt.Printf("%s\t%s\n", input, strings.Join(t, "\t"))
		case "mimetype":
			e := mimetypes.ExtensionsByType(input)
			fmt.Printf("%s\t%s\n", input, strings.Join(e, "\t"))
		default:
			log.Fatal("Invalid lookup type")
		}
	}

	os.Exit(0)
}
