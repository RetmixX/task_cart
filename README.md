# cart_api

# Переменные окружения
- DB_NAME
- DB_USER
- DB_PASSWORD
- DB_HOST
- DB_PORT
- SERVER_PORT
- SERVER_MODE

# Пример заполнения
- DB_NAME=cart_db
- DB_USER=cart_user
- DB_PASSWORD=password
- DB_HOST=db
- DB_PORT=5432
- SERVER_PORT=0.0.0.0:3000
- SERVER_MODE=debug

# Перед запуском создайте .env

# Запуск

cd .deploy/local

docker compose up -d --build

# Swagger

# host:port/swagger/index.html

