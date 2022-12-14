<h2>Тестовое задание для стажера-разработчика</h2>

<h3/>Задача</h3>

Реализовать сервис, предоставляющий API по созданию сокращённых ссылок

Ссылка должна быть:
- Уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка
- Длиной 10 символов
- Из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)

<h3>Условие</h3>

Сервис должен быть написан на Go и принимать следующие запросы по http:

- Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый
- Метод Get, который будет принимать сокращённый URL и возвращать оригинальный

<h3>Запуск</h3>
<b>Запуск Миграции базы данных:</b>

`make migration`

<b>Запуск API с in-memory-хранилищем:</b>

`make inMemory` или `docker compose -f inMemory.yml  up -d`

<b>Запуск API с PostgreSQL:</b>

`make postgres` или `docker compose -f postgres.yml  up -d`
