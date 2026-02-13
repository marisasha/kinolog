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
DB_USERNAME= postgres
DB_PASSWORD= postgres // можно любой другой
DB_HOST= db 
DB_PORT= 5432
DB_NAME= kinolog // можно любое другое
DB_SSLMODE= disable
```
2. Создать в корне папку `config` , а внутри папки файл `config.yaml`

```config.yaml
port: "8000"

db:
  host: "db"
  port: "5432"
  username: "postgres"
  password: "2425"
  dbname: "kinolog_db"
  sslmode: "disable"

ai:
  url: { url Groq api }
  key: { key Groq api }
  model: { model Groq api}

```

3. Запустить проект
```
docker compose up --build
```

4. Использовать api
```
http://localhost:8000/swagger/index.html
```

# Database Migrations

Миграции выполняются автоматически при запуске через контейнер `migrate`.
SQL-файлы находятся в папке: migrations

# Project Structure

- **cmd** — точка входа (main.go)  
- **pkg** — бизнес-логика (handler, service, repository)  
- **migrations** — SQL миграции  
- **docs** — Swagger документация  
- **configs** — конфигурационные файлы  
- **Dockerfile**  
- **docker-compose.yml**