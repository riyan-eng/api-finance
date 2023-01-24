package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	// var sqlDB *sql.DB
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TIMEZONE"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "finance.",
			SingularTable: false,
		},
	})

	if err != nil {
		panic("Connection failed to Database")
	}

	// connection pool
	// sqlDB, err = DB.DB()
	// sqlDB.SetConnMaxIdleTime(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic("Connection pool failed")
	}

	fmt.Println("Connection Opened to Database")

	// migrate database
	DB.AutoMigrate(&COA{}, &GeneralLedger{}, &Transaction{}, &CashReceiptJournal{}, &CashPaymentJuornal{}, &SalesJournal{}, &PurchaseJournal{}, &AdjustmentJournalEntry{})
	fmt.Println("Database migrated")
}
