# Golang Image Manipulator Library, aka GIMLy

Сделан потому что PHP Imagick::distort сосёт в распаралеливании. Потенциально Гимли это замена генератору, потому что быстро и все хотят ВЖУХ! и картинка появилась на главной.

## Алгоритм работы

[Sketchboard](https://sketchboard.me/nACNMWo6XpyJ#/), нажать `ALT` + `P`

# Зависимости

* ImageMagick 6-ой версии. Потенциально можно и на 7, но нужно тогда подключить `gopkg.in/gographics/imagick.v3/imagick` а не `v2` и немного изменить код.
* [Glide](https://glide.sh)

# Установка

* Imagick, Ubuntu

```shell
sudo apt-get install libmagickwand-dev
```

* Imagick, Mac

```shell
brew install imagemagick@6 --from-source
```

* Glide

```shell
curl https://glide.sh/get | sh
glide install
```

# Запуск

Отредактируйте `.env` файл по своему усмотрению.

```shell
go run gimly.go
```

Гимли будет доступен по адресу [http://localhost:8901](http://localhost:8901)

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4791073523e21c6d5364)

# Важно

* При изменении JSON схемы запроса, лучше всего воспользоваться [GoJson](http://github.com/ChimeraCoder/gojson)  

# Тестирование

```shell
go get github.com/xeipuuv/gojsonschema
go get github.com/smartystreets/goconvey
$GOPATH/bin/goconvey
```

Открыть в браузере [http://localhost:8080](http://localhost:8080)

Тесты написаны с помощью [GoConvey](http://goconvey.co)

# Документация используемых библиотек

* [Менеджер зависимостей Glide](https://glide.sh)
* [Биндинг ImageMagick в Go](https://github.com/gographics/imagick)
* [RESTful JSON API сервис](https://github.com/ant0ine/go-json-rest)
* [Чтение .env файлов](github.com/joho/godotenv)
* [BDD фреймворк тестирования GoConvey](http://goconvey.co)

### Опционально

* [Live reload utility for Go web servers](https://github.com/codegangsta/gin)

# TODO

* [x] Нормальная обработка ошибок - с помощью rest.Error()
* [x] Валидация запроса по JSON схеме
* [x] Кроп и поворот дизайна
* [ ] Докер контейнер
* [x] Тесты
* [x] Разобраться с Glide, почему-то не устанавливает зависимости, пришлось юзать go get
* [x] Postman коллекция, лучше после тестов
* [ ] Объявить об изменении JSON схемы: матрица искажений должна быть массивом, а не строкой, distortion_order стал интом
* [ ] Нормальные коды ошибок, а не просто 500 на всё
* [x] Смещение слоёв относительно финального изображение
* [x] Возвращать error при вызове image.Composite()
* [ ] Реализовать `PartialDistortMatrix`. Не стал этого делать сразу, потому что в оригинальном generator-laravel-package этот тип искажения сделан не понятно и вообще там очень сильное колдунство. Может получится полностью алгоритм переписать.
