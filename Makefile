# Переменные
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOMOD=$(GOCMD) mod

# Имя итогового бинарного файла
BINARY_NAME=shortener
BINARY_PATH=./bin/$(BINARY_NAME)

# Основной файл приложения
MAIN_FILE=cmd/server/main.go

# Установка зависимостей
.PHONY: deps
deps:
	$(GOMOD) tidy

# Сборка приложения
.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_PATH) $(MAIN_FILE)

# Запуск приложения
.PHONY: run
run:
	$(GOCMD) run $(MAIN_FILE)

# Запуск тестов
.PHONY: test
test:
	$(GOTEST) ./...

# Форматирование кода
.PHONY: fmt
fmt:
	$(GOFMT) ./...

# Очистка скомпилированных файлов
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)

# Установка зависимостей, форматирование и сборка
.PHONY: all
all: deps fmt build
