package wordcount

func WC(runes []rune, word string) (i int){
        var k int
        for _, r := range word {
                if r == runes[k] {
                        if k < len(runes)-1 {
                                k++
                        } else if k == len(runes)-1 {
                                i++
                        }
                } else {
                        k = 0
                }
        }
        return
}
