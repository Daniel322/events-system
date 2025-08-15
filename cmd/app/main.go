package main

import (
	"context"
	"events-system/infrastructure/providers/cron"
	db "events-system/infrastructure/providers/db"
	"events-system/infrastructure/providers/http/controllers"
	"events-system/infrastructure/providers/http/server"
	"events-system/infrastructure/providers/telegram"
	"events-system/internal/domain"
	"events-system/internal/repositories"
	"events-system/internal/services"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// init external providers
	db := db.NewDatabase(os.Getenv("GOOSE_DBSTRING"))
	server := server.NewEchoInstance()

	// init domain factories
	userFactory := domain.NewUserFactory()
	accountFactory := domain.NewAccountFactory()
	eventFactory := domain.NewEventFactory()
	taskFactory := domain.NewTaskFactory()

	// init repositories
	userRepository := repositories.NewRepository("UserRepository", db, userFactory)
	accountRepository := repositories.NewRepository("AccountRepository", db, accountFactory)
	eventRepository := repositories.NewRepository("EventRepository", db, eventFactory)
	taskRepository := repositories.NewRepository("TaskRepository", db, taskFactory)

	// init services
	userService := services.NewUserService(db, userRepository, accountRepository)
	accountService := services.NewAccountService(db, accountRepository)
	eventsService := services.NewEventService(db, eventRepository, taskRepository)
	tasksService := services.NewTaskService(db, taskRepository, eventRepository, accountRepository)

	// init controllers
	userController := controllers.NewUserController(
		server.Instance,
		userService,
	)

	tgBotProvider, err := telegram.NewTgBotProvider(os.Getenv("TG_BOT_TOKEN"), userService, accountService, eventsService)

	if err != nil {
		panic(err.Error())
	}

	cronProvider := cron.NewCronProvider(tgBotProvider, tasksService)

	cronProvider.Bootstrap()

	go tgBotProvider.Bootstrap()

	// init http routes
	userController.InitRoutes()

	// start http server
	go server.Start(os.Getenv("HTTP_PORT"))

	<-ctx.Done()

	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db.Close()
	tgBotProvider.Close()

	if err := server.Instance.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %w", err)
	}

}
