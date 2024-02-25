package postgresql

import (
	"fmt"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewClient creates a new Gorm database client and performs auto-migration.
func NewClient(cfg *config.Config) (*gorm.DB, error) {
	// Database connection string
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", cfg.Database.Host, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)

	// Open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	// Perform auto-migration for defined models
	db.AutoMigrate(&models.Team{}, &models.Player{}, &models.Match{}, &models.Score{}, &models.Assist{})

	// Seed initial data into the database
	seedData(db)

	return db, nil
}

// seedData populates the database with initial example data if no teams are found.
func seedData(db *gorm.DB) {
	var existingTeam models.Team
	result := db.First(&existingTeam)
	if result.Error == nil {
		fmt.Println("An example team already exists in the database.")
		return
	}

	// Create example teams
	teams := []models.Team{
		{TeamName: "Team 1"},
		{TeamName: "Team 2"},
		{TeamName: "Team 3"},
		{TeamName: "Team 4"},
		{TeamName: "Team 5"},
		{TeamName: "Team 6"},
		{TeamName: "Team 7"},
		{TeamName: "Team 8"},
	}

	// Add teams to the database
	for i := range teams {
		db.Create(&teams[i])

		// Create and add 5 players for each team to the database
		for j := 1; j <= len(teams); j++ {
			player := models.Player{PlayerName: fmt.Sprintf("Player %d", (i*5)+j), TeamID: teams[i].ID}
			db.Create(&player)
		}
	}
}
