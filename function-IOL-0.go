package main
package functionFunction
//package function
//package Function

import (
	"fmt"
	//"io"
	"net/http"
	//"os"
	"strings"
	//)
	"testing"
	//)
	"https://raw.githubusercontent.com/classicvalues/WOLFRAM-iolFunction/base/IOL.wl"
	//"github.com/LaudateCorpus1/WOLFRAM-iolFunction/"
	//"github.com/LaudateCorpus1/WOLFRAM-iolFunction/function"
	//"github.com/LaudateCorpus1/WOLFRAM-iolFunction/functionFunction"
	//)
	//"golang.org/x/net/html"
)

//func main() {
//	resp, err := http.Get("http://golang.org/")
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
//		os.Exit(1)
//	}

//var (
//	n = flag.Int("n", 1, "number of links")
//	url = flag.String("url", "%^&*", "url")
//)

func main(IOL) {
	resp, _ :=  http.Get("https://raw.githubusercontent.com/classicvalues/WOLFRAM-iolFunction/base/IOL.wl")
	//doc, err := html.Parse(resp.Body)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	//	os.Exit(1)
	//}
	//for _, link := range visit(nil, doc) {
	//	fmt.Println(link)
	//}
	//fmt.Println(resp.Body)}
	links := collectlinks.All(resp.html)
	fmt.Println(links)
}

//FAKE FUNCTION

//func main(IOL) {
//	resp, _ :=  http.Get("
//	doc, err := html.Parse(resp.Body)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
//		os.Exit(1)
//	}
//	for _, link := range visit(nil, doc) {
//		fmt.Println(link)
//	}
//	fmt.Println(resp.Body)}
//	links := collectlinks.All(resp.html)
//	fmt.Println(links)
//}

//CLEANER FUNCTION
//func main(IOL) {
//	resp, _ :=  http.Get("
//	doc, err := html.Parse(resp.Body)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
//		os.Exit(1)
//	}

func functionFunction(t *testing.T) {
	//resp, _ :=  http.Get("
	//doc, err := html.Parse(resp.Body)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	//	os.Exit(1)
	//}

	//for _, link := range visit(nil, doc) {
	//	fmt.Println(link)
	//}
	//fmt.Println(resp.Body)}
	//links := collectlinks.All(resp.html)
	//fmt.Println(links)

	//resp, _ :=  http.Get("
	//doc, err := html.Parse(resp.Body)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	//	os.Exit(1)
	//}

	//A paragraph of text</p>`) AFTER ('<p>   _______')
  reader := strings.NewReader(` <p> 
  <a href="https://raw.githubusercontent.com/classicvalues/WOLFRAM-iolFunction/base/IOL.wl"> n.th </a>
  https://raw.githubusercontent.com/classicvalues/WOLFRAM-iolFunction/base/IOL.wl
</p>`)
	//doc, err := html.Parse(reader)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	//	os.Exit(1)
	//}
	//for _, link := range visit(nil, doc) {
	//	fmt.Println(link)
	//}
	links := All(reader)//*strings.NewReader(IOL), OR, //*strings.Split()/*resp.html*/)
  fmt.Println(links)
//}
 // n := FIX FOR ERROR BELOW
  if len(links) != n {
	//t.Errorf("found %d links, want %d", len(links), n)
	//}
    t.Error("Error")
  }
}

func All(reader *strings.Reader) {
	panic("unimplemented")
}

//func Test_functionFunction(t *testing.T) {
//reader := strings.NewReader(` <p>
//<a href="
//"
//
//
//  import (
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//	"strings"
//	"github.com/LaudateCorpus1/WOLFRAM-iolFunction/"
//	"golang.org/x/net/html"
//)
//OR
//	import ( 
//	"crypto/tls"
//	"flag"
//	"fmt"
//	"io"
//	"strings"
//  "testing"
//	"github.com/LaudateCorpus1/WOLFRAM-iolFunction/"
//	"net/http"
//	"net/url"
//	"os"
//)


//package main
//package function
//package Function
//package functionFunction

//import (
//	"crypto/tls"
//	"flag"
//	"fmt"
//	"io"
//	"strings"
//	"testing"
//	"github.com/LaudateCorpus1/WOLFRAM-iolFunction/"
//	"net/http"
//	"net/url"
//	"os"
//)

//func main(IOL) {
//	resp, _ :=  http.Get("https://github.com/LaudateCorpus1WOLFRAM-iolFunction/")
//	links := collectlinks.All(resp.html)
//	fmt.Println(links)
//}

//func functionFunction(t *testing.T) {
//  reader := strings.NewReader(` <p>
//  <a href="https://github.com/LaudateCorpus1WOLFRAM-iolFunction"> n.th </a>
//  https://github.com/LaudateCorpus1WOLFRAM-iolFunction
//</p>`)
//
//  links := All(reader)
//
//  fmt.Println(links)
//
//  if len(links) != n {
//    t.Error("Error")
// }
//}







//FIRST FUNCTION VERSION
//------------------------------------------------------------------------------//

//package main
//package function
//package Function
//package functionFunction
//
//import (
//	"crypto/tls"
//	"flag"
//	"fmt"
//	"io"
//	"strings"
//  	"testing"
//	"github.com/LaudateCorpus1/WOLFRAM-iolFunction/"
//	"net/http"
//	"net/url"
//	"os"
//)
//
//func main(IOL) {
//	resp, _ :=  http.Get("https://github.com/LaudateCorpus1WOLFRAM-iolFunction/")
//	links := collectlinks.All(resp.html)
//	fmt.Println(links)
//}
//
//func functionFunction(t *testing.T) {
//  reader := strings.NewReader(` <p>
//  <a href="https://github.com/LaudateCorpus1WOLFRAM-iolFunction"> n.th </a>
//  https://github.com/LaudateCorpus1WOLFRAM-iolFunction
//</p>`)
//
//  links := All(reader)
//
//  fmt.Println(links)
//
//  if len(links) != n {
//    t.Error("Error")
//  }
//}
