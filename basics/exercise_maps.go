package main



import (
    "strings"
    "golang.org/x/tour/wc"
)


func WordCount(s string) map[string]int {
    
    var wordcounts = make(map[string]int)


    var words []string = strings.Fields(s)

    for _, word := range words {

        wordcounts[word]++
        // if _, ok := wordcounts[word]; ok == false {
        //     wordcounts[word] = 1
        // } else {
        //     wordcounts[word]++
        // }
    }
    return wordcounts
}


func main() {
    wc.Test(WordCount)
}
