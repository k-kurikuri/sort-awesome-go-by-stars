package main

import (
	"flag"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/http"
	"github.com/k-kurikuri/sort-awesome-go-by-stars/output"
)

func main() {
	err := realMain()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func realMain() error {
	flag.Parse()
	args := flag.Args()

	collect := http.NewCollector()
	collect.ErrorListener()
	collect.OnReadMe(args)
	collect.OnGithubStar()
	collect.BeforeRequest()
	if err := collect.VisitAweSomeGo(); err != nil {
		return err
	}

	for contentName, repositories := range collect.RepositoryMap() {
		repositories.SortDesc()
		topN := repositories.TopRankRepositories()

		table := output.NewTable(
			output.Header(simpletable.AlignCenter, "STAR", "PACKAGE_URL", "DESCRIPTION"),
			output.Footer(simpletable.AlignRight, contentName),
		)

		for _, repo := range topN {
			table.AddCells(repo.Star, repo.PackageURL, repo.Description)
		}

		table.Println()
	}

	return nil
}
