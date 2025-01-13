package main

import (
	"finance_backend/internal/config"
	"finance_backend/internal/handlers"
	"finance_backend/internal/middleware"
	"finance_backend/internal/repository"
	"finance_backend/internal/services"
	"finance_backend/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Setup dependencies
	deps := setupDependencies(db, cfg)

	// Initialize router
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	setupRoutes(router, deps, cfg.JWTSecret)

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Fatal(router.Run(":" + cfg.ServerPort))
}

type Dependencies struct {
	handlers *Handlers
}

type Handlers struct {
	user        *handlers.UserHandler
	transaction *handlers.TransactionHandler
	bill        *handlers.BillHandler
	dashboard   *handlers.DashboardHandler
}

func setupDependencies(db *gorm.DB, cfg *config.Config) *Dependencies {
	// Repositories
	userRepo := repository.NewUserRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	billRepo := repository.NewBillRepository(db)

	// Services
	userService := services.NewUserService(userRepo)
	transactionService := services.NewTransactionService(transactionRepo)
	billService := services.NewBillService(billRepo)
	dashboardService := services.NewDashboardService(transactionRepo, billRepo)

	// Handlers
	handlers := &Handlers{
		user:        handlers.NewUserHandler(userService, cfg.JWTSecret),
		transaction: handlers.NewTransactionHandler(transactionService),
		bill:        handlers.NewBillHandler(billService),
		dashboard:   handlers.NewDashboardHandler(dashboardService),
	}

	return &Dependencies{
		handlers: handlers,
	}
}

func setupRoutes(r *gin.Engine, deps *Dependencies, jwtSecret string) {
	// Public routes
	api := r.Group("/api")
	{
		api.POST("/register", deps.handlers.user.Register)
		api.POST("/login", deps.handlers.user.Login)

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(jwtSecret))
		{
			// User routes
			user := protected.Group("/user")
			{
				user.GET("/profile", deps.handlers.user.GetProfile)
				user.PUT("/profile", deps.handlers.user.UpdateProfile)
			}

			// Dashboard routes
			protected.GET("/dashboard", deps.handlers.dashboard.GetDashboard)

			// Transaction routes
			transactions := protected.Group("/transactions")
			{
				transactions.POST("", deps.handlers.transaction.Create)
				transactions.GET("", deps.handlers.transaction.List)
				transactions.GET("/:id", deps.handlers.transaction.Get)
				transactions.PUT("/:id", deps.handlers.transaction.Update)
				transactions.DELETE("/:id", deps.handlers.transaction.Delete)
			}

			// Bill routes
			bills := protected.Group("/bills")
			{
				bills.POST("", deps.handlers.bill.Create)
				bills.GET("", deps.handlers.bill.List)
				bills.GET("/:id", deps.handlers.bill.Get)
				bills.PUT("/:id", deps.handlers.bill.Update)
				bills.DELETE("/:id", deps.handlers.bill.Delete)
			}
		}
	}
}
