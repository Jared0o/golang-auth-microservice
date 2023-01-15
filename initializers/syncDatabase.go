package initializers

import "github.com/jared0o/auth-microservice/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
