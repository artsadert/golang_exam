package task2

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func NewBrainrotMeme(name string, trendLevel int, category string, views float64) (*BrainrotMeme, error) {
	if trendLevel < 1 || trendLevel > 10 {
		return nil, fmt.Errorf("trend level must be between 1 and 10")
	}

	if category != "Subo Bratik" && category != "TUNTUNTUNSAHUR" && category != "Sigma" && category != "Skibidi" && category != "Mewing" && category != "Other" {
		return nil, fmt.Errorf("category must match correct ones")
	}

	return &BrainrotMeme{
		Name:       name,
		TrendLevel: trendLevel,
		Category:   category,
		Views:      views,
	}, nil
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var topTrending []BrainrotMeme

	for _, meme := range memes {
		if meme.Views >= minViews {
			topTrending = append(topTrending, meme)
		}
	}

	sort.Slice(topTrending, func(i, j int) bool {
		return topTrending[i].TrendLevel < topTrending[j].TrendLevel
	})

	return topTrending
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	categoryImpact := make(map[string]float64)

	for _, meme := range memes {
		categoryImpact[meme.Category] += meme.Views
	}

	return categoryImpact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var filteredMemes []string

	for _, meme := range memes {
		if meme.TrendLevel >= 7 || meme.Views > 50 || meme.Category == "Sigma" {
			filteredMemes = append(filteredMemes, meme.Name)
		}
	}

	return filteredMemes
}
