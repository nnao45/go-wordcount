package main

import "github.com/nnao45/go-wordcount"
import "fmt"

var runes = []rune("たけ")
var word = `
このたけがきに　たけたてかけたのは
　　　たけたてかけたかったから　たけたてかけた
`

func main() {
        fmt.Printf("%d回\n", wordcount.WC(runes, word))
}
