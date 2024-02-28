# Установка

Для пользования прокси надо сгенерировать .env-файл, а также сгенерировать и установить корневой сертификат.

## Linux

### Скрипт

`sudo ./install.sh postgres postgres postgres ProxyStorage proxy-db`

Докер-контейнер с БД, прокси и веб-апи также собирается и поднимается скриптом

### Вручную

1. Создать файл `.env` в папке config
`touch config/.env`
2. Добавить туда следующий текст
```
MIGRATOR_PASSWORD="postgres"
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="postgres"
POSTGRES_DB="ProxyStorage"
POSTGRES_HOST="proxy-db"
```
3. Запустить скрипт `gen_ca.sh` из корневой папки
`./gen_ca.sh`
4. Скопировать полученный файл корневого сертификата в локальное хранилище сертификатов
`sudo cp proxy-serv-ca.crt /usr/local/share/ca-certificates/proxy-serv-ca.crt`
5. Обновить сертификаты
`sudo update-ca-certificates`
6. Поднять докер
`docker-compose up --build --detach`

## Windows

### Скрипт

`./install.ps1 postgres postgres postgres ProxyStorage proxy-db`

Докер-контейнер собирается отдельной командой

`docker-compose up --build --detach`

### Вручную

1-3. Повторить шаги 1-3 из инструкции для Linux
4. Установить сгенерированный сертификат
`Import-Certificate -FilePath 'proxy-serv-ca.crt' -CertStoreLocation Cert:\CurrentUser\Root`
4.1. Либо вручную установить его из проводника двойным нажатием
5. Поднять докер
`docker-compose up --build --detach`

# Использование

## cURL

`curl -x localhost:8080 http://example.org`
`curl -x localhost:8080 --ssl-revoke-best-effort https://example.org`

## Firefox

1. "Настройки" -> "Прокси" -> "Настроить..." -> localhost:8080 -> Ок
2. "Настройки" -> "Сертификаты" -> "Просмотр сертификатов..." -> "Импортировать..." -> Выбрать сгенерированный сертификат -> Отметить обе галочки "Доверять при..." -> Ок