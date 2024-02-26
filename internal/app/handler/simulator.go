package handler

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/config"
	"github.com/acergiyen/simulator/internal/app/services"
)

type Simulator struct {
	lgr           *log.Logger
	cfg           *config.Config
	teamService   *services.TeamService
	matchService  *services.MatchService
	scoreService  *services.ScoreService
	assistService *services.AssistService
}

func NewSimulatorHandler(lgr *log.Logger, cfg *config.Config, teamService *services.TeamService, matchService *services.MatchService, scoreService *services.ScoreService, assistService *services.AssistService) *Simulator {
	return &Simulator{lgr: lgr, cfg: cfg, teamService: teamService, matchService: matchService, scoreService: scoreService, assistService: assistService}
}

func (s *Simulator) Simulate(w http.ResponseWriter) {
	teams, err := s.teamService.GetAllTeams()
	if err != nil {
		s.lgr.Printf("err: %v", err)
		return
	}

	// Shuffle teams randomly
	matches := s.createMatches(teams)

	// Start matches for both teams simultaneously
	s.startMatches(matches)
	htmlContent, err := ioutil.ReadFile("../../template/index.html")
	if err != nil {
		s.lgr.Printf("Error reading HTML file: %v", err)
		return
	}

	// Write HTTP response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	// Populate {{.Matches}} expression in HTML with match information
	matches, err = s.matchService.GetAllMatchesWithStatistics()
	if err != nil {
		s.lgr.Printf("Database error:%v", err)
	}
	templateData := struct{ Matches []models.Match }{Matches: matches}
	tmpl, err := template.New("index").Parse(string(htmlContent))
	if err != nil {
		s.lgr.Println("Error parsing HTML template:", err)
		return
	}
	tmpl.Execute(w, templateData)
}

func (s *Simulator) createMatches(teams []models.Team) []models.Match {
	var matches []models.Match

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })

	for i := 0; i < len(teams)-1; i += 2 {
		match := models.Match{
			HomeTeamID: teams[i].ID,
			AwayTeamID: teams[i+1].ID,
			StartTime:  time.Now(),
			Duration:   time.Duration(s.cfg.Simulator.UpdateCount) * time.Minute,
		}
		matches = append(matches, match)
	}

	s.matchService.AddMatches(matches)
	matches, _ = s.matchService.GetAllMatches()
	return matches
}

func (s *Simulator) startMatches(matches []models.Match) {
	// Create waitgroup to wait for completion
	var wg sync.WaitGroup

	// Counter to control total time
	totalTime := time.Duration(0)

	// Create a goroutine for each match
	for _, match := range matches {
		wg.Add(1)
		go func(match models.Match) {
			defer wg.Done()
			s.simulateMatch(match)
		}(match)

		// Update total time
		totalTime += match.Duration
	}

	// Wait for all matches to finish and control total time
	wg.Wait()

	// If total time exceeds 240 seconds, issue a warning
	if totalTime > time.Duration(240)*time.Second {
		s.lgr.Println("Total time exceeded 240 seconds!")
	}
}

func (s *Simulator) simulateMatch(match models.Match) {
	s.lgr.Printf("Match between Team %d and Team %d started at %s\n", match.HomeTeamID, match.AwayTeamID, match.StartTime.Format("15:04:05"))

	startTime := time.Now()
	updateCount := 0

	for {
		elapsedTime := time.Since(startTime)
		if updateCount >= s.cfg.Simulator.UpdateCount {
			s.lgr.Printf("Match between Team %d and Team %d ended at %s\n", match.HomeTeamID, match.AwayTeamID, time.Now().Format("15:04:05"))
			break
		}

		// Update score and assist every 5 seconds
		if elapsedTime.Seconds()-match.CurrentTime.Seconds() >= float64(s.cfg.Simulator.SimulateTime) {
			match.CurrentTime = elapsedTime
			s.updateScore(match)
			s.updateAssist(match)
			updateCount++
		}

		time.Sleep(1 * time.Second)
	}
}

func (s *Simulator) updateScore(match models.Match) {
	// Randomly select a player
	playerID := match.HomeTeam.Players[rand.Intn(len(match.HomeTeam.Players))].ID

	score := rand.Intn(2) + 2

	s.lgr.Printf("Match between Team %d and Team %d - Player %d scored %d points at %s\n", match.HomeTeamID, match.AwayTeamID, playerID, score, time.Now().Format("15:04:05"))
	// You can add score information to the database here.
	s.scoreService.AddScore(&models.Score{MatchID: match.ID, PlayerID: uint(playerID), Score: score})
}

func (s *Simulator) updateAssist(match models.Match) {
	// Randomly select a player
	playerID := match.HomeTeam.Players[rand.Intn(len(match.HomeTeam.Players))].ID
	assistCount := 1

	s.lgr.Printf("Match between Team %d and Team %d - Player %d made %d assists at %s\n", match.HomeTeamID, match.AwayTeamID, playerID, assistCount, time.Now().Format("15:04:05"))
	// You can add assist information to the database here.
	s.assistService.AddAssist(&models.Assist{MatchID: match.ID, PlayerID: uint(playerID), AssistCount: assistCount})
}
