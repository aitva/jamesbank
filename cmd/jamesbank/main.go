package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type congressMan struct {
	FirstName string
	LastName  string
	Territory string
	Group     string
	Birth     string
	Job       string
	Title     string
}

func parseCongressMan(tr *goquery.Selection) congressMan {
	const prefix = "    "
	c := congressMan{}

	tds := tr.Children()
	names := strings.Split(tds.First().Children().First().Text(), ",")
	c.LastName = names[0]
	if len(names) > 1 {
		c.FirstName = names[1]
	}

	tds = tds.Next()
	c.Territory = tds.First().Text()

	tds = tds.Next() // jump empty

	tds = tds.Next()
	c.Group = tds.First().Text()

	tds = tds.Next()
	c.Birth = tds.First().Text()

	tds = tds.Next()
	c.Job = tds.First().Text()

	tds = tds.Next()
	c.Title = strings.Replace(tds.First().Text(), "\n", "; ", -1)
	return c
}

func congressManScrape() {
	doc, err := goquery.NewDocument("https://fr.wikipedia.org/wiki/Liste_des_d%C3%A9put%C3%A9s_de_la_XIVe_l%C3%A9gislature_de_la_Cinqui%C3%A8me_R%C3%A9publique")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("table.wikitable.sortable tr").Each(func(i int, s *goquery.Selection) {
		fmt.Println(goquery.NodeName(s))
		c := parseCongressMan(s)
		fmt.Println(c)
	})
}

func main() {
	congressManScrape()
}
