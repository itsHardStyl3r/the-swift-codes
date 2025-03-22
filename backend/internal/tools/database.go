package tools

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/itsHardStyl3r/the-swift-codes/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDb() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("dbUser"),
		os.Getenv("dbPassword"),
		os.Getenv("dbAddr"),
		os.Getenv("dbDatabase"),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	return nil
}

func SetupDb(shouldAutoMigrate bool) error {
	if shouldAutoMigrate {
		err := DB.AutoMigrate(&models.Country{}, &models.Bank{}, &models.Bic{})
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
			return err
		}
	}

	var dataPath = os.Getenv("csvDataPath")
	records := ReadCsv(dataPath, true)
	if len(records) == 0 {
		log.Warnf("File '%v' appears to be empty.", dataPath)
	}

	var count int64
	if (DB.Migrator().HasTable(&models.Country{})) {
		DB.Model(&models.Country{}).Count(&count)
		if count != 0 {
			log.Info("Found non-empty table 'countries'. No data added.")
		} else {
			log.Info("Populating table 'countries'...")
			var countries []models.Country
			countrySet := make(map[string]struct{})
			for _, record := range records {
				iso2 := record[0]
				if _, exists := countrySet[iso2]; !exists {
					countries = append(countries, models.Country{
						Name: record[6],
						Iso2: iso2,
					})
					countrySet[iso2] = struct{}{}
				}
			}
			err := addDataTransaction(countries)
			if err != nil {
				log.Warnf("Failed to add counties via transaction. %v", err)
			}
		}
	}

	if (DB.Migrator().HasTable(&models.Bank{})) {
		DB.Model(&models.Bank{}).Count(&count)
		if count != 0 {
			log.Info("Found non-empty table 'banks'. No data added.")
		} else {
			log.Info("Populating table 'banks'...")
			var banks []models.Bank
			bankSet := make(map[string]struct{})
			for _, record := range records {
				bankCode := record[1][0:4]
				if _, exists := bankSet[bankCode]; !exists {
					banks = append(banks, models.Bank{
						Name:     record[3],
						BankCode: bankCode,
					})
					bankSet[bankCode] = struct{}{}
				}
			}
			err := addDataTransaction(banks)
			if err != nil {
				log.Warnf("Failed to add banks via transaction. %v", err)
			}
		}
	}

	if (DB.Migrator().HasTable(&models.Bic{})) {
		DB.Model(&models.Bic{}).Count(&count)
		if count != 0 {
			log.Info("Found non-empty table 'bics'. No data added.")
		} else {
			log.Info("Populating table 'bics'...")
			countryMap := loadCountryMap()
			bankMap := loadBankMap()
			var bics []models.Bic
			for _, record := range records {
				bankCode := record[1][:4]
				countryIso := record[0]
				bankId, bankExists := bankMap[bankCode]
				countryId, countryExists := countryMap[countryIso]
				if !bankExists {
					log.Warnf("Bank with bankcode '%s' not found. Skipping %s.", bankCode, record[1])
					continue
				}
				if !countryExists {
					log.Warnf("Country with iso2 '%s' not found. Skipping %s.", countryIso, record[1])
					continue
				}
				branch := "XXX"
				if record[1][8:11] != "XXX" {
					branch = record[1][8:11]
				}
				bics = append(bics, models.Bic{
					Bic:          record[1],
					BankId:       bankId,
					CountryId:    countryId,
					LocationCode: record[1][6:8],
					Branch:       branch,
					Address:      record[4],
					Town:         record[5],
					TimeZone:     record[7],
					CodeType:     record[2],
				})
			}
			err := addDataTransaction(bics)
			if err != nil {
				log.Warnf("Failed to add bics via transaction. %v", err)
			}
		}
	}
	return nil
}

func addDataTransaction[T models.Country | models.Bank | models.Bic](slice []T) error {
	tx := DB.Begin()
	err := tx.Error
	if err != nil {
		return err
	}
	if err = tx.Create(slice).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func LogDatabaseStats() {
	log.Info("Current database status:")
	var count int64
	DB.Model(&models.Country{}).Count(&count)
	log.Infof("- countries: %d entries", count)
	DB.Model(&models.Bank{}).Count(&count)
	log.Infof("- banks: %d entries", count)
	DB.Model(&models.Bic{}).Count(&count)
	log.Infof("- bics: %d entries", count)
}

func loadCountryMap() map[string]int {
	var countries []models.Country
	DB.Model(&models.Country{}).Select("id, iso2").Find(&countries)

	countryMap := make(map[string]int)
	for _, country := range countries {
		countryMap[country.Iso2] = country.Id
	}
	return countryMap
}

func loadBankMap() map[string]int {
	var banks []models.Bank
	DB.Model(&models.Bank{}).Select("id, bank_code").Find(&banks)

	bankMap := make(map[string]int)
	for _, bank := range banks {
		bankMap[bank.BankCode] = bank.Id
	}
	return bankMap
}
