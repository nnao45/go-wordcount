package wordcount

func WC(runes []rune, word string) {
        var k int
        var i int
        for _, r := range text {
                if r == runeSlice[k] {
                        if k < len(runeSlice)-1 {
                                k++
                        } else if k == len(runeSlice)-1 {
                                i++
                        }
                } else {
                        k = 0
                }
        }
}
