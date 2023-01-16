package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/webtoon/internal/app/controller"
	"github.com/webtoon/internal/app/model"
	"github.com/webtoon/internal/app/repository"
	"github.com/webtoon/internal/app/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	router           *gin.Engine
	manhwaController *controller.ManhwaController
}

//TODO: separate function setup databasee

func NewServer() *Server {

	//initialize database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conncting to database")
	}
	db.AutoMigrate(&model.Manhwa{})

	//new router
	router := gin.Default()
	//TODO: add middleware like JWT auth

	//initialize manhwa controller
	manhwaRepository := repository.NewManhwaRepository(db)
	manhwaService := service.NewManhwaService(manhwaRepository)
	manhwaController := controller.NewController(manhwaService)

	return &Server{
		router:           router,
		manhwaController: manhwaController,
	}
}

func getDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)

}

// Start starts the server and listens for incoming requests.
// It applies the controller routes and sets the listening port.
// If a PORT environment variable is set, it uses that value.
// Otherwise, it uses the default value "8000".
// It handles graceful shutdown when an interrupt signal is received.
func (s *Server) Start() {
	s.manhwaController.SetupRoutes(s.router)

	port := getDotEnvVariable("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: s.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
