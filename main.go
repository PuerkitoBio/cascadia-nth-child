package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
	"github.com/andybalholm/cascadia"
)

func main() {
	f, err := os.Open("src.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	ss := "#tbl tr:nth-child(1) td"
	t := time.Now()
	sel := doc.Find(ss)
	log.Printf("%s: %d elements in %s via goquery", ss, sel.Size(), time.Since(t))

	m := cascadia.MustCompile(ss)
	t = time.Now()
	ns := m.MatchAll(doc.Nodes[0])
	log.Printf("%s: %d elements in %s via cascadia", ss, len(ns), time.Since(t))

	t = time.Now()
	sel2 := doc.Find("#tbl tr").First().Find("td")
	log.Printf("tr + First + td: %d elements in %s", sel2.Size(), time.Since(t))

	ns1, ns2 := sel.Nodes, sel2.Nodes
	if !sameNodes(ns, ns1) {
		log.Printf("cascadia nodes != goquery nodes (1)")
	}
	if !sameNodes(ns, ns2) {
		log.Printf("cascadia nodes != goquery nodes (2)")
	}
}

func sameNodes(ns1, ns2 []*html.Node) bool {
	if len(ns1) != len(ns2) {
		return false
	}
	for i, n := range ns1 {
		if n != ns2[i] {
			return false
		}
	}
	return true
}
