# запуск docker
sudo docker run --name=notes-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
# создание папки миграций
migrate create -ext sql -dir ./schema -seq init
# применение миграции инициализации
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
# применение миграции удаления
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
# завершение работы в docker базы данных
sudo docker stop notes-db
