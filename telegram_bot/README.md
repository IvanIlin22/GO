# Telegram bot
Рассмотрим создание Telegram бота на Go, который будет выдавать новости с "Habr" и "Tproger". Мы не будем писать обработку совсем всех
событий с нуля, а воспользуемся готовой библиотекой. Есть готовая библиотека Telegram Bot API. Для
начала нужно ее установить. Устанавливается она через команду"

`go get gopkg.in/telegram-bot-api.v4`

Зарегистрируем нашего бота в специализированном боте BotFather и получим токен. Последняя подготовительная часть — нам необходимо
каким-то образом выложить бота в Интернет либо еще как-то сделать, чтобы Telegram мог до него достучаться. В этом
примере мы воспользуемся сервисом ngrok.


