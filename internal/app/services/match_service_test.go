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

var _ = Describe("MatchService", func() {

	var (
		mockRepo   *repositoriesfakes.FakeMatchRepositoryInterface
		matchSvc   *services.MatchService
		mockLogger *log.Logger
	)

	BeforeEach(func() {
		var buff bytes.Buffer
		mockLogger = log.New(&buff, "", log.LstdFlags)

		mockRepo = new(repositoriesfakes.FakeMatchRepositoryInterface)
		matchSvc = services.NewMatchService(mockLogger, mockRepo)
	})

	Describe("AddMatches", func() {
		Context("when repository returns no error", func() {
			It("should not return an error", func() {
				mockRepo.AddMatchReturns(nil) // Simulate repository success
				err := matchSvc.AddMatches([]models.Match{{}, {}})
				Expect(err).To(BeNil())
			})
		})

		Context("when repository returns an error", func() {
			It("should return an error", func() {
				mockRepo.AddMatchReturns(errors.New("some error")) // Simulate repository error
				err := matchSvc.AddMatches([]models.Match{{}, {}})
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("GetAllMatches", func() {
		Context("when repository returns matches", func() {
			It("should return matches", func() {
				mockRepo.GetAllMatchesReturns([]models.Match{{Model: gorm.Model{ID: 1}}, {Model: gorm.Model{ID: 2}}}, nil) // Simulate repository success
				matches, err := matchSvc.GetAllMatches()
				Expect(err).To(BeNil())
				Expect(matches).To(HaveLen(2))
			})
		})

		Context("when repository returns an error", func() {
			It("should return an error", func() {
				mockRepo.GetAllMatchesReturns(nil, errors.New("some error")) // Simulate repository error
				matches, err := matchSvc.GetAllMatches()
				Expect(err).To(HaveOccurred())
				Expect(matches).To(BeNil())
			})
		})
	})

	Describe("GetAllMatchesWithStatistics", func() {
		Context("when repository returns matches with statistics", func() {
			It("should return matches with statistics", func() {
				mockRepo.GetAllMatchesWithStatisticsReturns([]models.Match{{}, {}}, nil) // Simulate repository success
				matches, err := matchSvc.GetAllMatchesWithStatistics()
				Expect(err).To(BeNil())
				Expect(matches).To(HaveLen(2))
			})
		})

		Context("when repository returns an error", func() {
			It("should return an error", func() {
				mockRepo.GetAllMatchesWithStatisticsReturns(nil, errors.New("some error")) // Simulate repository error
				matches, err := matchSvc.GetAllMatchesWithStatistics()
				Expect(err).To(HaveOccurred())
				Expect(matches).To(BeNil())
			})
		})
	})
})
