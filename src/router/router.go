package router

import (
	"app/src/config"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(app *fiber.App, db *gorm.DB) {

	UserService := service.NewUserService(db)
	CommentService := service.NewCommentService(db)
	ComService := service.NewComService(db)
	AdminService := service.NewAdminService(db)

	v1 := app.Group("/v1")
	AuthRoutes(v1)

	//HealthCheckRoutes(v1, healthCheckService)

	UserRoutes(v1, UserService)
	CommentRoutes(v1, CommentService)
	ComRoutes(v1, ComService)
	AdminRoutes(v1, AdminService)

	// TODO: add another routes here...

	if !config.IsProd {
		DocsRoutes(v1)
	}
}
