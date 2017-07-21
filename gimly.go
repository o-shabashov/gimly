package main

import (
    "gopkg.in/gographics/imagick.v3/imagick"
    "gimly/models"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "log"
)

func main() {
    api := rest.NewApi()

    // Веб страница со статусом, например http://gimly.com/.status
    statusMw := &rest.StatusMiddleware{}

    api.Use(statusMw)
    api.Use(rest.DefaultDevStack...)

    // Маршрутизация
    router, err := rest.MakeRouter(
        rest.Post("/image", GetImage),

        rest.Get("/.status", func(w rest.ResponseWriter, r *rest.Request) {
            w.WriteJson(statusMw.GetStatus())
        }),
    )

    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler())) // TODO взять данные из os.GetEnv()
}

func GetImage(w rest.ResponseWriter, r *rest.Request) {
    imagick.Initialize()
    defer imagick.Terminate()

    // Чтение и валидация запроса
    request := models.ReadRequest(r)

    // В этот канал будем отправлять искажённые слои
    channel := make(chan *imagick.MagickWand, len(request.Layers))

    // Финальное изображение
    image := imagick.NewMagickWand()

    // Инит финального изображения без цвета
    pw := imagick.NewPixelWand()
    pw.SetColor("none")

    image.NewImage(request.Height, request.Width, pw)
    image.SetImageFormat("JPG") // TODO взять формат из запроса

    // Чтобы на месте перемещённых пикселей была прозрачность
    image.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    // Запустить искажение всех слоёв в отдельных потоках, результат прилетит в канал channel
    for _, layer := range request.Layers {
        go models.DistortLayer(channel, layer)
    }

    // Подписываемся на канал и ждём данных от горутин
    for range request.Layers {
        select {
        case layer := <-channel:
            image.CompositeImage(layer, imagick.COMPOSITE_OP_OVER, false, 0, 0) // TODO top, left
        }
    }

    w.Header().Set("Content-Type", "image/jpeg")
    w.(http.ResponseWriter).Write(image.GetImageBlob())
}
