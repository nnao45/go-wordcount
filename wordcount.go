package wordcount

func WC(word string, runes string) (i int){
        rs := []rune(runes)
        var k int
        for _, r := range word {
                if r == rs[k] {
                        if k < len(rs)-1 {
                                k++
                        } else if k == len(rs)-1 {
                                i++
                                k = 0
                        }
                } else {
                        k = 0
                }
        }
        return
}
