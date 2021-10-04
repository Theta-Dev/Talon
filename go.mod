module github.com/Theta-Dev/Talon

go 1.16

require (
	github.com/stretchr/testify v1.7.0
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.1
	gorm.io/driver/sqlite v1.1.5
	gorm.io/gorm v1.21.15
)

// Patch to allow migrating constraints
replace gorm.io/driver/sqlite => github.com/Theta-Dev/gorm-sqlite v1.1.5-0.20211003192412-c8ad3266aa58
