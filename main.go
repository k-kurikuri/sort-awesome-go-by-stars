package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexeyco/simpletable"
	"github.com/gocolly/colly"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/model"
)

const (
	visitURL        = "https://github.com/avelino/awesome-go"
	targetGithubURL = `^https://github.com\/[\w-]+\/[\w-]+`
	order           = 5
)

var (
	c          = colly.NewCollector()
	r          = regexp.MustCompile(targetGithubURL)
	packageMap = make(map[string]model.Repositories)
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	args := flag.Args()

	// get content
	c.OnHTML("#readme h2", func(e *colly.HTMLElement) {
		if !contains(args, e.Text) {
			return
		}

		e.DOM.Next().NextAllFiltered("ul").First().Find("li > a").Each(func(i int, selection *goquery.Selection) {
			href, exists := selection.Attr("href")
			if !exists {
				return
			}

			if !r.MatchString(href) {
				return
			}

			e.Request.Ctx.Put("contentName", e.Text)
			e.Request.Ctx.Put("description", selection.Parent().Text())
			err := e.Request.Visit(href)
			if err != nil {
				return
			}
		})
	})

	// star数を取得
	c.OnHTML("a.social-count.js-social-count", func(e *colly.HTMLElement) {
		// awesome goは除外する
		if e.Request.URL.Path == "/avelino/awesome-go" {
			return
		}

		star, err := getStar(e.Text)
		if err != nil {
			return
		}

		repo := model.Repository{
			URL:         e.Request.URL.String(),
			Star:        star,
			Description: e.Request.Ctx.Get("description"),
		}

		contentName := e.Request.Ctx.Get("contentName")
		packageMap[contentName] = append(packageMap[contentName], repo)
	})

	// before request
	c.OnRequest(func(r *colly.Request) {
		fmt.Print(".")
	})

	// error listener
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("something went wrong:", err)
	})

	err := c.Visit(visitURL)
	if err != nil {
		return 1
	}

	for k, repositories := range packageMap {
		sort.Slice(repositories, func(i, j int) bool {
			return repositories[i].Star > repositories[j].Star
		})

		var topN model.Repositories
		switch {
		case len(repositories) >= order:
			topN = repositories[0:order]
		default:
			topN = repositories[0:]
		}

		table := simpletable.New()
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "STAR"},
				{Align: simpletable.AlignCenter, Text: "PACKAGE_URL"},
				{Align: simpletable.AlignCenter, Text: "DESCRIPTION"},
			},
		}

		for _, repo := range topN {
			r := []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", repo.Star)},
				{Text: repo.URL},
				{Text: repo.Description},
			}

			table.Body.Cells = append(table.Body.Cells, r)
		}

		table.Footer = &simpletable.Footer{
			Cells: []*simpletable.Cell{
				{},
				{},
				{Align: simpletable.AlignRight, Text: k},
			},
		}

		fmt.Println()
		table.SetStyle(simpletable.StyleUnicode)
		table.Println()
	}

	return 0
}

func getStar(val string) (int, error) {
	starStr := strings.Replace(val, ",", "", -1)
	starStr = strings.Replace(starStr, "\n", "", -1)
	starStr = strings.Replace(starStr, " ", "", -1)

	star, err := strconv.Atoi(starStr)
	if err != nil {
		return 0, err
	}

	return star, err
}

func contains(args []string, targetTxt string) bool {
	for _, arg := range args {
		if arg == targetTxt {
			return true
		}
	}
	return false
}
