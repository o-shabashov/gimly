package main

import (
    "net/http"
    "log"
    "gopkg.in/gographics/imagick.v2/imagick"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/joho/godotenv"
    "os"
    "sort"
    "gimly/models"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    api := rest.NewApi()

    // Веб страница со статусом, например http://gimly.dev/.status
    statusMw := &rest.StatusMiddleware{}

    api.Use(statusMw)
    api.Use(rest.DefaultDevStack...)

    // Маршрутизация
    router, err := rest.MakeRouter(
        rest.Post("/generate-image", GetImage),

        rest.Get("/.status", func(w rest.ResponseWriter, r *rest.Request) {
            w.WriteJson(statusMw.GetStatus())
        }),
    )

    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(os.Getenv("GIMLY_ADDRESS"), api.MakeHandler()))
}

func GetImage(w rest.ResponseWriter, r *rest.Request) {
    imagick.Initialize()
    defer imagick.Terminate()

    // Сюда запишем данные от горутин в виде позиция - слой
    mapPositionMw := make(map[int]models.PositionMagicWand)

    // Чтение и валидация запроса
    postData := models.PostData{}
    if err := postData.Validate(r); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Перевод процентного позиционирования в пиксельное TODO ошибки
    postData.ConvertPositioning()

    // В этот канал будем отправлять искажённые слои
    channel := make(chan models.PositionMagicWand, len(postData.Layers))

    // А в этот канал будем отправлять ошибки
    errors := make(chan error)

    // Финальное изображение
    image := imagick.NewMagickWand()

    // Инит финального изображения без цвета
    pw := imagick.NewPixelWand()
    pw.SetColor("none")

    image.NewImage(uint(postData.Height), uint(postData.Width), pw)
    image.SetImageFormat(postData.Format)

    // Чтобы на месте перемещённых пикселей была прозрачность
    image.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    // Запустить обработку всех слоёв в отдельных потоках, без учёта порядка, результат прилетит в канал channel
    for _, layer := range postData.Layers {
        go layer.Build(channel, errors)
    }

    // Подписываемся на оба канала, ждём данных от горутин и записываем их в массив в виде позиция - слой
    for range postData.Layers {
        select {
        case pmw := <-channel:
            mapPositionMw[pmw.Layer.Position] = pmw
        case err := <-errors:
            rest.Error(w, err.Error(), 500) // TODO нормальные коды ошибок
        }
    }

    // Сортируем слои по порядку https://stackoverflow.com/questions/23330781/sort-golang-map-values-by-keys
    var keys []int
    for i := range mapPositionMw {
        keys = append(keys, i)
    }
    sort.Ints(keys)

    // Накладываем слои по порядку на финальное изображение
    for _, v := range keys {
        err := image.CompositeImage(
            mapPositionMw[v].MagicWand,
            imagick.COMPOSITE_OP_OVER,
            int(mapPositionMw[v].Layer.Left),
            int(mapPositionMw[v].Layer.Top),
        )

        // Без этого горутины зависнут и рест не отдаст контент, т.к. *MagickWand передаётся в канал по ссылке внутри
        // другой структуры, и не может сам себя уничтожить.
        mapPositionMw[v].MagicWand.Destroy()

        if err != nil {
            rest.Error(w, err.Error(), 500)
        }
    }

    w.Header().Set("Content-Type", "image/"+postData.Format)
    w.(http.ResponseWriter).Write(image.GetImageBlob())

    // Освобождаем память, нам ведь не нужны утечки?
    image.Destroy()
}
