package test

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/itsHardStyl3r/the-swift-codes/cmd/api"
	"github.com/itsHardStyl3r/the-swift-codes/internal/models"
	"github.com/itsHardStyl3r/the-swift-codes/internal/tools"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type APITestSuite struct {
	suite.Suite
	DB   *gorm.DB
	gin  *gin.Engine
	path string
}

func (s *APITestSuite) SetupSuite() {
	s.path = "/v1/swift-codes"
	var err error
	if tools.DB, err = gorm.Open(sqlite.Open(":memory:")); err != nil {
		log.Fatalf("Failed to connect to database: %v.", err)
	}
	s.DB = tools.DB
	if err = tools.SetupDb(true, false); err != nil {
		log.Fatalf("Failed to setup database: %v.", err)
	}
	gin.SetMode(gin.TestMode)
	s.gin = gin.Default()
	s.populateWithMock()
	v1 := s.gin.Group("/v1")
	api.DeleteBySwiftCode(v1)
	tools.LogDatabaseStats()
}

func (s *APITestSuite) TearDownSuite() {
	if s.DB != nil {
		db, err := s.DB.DB()
		if err != nil {
			log.Errorf("DB() failed: %v.", err)
			return
		}
		if err = db.Close(); err != nil {
			log.Error("Failed to close database connection.")
			return
		}
	}
}

func TestAPI(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) populateWithMock() {
	s.DB.Create(&models.Country{Name: "POLAND", Iso2: "PL"})
	s.DB.Create(&models.Country{Name: "LATVIA", Iso2: "LT"})
	s.DB.Create(&models.Country{Name: "MALTA", Iso2: "MT"})

	s.DB.Create(&models.Bank{Name: "Bank of Poland", BankCode: "POLS"})
	s.DB.Create(&models.Bank{Name: "Bank of Latvia", BankCode: "LITW"})
	s.DB.Create(&models.Bank{Name: "Bank of Malta", BankCode: "MALT"})

	s.DB.Create(&models.Bic{
		CountryId:    1,
		Bic:          "POLSPLAWXXX",
		BankId:       1,
		Address:      "Some address",
		Town:         "WARSAW",
		LocationCode: "AW",
		Branch:       "XXX",
	})

	s.DB.Create(&models.Bic{
		CountryId:    2,
		Bic:          "LITWLTDEXXX",
		BankId:       2,
		Address:      "Some address",
		Town:         "Riga",
		LocationCode: "DE",
		Branch:       "XXX",
	})

	s.DB.Create(&models.Bic{
		CountryId:    3,
		Bic:          "MALTMTIPR2T",
		BankId:       3,
		Address:      "Some address",
		Town:         "Birgu",
		LocationCode: "IP",
		Branch:       "R2T",
	})
}
