Чтобы запустить проект надо сначала: 

В корневой директории проекта запустить:
- swag init (Для того чтобы был swagger)
- go mod download (Качаем зависимости)

В папке проекта: 

- bash backend/scripts/generate_db.sh (Для генерации SQLC можно в git терминале) 
- docker compose up -d (Запустит докер контейнеры с PG базой данных + PG admin)

Для подключения к PG admin: 
http://localhost:5050/browser/ в браузере 

Пароль 1qaz2WSX 
    ПКМ по servers register server: 
    в General написать любое name 
    В Connection 
        Host name/address: service-db 
        username: ifigurin12
        password: 1qaz2WSX

Если вы в vscode то там есть 2 launch config'a 
1. Запустит сам backend 
2. Запустит тесты 

.env файл будет лежать на google диск, ссылка на него будет на hh.ru (мог бы тут оставить, но там пароль к smtp email) его положить надо в корень проекта 


