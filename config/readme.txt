Установка Сервера Базы данных
sudo apt-get install postgresql

Установка компилятора Go
sudo apt-get install golang-src

В переменные окружения нужно будет вписать GOPATH и прочее из гайда.
Кажется всего 3-4 пути через EXPORT добавить в конфигурационку твоего терминала.
https://www.tecmint.com/install-go-in-linux/    <-- начинать с главы configure go environment


Склонировать/скопировать текущий проект в папку $GOPATH/src/MatchaServer/
Это важный строгий пункт. Все доп пакеты относительно $GOPATH/src указываются...

Создать базу данных и пользователя скриптом setup.sh, который использует setup.sql
sudo ./setup.sh

выполнить команды
sudo -i -u postgres
psql -d matcha_db    // Последние две команды на маке выполняются как psql postgres
SELECT * FROM pg_database;
SELECT * FROM pg_shadow;
SELECT * FROM information_schema.tables  WHERE table_schema='public' ORDER BY table_name;
SELECT table_name, column_name FROM information_schema.columns WHERE table_schema='public' ORDER BY table_name;
SELECT * FROM test WHERE interest[1] IN (1, 2)
SELECT * FROM test WHERE interest[1123]=2 OR interest[2]=2 OR interest[3]=2 OR interest[4]=2 OR interest[5]=2
SELECT * FROM test WHERE interest='{2,3,5}'
INSERT INTO test (name, interest) VALUES ('vasya', '{2, 3, 5}')

в списках должны находиться база данных matcha_db и пользователь bsabre
запустить программу setup.go

go run setup.go

Все пункты должны быть зелеными
Запускать тесты не обязательно - если я запушил - значит они проходят хорошо (от машины они не зависят).
Тесты - это файлы *_test.go. Лежат в подпапках

go test -v -cover

Все тесты должны быть зелеными

go test -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html

Теперь в cover.html содержится визуализация покрытия кода тестами