package services_test

import (
	"bytes"
	"errors"
	"log"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/repositories/repositoriesfakes"
	"github.com/acergiyen/simulator/internal/app/services"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TeamService", func() {

	var (
		mockLogger *log.Logger
		mockRepo   *repositoriesfakes.FakeTeamRepositoryInterface
		teamSvc    *services.TeamService
	)

	BeforeEach(func() {
		var buff bytes.Buffer
		mockLogger = log.New(&buff, "", log.LstdFlags)

		mockRepo = new(repositoriesfakes.FakeTeamRepositoryInterface)
		teamSvc = services.NewTeamService(mockLogger, mockRepo)
	})

	Describe("GetAllTeams", func() {
		Context("when repository returns teams", func() {
			It("should return teams", func() {
				// Arrange
				expectedTeams := []models.Team{{Model: gorm.Model{ID: 1}}, {Model: gorm.Model{ID: 2}}}
				mockRepo.GetAllTeamsReturns(expectedTeams, nil)

				// Act
				teams, err := teamSvc.GetAllTeams()

				// Assert
				Expect(err).To(BeNil())
				Expect(teams).To(Equal(expectedTeams))
			})
		})

		Context("when repository returns an error", func() {
			It("should return an error", func() {
				// Arrange
				mockRepo.GetAllTeamsReturns(nil, errors.New("some error"))

				// Act
				teams, err := teamSvc.GetAllTeams()

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(teams).To(BeNil())
			})
		})
	})
})
