# Simple-Telegram-Golang-bot
## Описание / Description
Первый опыт работы с Golang.
Простой Телеграм бот с выдачей ответов на действия пользователей через кнопки.

First attempt on Golang
Simple Telegram bot that provides responses to user buttons via buttions.



## Настройка / Setting up

Для корректной работы, в Dockerfile нужно прописать токен Телеграм бота.
Токен получается при создании бота в @BotFather.
Текст ответов изменяется в файлах, находящиеся в папке messages.

Перед запуском, создаете Docker образ через ``` docker build . --tag <Название образа>:<версия> ```
Дальнейший запуск производится через создание контейнера, например ``` docker run -it --name <Название контейнера> <Название образа>:<версия> ```

You need to insert token of your Telegram bot in Dockerfile.
You can get token after creation your bot via @BotFather.
Messges text can be changed in files that places in messages folder.

Before starting, you need to create Docker image via ``` docker build . --tag <Image name>:<version> ```
Further start up proceeds via ``` docker run -it --name <Container name> <Image name>:<version> ```
