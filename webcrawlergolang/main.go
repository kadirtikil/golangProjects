package main 

import (
    "fmt"
    "os"
    "webcrawler/webcrawler"
)

func main() {
    fmt.Println("Main started")

    if len(os.Args) < 2 {
        fmt.Println("no website providet")
        os.Exit(1)
    }    
    
    if len(os.Args) > 2 {
        fmt.Println("too many arguments. only provide 1 website to crawl")
        os.Exit(1)
    }
    

    fmt.Printf("starting crawl of: %s\n", os.Args[1])
   

    // AAAAh yes, note to myself: handle your error dumbo. couldve saved hours of wondering why this function is not working if
    // i just handled the error....
    temp, err := webcrawler.FetchHTML(os.Args[1])
    
    if err != nil {
        fmt.Println("Error at line 31, trying FetchHTML")
    }

    fmt.Println(temp)
}
