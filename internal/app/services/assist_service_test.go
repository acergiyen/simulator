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

var _ = Describe("AssistService", func() {

	var (
		mockLogger *log.Logger
		mockRepo   *repositoriesfakes.FakeAssistRepositoryInterface
		assistSvc  *services.AssistService
	)
	BeforeEach(func() {
		var buff bytes.Buffer
		mockLogger = log.New(&buff, "", log.LstdFlags)

		mockRepo = new(repositoriesfakes.FakeAssistRepositoryInterface)
		assistSvc = services.NewAssistService(mockLogger, mockRepo)
	})

	Describe("AddAssist", func() {

		Context("when repository returns no error", func() {
			It("should not return an error", func() {
				mockRepo.AddAssistReturns(nil) // Simulate repository success
				err := assistSvc.AddAssist(&models.Assist{})
				Expect(err).To(BeNil())
			})
		})

		Context("when repository returns an error", func() {
			It("should return an error", func() {
				mockRepo.AddAssistReturns(errors.New("some error")) // Simulate repository error
				err := assistSvc.AddAssist(&models.Assist{Model: gorm.Model{ID: 1}})
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
