//
package http

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/model"
)

const (
	targetGithubURL = `^https://github.com\/[\w-]+\/[\w-]+`
	visitURL        = "https://github.com/avelino/awesome-go"
)

var (
	r = regexp.MustCompile(targetGithubURL)
)

// Collector struct
type Collector struct {
	colly         *colly.Collector
	repositoryMap map[string]model.Repositories
}

// NewCollector constructor
func NewCollector() *Collector {
	return &Collector{
		colly:         colly.NewCollector(),
		repositoryMap: make(map[string]model.Repositories),
	}
}

// RepositoryMap
func (c *Collector) RepositoryMap() map[string]model.Repositories {
	return c.repositoryMap
}

// BeforeRequest
func (c *Collector) BeforeRequest() {
	// before request
	c.colly.OnRequest(func(r *colly.Request) {
		fmt.Print(".")
	})
}

func (c *Collector) OnReadMe(texts []string) {
	// get content
	c.colly.OnHTML("#readme h2", func(e *colly.HTMLElement) {
		if !c.contains(texts, e.Text) {
			return
		}

		e.DOM.NextAllFiltered("ul").First().Find("li > a").Each(func(i int, selection *goquery.Selection) {
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
}

func (c *Collector) OnGithubStar() {
	// star数を取得
	c.colly.OnHTML("a.social-count.js-social-count", func(e *colly.HTMLElement) {
		// awesome goは除外する
		if e.Request.URL.Path == "/avelino/awesome-go" {
			return
		}

		star, err := c.getStar(e.Text)
		if err != nil {
			return
		}

		repo := model.Repository{
			PackageURL:  e.Request.URL.String(),
			Star:        star,
			Description: e.Request.Ctx.Get("description"),
		}

		contentName := e.Request.Ctx.Get("contentName")
		c.repositoryMap[contentName] = append(c.repositoryMap[contentName], repo)
	})
}

// ErrorListener
func (c *Collector) ErrorListener() {
	c.colly.OnError(func(_ *colly.Response, err error) {
		log.Println("something went wrong:", err)
	})
}

func (c *Collector) VisitAweSomeGo() error {
	return c.colly.Visit(visitURL)
}

func (c *Collector) contains(args []string, targetTxt string) bool {
	for _, arg := range args {
		if arg == targetTxt {
			return true
		}
	}
	return false
}

func (c *Collector) getStar(val string) (int, error) {
	starStr := strings.Replace(val, ",", "", -1)
	starStr = strings.Replace(starStr, "\n", "", -1)
	starStr = strings.Replace(starStr, " ", "", -1)

	star, err := strconv.Atoi(starStr)
	if err != nil {
		return 0, err
	}

	return star, err
}
