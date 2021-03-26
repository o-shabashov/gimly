# Golang Image Manipulator Library, aka GIMLy

Многопоточное генерирование превью картинки с помощью Imagick::distort на основе матрицы искажений.

# Зависимости

* ImageMagick 6-ой версии. Потенциально можно и на 7, но нужно тогда подключить `gopkg.in/gographics/imagick.v3/imagick` а не `v2` и немного изменить код.

# Установка

* Imagick, Ubuntu

```shell
sudo apt-get install libmagickwand-dev
```

* Imagick, Mac

```shell
brew install imagemagick@6 --from-source
```
# Запуск

Отредактируйте `.env` файл по своему усмотрению.

```shell
go run gimly.go
```

Гимли будет доступен по адресу [http://localhost:8901](http://localhost:8901)

[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.getpostman.com/collections/2ff63a9ff0f4c9050078)

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
* [x] Разобраться с Glide, почему-то не устанавливает зависимости, пришлось использовать go get
* [x] Postman коллекция, лучше после тестов
* [x] Объявить об изменении JSON схемы: матрица искажений должна быть массивом, а не строкой, distortion_order стал числом
* [ ] Нормальные коды ошибок, а не просто 500 на всё
* [x] Смещение слоёв относительно финального изображение
* [x] Возвращать error при вызове image.Composite()
* [x] Реализовать `PartialDistortMatrix`
