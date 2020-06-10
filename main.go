package main

import (
	"flag"
	"log"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/http"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/output"
)

func main() {
	err := realMain()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func realMain() error {
	flag.Parse()
	contentName := flag.Arg(0)

	c := http.NewCollector(contentName)
	c.ErrorListener()
	c.BeforeRequest()
	c.OnReadMe(contentName)
	c.OnGithubStar()
	c.OnDescription()
	c.OnCompleted()
	if err := c.VisitAweSomeGo(); err != nil {
		return err
	}
	c.Wait()

	repos := c.Repositories()
	repos.SortDesc()
	topN := repos.TopRankRepositories()

	table := output.NewTable(
		output.Header(simpletable.AlignCenter, "STAR", "PACKAGE_URL", "DESCRIPTION"),
		output.Footer(simpletable.AlignRight, contentName),
	)

	for _, repo := range topN {
		table.AddCells(repo.Star, repo.PackageURL, repo.Description)
	}

	table.Println()

	return nil
}
