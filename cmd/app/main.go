package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
	"tasklist/internal/database"
	"tasklist/internal/handler"
	"tasklist/internal/repository"
	"tasklist/routes"
)

// Инициализация подключения через переменные окружения
func mustURL() string {
	databaseURL := os.Getenv("DD_URL")
	if databaseURL == "" {
		log.Fatal("$DD_URL environment variable not set")
	}
	return databaseURL
}

func main() {
	//databaseURL := "postgres://???:???@localhost:5432/tasklist"
	databaseURL := mustURL()
	// Подключение к бд
	pool, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Инициализация репозитория и обработчика
	taskRepo := repository.NewTaskRepository(pool)
	taskHandler := handler.NewTaskHandler(taskRepo)

	// Создаём новое приложение Fiber
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Подключаем middleware
	app.Use(logger.New())   // Логирование запросов
	app.Use(compress.New()) // Сжатие ответов
	app.Use(recover.New())  // Восстановление после паники
	app.Use(limiter.New())  // Лимит запросов для предотвращения DDOS атак

	// Регистрация маршрутов
	routes.TaskListRoutes(app, taskHandler)

	// Запускаем сервер
	log.Fatal(app.Listen(":3000"))
}
