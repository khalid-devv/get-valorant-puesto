package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.6" )

func Ygea17OiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGcD7fkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NW1ZRpBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjTv05ugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFpoHj3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ljn4Jq5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYSiHNgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Pza6sX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qok0Ni6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g73yMTvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhHfRMZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vxr32SPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdmTaChaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7sdzKfEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUnrE4RIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVe2WkNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VG4Mr4h9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFdWVokKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKH4oePnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efTfnwfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAeN6gloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjF31EPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcmEPCVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBcIlaP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phSLk3p1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhHP0IWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVNmx19UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DZ5X6KjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1x6HrBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ft8q4yPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVJnroEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjJrUwzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQrdovEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kt6dnWj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqbKa1LtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cPYndEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6idXUoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ohN51YNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHDS4oJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbSNuMNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQesZ0AnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLmAnXDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFhqWkDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGuGnEkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxFEmf5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dUtIPmfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTSMYvXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCOQzlTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9u04BoguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxuJQHFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xpI7rGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkEe0RsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpSTcfGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwQfT5P5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3g62WuLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnQqGZTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IO4DQwtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FI9DcDKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnEXecynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cv0xVi8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FW4XnYkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xk1A1EhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1wiebzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khK9Jm2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAZk5qC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u30sQDO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfj9vhQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABzFb5roWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqqKVSdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfyaWgopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYaDdya0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjX6oVklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func os4ucknhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9AmvW0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lK0EdYq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoYGfl7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxtcSEn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYSW2YVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhxdRNfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3vyVqbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLlEJICTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpW8o0AbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKTnhzz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVL1NvsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWVhsBuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wagm0yefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZDaFWvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elwCJrJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XP3sGzC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pinvxM8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsETDv4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoWVwI57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9LR1DslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjHEKrGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIJRmZirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6n9AwSUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGKKDVBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAvFf3PFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C11qVPZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVif5dtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6GUEwNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65yZi5trWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94ivIridWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2ELvgwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrnHkwUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRYI7jwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyiE0LQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqPxDwp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWmbsYgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4x97upgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a514JJUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovq6C3NMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEEz6J6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGkicsCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpIA9MgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCl0xEdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTezqxr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xx4PCxioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrdTIjZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjrztY7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNfKXJszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMSKisVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rp7npClxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7EklatrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYsisk6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERq6QEecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQEYJUB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbZ1ekjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoQu77jyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMxJoNRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUSAsUX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EW5Vet7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHCEjuI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3X6BNQJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j569vKO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFbg0x3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDCCkXH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyHqYvckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scqBBh4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xaw4xh7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 175Su58VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjyVVwltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbSE8G11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hVleqVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWEFBPP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQ4WON35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eg4Xsm1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMxMDfwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOxNHAByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NntDcKj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwteRXwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yqRPGGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPs5V04QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDCmwVTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDAQic7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3BjzIuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKPSEFKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TibmUUqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QKxCfDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Upw3QrMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQvU0kN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxg5mLyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5CUeBqe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwOWiZZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMEHVS5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyQPK1uxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrJ0T0nFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlYyCS6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lE274OkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7g0PTloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txlMpZ88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFIGhzoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9xb32esWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00xoKjcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7hbHA6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func roQOq3ODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCJjFn7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvOYPhquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rq4g01zaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZaMKVYJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuDPsy2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dd2agWkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyeZLDBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpzZxgvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0YEiyUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kHbByY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sv41bw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEvEJrS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRLWO0OXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLEFuX6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8lR6AMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7pfReOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izQ1GCDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cU61tWaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozaSUAGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIIwpPF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpboCfydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IiGquRKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nsefjBVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HB4xhxHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDaZXPzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXLAQi7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnHFPa1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fn9rsxPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wm7wkYnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dy9kOSPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78HFLfk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGdBbJ5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmEDwObsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRaVIjP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4bIbnoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxr3lC9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rI1K2M43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8nW7cR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzuXCtS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CCyyx9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXLY92ICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8ggR5nsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJOVc1MuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymCJJl6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFi5gybRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LP9SWlVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWoKYLsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWyRsidTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGm4LeY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NA19LBb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQoS0rMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7WxGlzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjWSrJSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKcKsVSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sto0Eho8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4x22U3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUEJsRMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gI4Sno2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIK8mYHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHt3N4T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSAJIstTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qzl7jYbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVPNcbGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWlaUJbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5n6j9cyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxmRH4NJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBe5L3ndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84OidearWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvXHugjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bax99SqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e982asYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F25ZhrsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAiCZVLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQEDKRhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7O8vJF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0z1vG8IhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLY6s1ESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWzJHz0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QYuqH4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yo1zRRuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wL1i4yJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NG4B4AInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WN1hhHOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZWViU5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OHuG73yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mwva3Yu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzpLFqf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4XZkupeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8q6XwIRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adUH5eFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGitUifjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4wdWcGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zj4cg2vIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9KoCwubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvW6GjuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65oHslmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUEe65W5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Awemfd14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVs3xpzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UoJlcGzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5hZ1Jy3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PlNb9tQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyopvZepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yM2bZA0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ipA9mtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3GceQKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lR5icRjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgvcShjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cy833QwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func La8jti8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1BAIgiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iv5PPYRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PWLJYvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJYqp4WCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCUSYoS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdB9cFtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWXiojeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTyxuN1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEGbguZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEOTQnMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3miMcqYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYPZb3nrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjmRb3GdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvZKMCWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mU42tuc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvlzaXquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzZZvr3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFI9HU0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRHkQicZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYyYkl8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQJFzviwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQ1vaw8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IGpekO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFxGGeXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPubJgepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSe3PwiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvhz5uArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VaLXsDr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3NlX8uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOlRaxQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEypZE0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vDX674hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DctZ3aDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEUPENjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Py0EK7YQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tY1WGCKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aaWPS9Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAwWL4hJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBmIn8VxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c86hfIY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8LnSgFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4Fg1vzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Itb6nGKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPfXwKSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlHrwSYMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMreP97WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJM9PXQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYtjDqqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rq89XEqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0emTVfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

