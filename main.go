package main

import (
	"flag"
	"log"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/output"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/scraper"
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

	sc := scraper.New(contentName)
	sc.ErrorListener()
	sc.BeforeRequest()
	sc.OnReadMe(contentName)
	sc.OnGithubStar()
	sc.OnDescription()
	sc.OnCompleted()
	if err := sc.VisitAweSomeGo(); err != nil {
		return err
	}
	sc.Wait()

	repos := sc.Repositories()
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
