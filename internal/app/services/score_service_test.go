package services_test

import (
	"bytes"
	"errors"
	"log"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/repositories/repositoriesfakes"
	"github.com/acergiyen/simulator/internal/app/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScoreService", func() {

	var (
		mockLogger *log.Logger
		mockRepo   *repositoriesfakes.FakeScoreRepositoryInterface
		scoreSvc   *services.ScoreService
	)

	BeforeEach(func() {
		var buff bytes.Buffer
		mockLogger = log.New(&buff, "", log.LstdFlags)
		mockRepo = new(repositoriesfakes.FakeScoreRepositoryInterface)
		scoreSvc = services.NewScoreService(mockLogger, mockRepo)
	})

	Describe("AddScore", func() {

		Context("when repository returns no error", func() {
			It("should not return an error", func() {
				// Arrange
				mockRepo.AddScoreReturns(nil) // Simulate repository success

				// Act
				err := scoreSvc.AddScore(&models.Score{})

				// Assert
				Expect(err).To(BeNil())
			})
		})

		Context("when repository returns an error", func() {
			It("should return an error", func() {
				// Arrange
				mockRepo.AddScoreReturns(errors.New("some error")) // Simulate repository error

				// Act
				err := scoreSvc.AddScore(&models.Score{})

				// Assert
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
