# thumbnail-proxy-client

## Описание
thumnail-proxy-client является CLI утилитой, реализующей GRPC-клиент для сервиса thumbnail-proxy: https://github.com/sletkov/thumbnail-proxy

## Команды и флаги
### thumbnailproxyclient get [urls...]
Выполняет скачивание превью YouTube видео для указанных сссылок.
#### Флаг --async
Если указать флаг --async, превью будут скачаны ассинхронно.

## Установка и запуск
### Для linux и macOS
1. Склонируйте репозиторий на локальную машину командой `git clone [github.com/repository/path]`
2. Перейдите в директорию сервиса командой `cd [path/to/service]`
3. Выполните команду `sudo make install`

### Для windows
1. Склонируйте репозиторий на локальную машину командой `git clone [github.com/repository/path]`
2. Перейдите в директорию сервиса командой `cd [path/to/service]`
3. Выполните команду `go build -o thumbnail-proxy-client.exe`
4. Добавьте путь к директории, содержащей файл "thumbnail-proxy-client.exe" в системную переменную Path
5. Поместите файл thumbnail-proxy-client в директорию $HOME