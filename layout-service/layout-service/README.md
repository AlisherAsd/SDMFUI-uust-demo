# Запуск сервиса
1. Установите зависимости
```bash
go get "github.com/gin-contrib/cors"
```
```bash
go get "github.com/gin-gonic/gin"
```
```bash
go get github.com/jackc/pgx/v5
```
```bash
go get github.com/joho/godotenv  
``` 
2. Запустите postgres в docker
```bash
make env-up
```
3. Примените миграции
```bash
make migrate-up
```
