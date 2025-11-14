package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	goals = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "fbref_team_goals_total", Help: "Total goals scored by each Premier League team"},
		[]string{"team"},
	)
	possession = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "fbref_team_possession_percent", Help: "Average possession percentage by team"},
		[]string{"team"},
	)
	shots = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "fbref_team_shots_total", Help: "Total shots taken by each team"},
		[]string{"team"},
	)
	shotsOnTarget = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "fbref_team_shots_on_target", Help: "Shots on target by each team"},
		[]string{"team"},
	)
	passCompletion = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "fbref_team_pass_completion_percent", Help: "Pass completion percentage by team"},
		[]string{"team"},
	)
	expectedGoals = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "fbref_team_expected_goals", Help: "Expected goals (xG) by team"},
		[]string{"team"},
	)
)

func init() {
	prometheus.MustRegister(goals, possession, shots, shotsOnTarget, passCompletion, expectedGoals)
}

// ----------- FIXED SCRAPER -------------
func scrapeFBref() {
	url := "https://fbref.com/en/comps/9/Premier-League-Stats"

	client := &http.Client{Timeout: 15 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		log.Println("Error fetching FBref:", err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return
	}

	// Updated selector — FBref moved the table!
	table := doc.Find("table#stats_squads_standard_for")
	if table.Length() == 0 {
		table = doc.Find("table#stats_squads_standard_9")
	}
	if table.Length() == 0 {
		table = doc.Find("table#stats_squads_standard") // fallback
	}

	if table.Length() == 0 {
		log.Println("ERROR: No stats table found — FBref likely changed layout again")
		return
	}

	rowCount := 0

	table.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		team := strings.TrimSpace(s.Find("th[data-stat='team']").Text())
		if team == "" {
			return
		}
		rowCount++

		parseFloat := func(v string) float64 {
			v = strings.TrimSpace(strings.TrimSuffix(v, "%"))
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0
			}
			return f
		}

		goals.WithLabelValues(team).Set(parseFloat(s.Find("td[data-stat='goals']").Text()))
		possession.WithLabelValues(team).Set(parseFloat(s.Find("td[data-stat='possession']").Text()))
		shots.WithLabelValues(team).Set(parseFloat(s.Find("td[data-stat='shots']").Text()))
		shotsOnTarget.WithLabelValues(team).Set(parseFloat(s.Find("td[data-stat='shots_on_target']").Text()))
		passCompletion.WithLabelValues(team).Set(parseFloat(s.Find("td[data-stat='passes_pct']").Text()))
		expectedGoals.WithLabelValues(team).Set(parseFloat(s.Find("td[data-stat='xg']").Text()))
	})

	if rowCount == 0 {
		log.Println("WARNING: Scraper ran but no team rows were parsed — selectors may need updating")
	} else {
		log.Printf("Scraped %d teams successfully\n", rowCount)
	}
}

// -------------- MAIN ----------------
func main() {
	// Background scrape loop
	go func() {
		for {
			log.Println("Running FBref scrape…")
			scrapeFBref()
			time.Sleep(5 * time.Minute) // use 5m for debugging, then change back to 15m
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Exporter running on :9101/metrics")
	log.Fatal(http.ListenAndServe(":9101", nil))
}
