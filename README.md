# Golang Image Manipulator Library, aka GIMLy

Сделан потому что PHP Imagick::distort сосёт в распаралеливании. Потенциально Гимли это замена генератору, потому что быстро и все хотят ВЖУХ! и картинка появилась на главной.

# Зависимости

* ImageMagick 6-ой версии. Потенциально можно и на 7, но нужно тогда подключить `gopkg.in/gographics/imagick.v3/imagick` а не `v2` и немного изменить код.
* [Glide](https://glide.sh/)

# Установка

```shell
curl https://glide.sh/get | sh
glide install
```

# Запуск

```shell
go run gimly.go
```

# Важно

* При изменении JSON схемы запроса, лучше всего воспользоваться [GoJson](http://github.com/ChimeraCoder/gojson/gojson)  

# Тестирование

```shell
go get github.com/xeipuuv/gojsonschema
go get github.com/smartystreets/goconvey
$GOPATH/bin/goconvey
```

Открыть в браузере [http://localhost:8080](http://localhost:8080)

Тесты написаны с помощью [GoConvey](http://goconvey.co)

# TODO

* [x] Нормальная обработка ошибок - с помощью rest.Error()
* [x] Валидация запроса по JSON схеме
* [x] Кроп и поворот дизайна
* [ ] Докер контейнер
* [ ] Тесты
* [x] Разобраться с Glide, почему-то не устанавливает зависимости, пришлось юзать go get
* [ ] Postman коллекция, лучше после тестов
* [ ] Объявить об изменении JSON схемы: матрица искажений должна быть массивом, а не строкой, distortion_order стал интом
* [ ] Нормальные коды ошибок, а не просто 500 на всё
* [x] Смещение слоёв относительно финального изображение
