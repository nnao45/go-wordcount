# go-wordcount
Counting word in a string.

```bash
package main

import "github.com/nnao45/go-wordcount"
import "fmt"

var kensaku = "たけ"
var word = `
このたけがきに　たけたてかけたのは
　　　たけたてかけたかったから　たけたてかけた
`

func main() {
        fmt.Printf("%d回\n", wordcount.WC(kensaku, word)) //=>4回
}
```

