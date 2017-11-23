# go-wordcount
Counting word in a string.

```bash
package main

import "github.com/nnao45/go-wordcount"
import "fmt"

var word = `
このたけがきに　たけたてかけたのは
　　　たけたてかけたかったから　たけたてかけた
`
var kensaku = "たけ"

func main() {
        fmt.Printf("%d回\n", wordcount.WC(word, kensaku)) //=>4回
}
```

