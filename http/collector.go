//
package http

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/model"
)

const (
	targetGithubURL = `^https://github.com\/[\w-]+\/[\w-]+`
	visitURL        = "https://github.com/avelino/awesome-go"
)

var (
	r     = regexp.MustCompile(targetGithubURL)
	mutex = &sync.Mutex{}
)

// Collector struct
type Collector struct {
	colly          *colly.Collector
	repositories   model.Repositories
	starMap        *sync.Map
	descriptionMap *sync.Map
	contentName    string
}

// NewCollector constructor
func NewCollector(contentName string) *Collector {
	c := colly.NewCollector(
		colly.Async(true),
		colly.AllowedDomains("github.com"),
	)
	return &Collector{
		colly:          c,
		repositories:   make(model.Repositories, 0, model.Rank),
		starMap:        &sync.Map{},
		descriptionMap: &sync.Map{},
		contentName:    contentName,
	}
}

// Repositories
func (c *Collector) Repositories() model.Repositories {
	return c.repositories
}

// BeforeRequest
func (c *Collector) BeforeRequest() {
	// before request
	c.colly.OnRequest(func(r *colly.Request) {
		fmt.Print(".")
	})
}

func (c *Collector) OnReadMe(text string) {
	// get content
	c.colly.OnHTML("#readme h2", func(e *colly.HTMLElement) {
		if text != e.Text {
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

			err := e.Request.Visit(href)
			if err != nil {
				log.Println(err)
				return
			}
		})
	})
}

func (c *Collector) OnGithubStar() {
	c.colly.OnHTML("a.social-count.js-social-count", func(e *colly.HTMLElement) {
		// awesome goは除外する
		if e.Request.URL.Path == "/avelino/awesome-go" {
			return
		}

		star, err := c.getStarFromScrapedText(e.Text)
		if err != nil {
			log.Println(err)
			return
		}

		packageURL := e.Request.URL.String()
		c.starMap.Store(packageURL, star)
	})
}

func (c *Collector) OnDescription() {
	c.colly.OnHTML("span.text-gray-dark.mr-2", func(e *colly.HTMLElement) {
		packageURL := e.Request.URL.String()
		c.descriptionMap.Store(packageURL, strings.Replace(e.Text, "\n", "", -1))
	})
}

func (c *Collector) OnCompleted() {
	c.colly.OnScraped(func(res *colly.Response) {
		packageURL := res.Request.URL.String()
		repo := model.Repository{
			PackageURL:  packageURL,
			Star:        c.getStar(packageURL),
			Description: c.getDescription(packageURL),
		}

		mutex.Lock()
		c.repositories = append(c.repositories, repo)
		mutex.Unlock()
	})
}

// ErrorListener
func (c *Collector) ErrorListener() {
	c.colly.OnError(func(_ *colly.Response, err error) {
		log.Println(err)
	})
}

func (c *Collector) VisitAweSomeGo() error {
	return c.colly.Visit(visitURL)
}

func (c *Collector) Wait() {
	c.colly.Wait()
}

func (c *Collector) getStarFromScrapedText(val string) (int, error) {
	starStr := strings.Replace(val, ",", "", -1)
	starStr = strings.Replace(starStr, "\n", "", -1)
	starStr = strings.Replace(starStr, " ", "", -1)

	switch {
	case strings.Contains(starStr, "."):
		starStr = strings.Replace(starStr, ".", "", 1)
		starStr = strings.Replace(starStr, "k", "00", 1)
	case strings.Contains(starStr, "k"):
		starStr = strings.Replace(starStr, "k", "000", 1)
	}

	star, err := strconv.Atoi(starStr)
	if err != nil {
		return 0, err
	}

	return star, err
}

func (c *Collector) getStar(key string) int {
	val, ok := c.starMap.Load(key)
	if !ok {
		return 0
	}
	return val.(int)
}

func (c *Collector) getDescription(key string) string {
	val, ok := c.descriptionMap.Load(key)
	if !ok {
		return ""
	}
	return val.(string)
}
