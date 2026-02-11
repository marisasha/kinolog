# KINOLOG 

REST API сервис позволяющий пользователям искать информацию о фильмах/сериалах , ставить оценки и писать рецензии.
Для поиска используется AI Groq

# Tech Stack

- Go (Golang)
- Gin Web Framework
- PostgreSQL

# Run with Docker
1. Создать файл `.env` с настройками для PostgreSQL

```.env
DB_USERNAME={ username }
DB_PASSWORD={ password }
DB_HOST={ host } //пользуйтесь db для Docker
DB_PORT={ port }
DB_NAME={ name }
DB_SSLMODE={ sslmode }
```

2. Запустить проект
```
docker compose up --build
```

3. Использовать api
```
http://localhost:8000/swagger/index.html
```

# Database Migrations

Миграции выполняются автоматически при запуске через контейнер `migrate`.
SQL-файлы находятся в папке: schema

# Project Structure

- **cmd** — точка входа (main.go)  
- **pkg** — бизнес-логика (handler, service, repository)  
- **schema** — SQL миграции  
- **docs** — Swagger документация  
- **configs** — конфигурационные файлы  
- **Dockerfile**  
- **docker-compose.yml**