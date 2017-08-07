package main

import (
    "gimly/models"
    "net/http"
    "log"
    "gopkg.in/gographics/imagick.v3/imagick"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/joho/godotenv"
    "os"
    "sort"
    "fmt"
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
        rest.Post("/image", GetImage),

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
    mapPositionMw := make(map[int]*imagick.MagickWand)

    // Чтение и валидация запроса
    data := models.PostData{}
    if err := r.DecodeJsonPayload(&data); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // В этот канал будем отправлять искажённые слои
    channel := make(chan models.PositionMagicWand, len(data.Layers))

    // А в этот канал будем отправлять ошибки
    errors := make(chan error)

    // Финальное изображение
    image := imagick.NewMagickWand()

    // Инит финального изображения без цвета
    pw := imagick.NewPixelWand()
    pw.SetColor("none")

    image.NewImage(data.Height, data.Width, pw)
    image.SetImageFormat(data.Format)

    // Чтобы на месте перемещённых пикселей была прозрачность
    image.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    // Запустить искажение всех слоёв в отдельных потоках, без учёта порядка, результат прилетит в канал channel
    for _, layer := range data.Layers {
        go models.DistortLayer(channel, errors, layer)
    }

    // Подписываемся на оба канала, ждём данных от горутин и записываем их в массив в виде позиция - слой
    for range data.Layers {
        select {
        case pmw := <-channel:
            mapPositionMw[pmw.Position] = pmw.MagicWand
        case err := <-errors:
            rest.Error(w, err.Error(), 500) // TODO нормальные коды ошибок
        }
    }

    // Сортируем слои по порядку https://stackoverflow.com/questions/23330781/sort-golang-map-values-by-keys
    var keys []int
    for k := range mapPositionMw {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    // Накладываем слои по порядку на финальное изображение
    for _, k := range keys {
        image.CompositeImage(mapPositionMw[k], imagick.COMPOSITE_OP_OVER, false, 0, 0) // TODO top, left
        image.WriteImage(fmt.Sprintf("%v.png", k))

        // Без этого горутины зависнут и рест не отдаст контент, т.к. *MagickWand передаётся в канал по ссылке внутри
        // другой структуры, и не может сам себя уничтожить.
        mapPositionMw[k].Destroy()
    }

    w.Header().Set("Content-Type", "image/jpeg")
    w.(http.ResponseWriter).Write(image.GetImageBlob())
}
