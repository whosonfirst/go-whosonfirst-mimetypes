package main

import (
	"bufio"
	"encoding/json"
	_ "flag"
	"log"
	"net/http"
	"regexp"
	"strings"
)

/*
# MIME type (lowercased)			Extensions
# ============================================	==========
# application/1d-interleaved-parityfec
# application/3gpp-ims+xml
# application/activemessage
application/andrew-inset			ez
*/

func main() {

	by_type := make(map[string][]string)
	by_ext := make(map[string][]string)

	re, err := regexp.Compile(`\s{2,}`)

	if err != nil {
		log.Fatal(err)
	}

	url := "https://svn.apache.org/viewvc/httpd/httpd/branches/2.2.x/docs/conf/mime.types?view=co"

	rsp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer rsp.Body.Close()

	scanner := bufio.NewScanner(rsp.Body)

	for scanner.Scan() {
	
		ln := scanner.Text()

		if strings.HasPrefix(ln, "#") {
			continue
		}

		ln = re.ReplaceAllString(ln, " ")
		parts := strings.Split(ln, " ")

		m := parts[0]
		e := parts[1:]

		by_type[m] = e

		for _, x := range e {

			_, ok := by_ext[x]

			if ok {
				by_ext[x] = append(by_ext[x], m)
			} else {
				by_ext[x] = []string{m}
			}
		}

		log.Println(m, len(e))
	}

	// TO DO : append missing here
	
	bt, _ := json.Marshal(by_type)
	log.Println(string(bt))

	be, _ := json.Marshal(by_ext)
	log.Println(string(be))
}
