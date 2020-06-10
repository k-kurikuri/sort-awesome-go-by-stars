package model

import (
	"sort"
)

const (
	Rank = 5
)

// Repository struct
type Repository struct {
	PackageURL  string
	Star        int
	Description string
}

// Repositories Repository list
type Repositories []Repository

// SortDesc sort by desc
func (rs Repositories) SortDesc() {
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].Star > rs[j].Star
	})
}

// TopRankRepositories return top N repository list
func (rs Repositories) TopRankRepositories() Repositories {
	switch {
	case len(rs) >= Rank:
		return rs[0:Rank]
	default:
		return rs[0:]
	}
}
