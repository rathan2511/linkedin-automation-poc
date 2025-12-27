package search

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// SearchPeopleAndCollectProfiles collects profile URLs from FIRST SEARCH PAGE
func SearchPeopleAndCollectProfiles(page *rod.Page, keyword string) []string {
	fmt.Println("🔍 Searching people for keyword:", keyword)

	keyword = strings.ReplaceAll(keyword, " ", "%20")
	searchURL := "https://www.linkedin.com/search/results/people/?keywords=" + keyword

	page.MustNavigate(searchURL)

	// LinkedIn search NEVER stabilizes – use human waits
	time.Sleep(6 * time.Second)

	// Scroll multiple times to force lazy loading
	for i := 0; i < 3; i++ {
		page.Mouse.MustScroll(0, 900)
		time.Sleep(2 * time.Second)
	}

	// THIS is the key container for people search results
	results, err := page.Elements(`li.reusable-search__result-container`)
	if err != nil || len(results) == 0 {
		fmt.Println("❌ No search result containers found")
		return nil
	}

	profileSet := make(map[string]bool)
	var profiles []string

	for _, result := range results {
		link, err := result.Element(`a[href*="/in/"]`)
		if err != nil {
			continue
		}

		href, err := link.Attribute("href")
		if err != nil || href == nil {
			continue
		}

		url := *href

		// Clean tracking params
		if strings.Contains(url, "?") {
			url = strings.Split(url, "?")[0]
		}

		if !strings.HasPrefix(url, "https://") {
			url = "https://www.linkedin.com" + url
		}

		if !profileSet[url] {
			profileSet[url] = true
			profiles = append(profiles, url)
		}
	}

	fmt.Println("✅ Profiles collected:", len(profiles))
	return profiles
}
