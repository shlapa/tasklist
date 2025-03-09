# Tasklist ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

**Tasklist** — это простое CRUD-приложение для управления задачами, написанное на Go с использованием Fiber и PostgreSQL. Приложение предоставляет REST API для создания, чтения, обновления и удаления задач.

---

## Оглавление

1. [Особенности](#особенности)
2. [Технологии](#технологии)
3. [Установка](#установка)
4. [Настройка](#настройка)
5. [Использование](#использование)
6. [Миграции](#миграции)
7. [API](#api)
8. [Лицензия](#лицензия)

---

## Особенности

- Создание, чтение, обновление и удаление задач.
- Поддержка статусов задач: `new`, `in_progress`, `done`.
- Автоматическое применение миграций при запуске.
- Простая настройка через переменные среды.

---

## Технологии

- **Язык программирования:** Go
- **Фреймворк:** [Fiber](https://github.com/gofiber/fiber)
- **База данных:** PostgreSQL
- **Миграции:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **Контейнеризация:** Docker (опционально)

---

## Установка

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/shlapa/tasklist.git
cd tasklist
```
---

### 2. Установка зависимости

```bash
go mod download
```
---

### 3. Настройте базу данных

- Создайте базу данных tasklist в PostgreSQL:

```bash
CREATE DATABASE tasklist;
```
---
## Настройка

### 1. Переменные среды
- Создайте файл .env в корне проекта и добавьте следующие переменные:
```bash
DD_URL=postgres://user:password@localhost:5432/tasklist?sslmode=disable
```
- Замените user и password от PostgreSQL.
---
### 2. Применение миграций
- Миграции применяются автоматически при запуске приложения. Если нужно применить их вручную, выполните:
```bash
migrate -path ./migrations -database "$DD_URL" up
```
- Откат миграции
```bash
migrate -path ./migrations -database "$DD_URL" down
```
- Замените user и password от PostgreSQL.
---
## Использование
- Запуск приложения
```bash
go run main.go
```
- Приложение будет доступно по адресу: http://localhost:<порт>, где <порт> — порт, указанный в вашем коде (по умолчанию 3000).
## API
- В pkg/postmanCollection предоставленна коллекция для Postman
### Создание задачи
Метод: POST

URL: /tasks

Тело запроса (JSON):
```bash
{
  "title": "Изучить Go",
  "description": "Изучить основы языка Go",
  "status": "new"
}
```
### Получение всех задач
Метод: GET

URL: /tasks

### Получение задачи по ID
Метод: GET

URL: /tasks/:id

### Обновление задачи
Метод: PUT

URL: /tasks/:id

Тело запроса (JSON):
```bash
{
  "title": "Изучить Go (обновлено)",
  "description": "Изучить основы языка Go и Fiber",
  "status": "in_progress"
}
```
### Удаление задачи
Метод: DELETE

URL: /tasks/:id
