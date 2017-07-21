# Golang Image Manipulator Library, aka GIMLy

Сделан потому что PHP Imagick::distort сосёт в распаралеливании. Потенциально Гимли это замена генератору, потому что быстро и все хотят ВЖУХ! и картинка появилась на главной.

# Зависимости

Imagick 7-ой версии.
Потенциально можно и на 6, но нужно тогда подключить `gopkg.in/gographics/imagick.v2/imagick` а не `v3`. 

# Установка
Glide глючит, так что пока вот так

```shell
go get gopkg.in/gographics/imagick.v3/imagick github.com/xeipuuv/gojsonschema github.com/ant0ine/go-json-rest/rest
glide up
```

# Запуск

```shell
go run gimly.go
```

# TODO

* Валидация запроса по JSON схеме
* Кроп и поворот дизайна
* Нормальная обработка ошибок - с помощью rest.Error()
* Докер контейнер
* Тесты
* Разобраться с Glide, почему-то не устанавливает зависимости, пришлось юзать go get
* Postman коллекция, лучше после тестов