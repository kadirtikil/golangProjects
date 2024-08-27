package webcrawler 

import (
    "fmt"
    "net/url"  
    "net/http"
    "strings"
    "io"
    "golang.org/x/net/html"
)


// Never use url as String, couldnt use url. methods because string was url before
func NormalizeURL(urlString string) (string, error) {
    normURL, err := url.Parse(urlString)
    
    if err != nil {
        return "", err
    }
        
    normString := normURL.Host + normURL.Path

    return normString, nil    

}


func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
    // init return value
    var result []string 

    // read the html body
    readBody := strings.NewReader(htmlBody)
    
    // find the url inside the body
    parsedBody, err := html.Parse(readBody)
    // check err
    if err != nil {
        return nil, err 
    }
    
    // recursive func to traverse the *html.Node object. have to find the a tags and get its href
    var recursiveIterationOfHTMLNode func(*html.Node)
    recursiveIterationOfHTMLNode = func(node *html.Node) {
        // check if current node is a tag
        if node.Type == html.ElementNode && node.Data == "a" {
            for _, n := range node.Attr {
                if n.Key == "href" {
                    // the href is here.
                    // have to check, if the href has a baseurl. if it does, then just append it to result, if it doesnt,
                    // add it to the baseeurl and then append it to result.
                    fmt.Println(n.Val)
                    if n.Val[0:4] == "http" {
                        result = append(result, n.Val)
                    } else {
                        result = append(result, rawBaseURL + n.Val)
                    }
                        
                }
            }
             
        }
        
        // check its child elements by using for loop
        for child := node.FirstChild; child != nil; child = child.NextSibling {
            recursiveIterationOfHTMLNode(child)
        }
    }

    // recursive call
    recursiveIterationOfHTMLNode(parsedBody)

    
    // return the urls as a slice
    return result, nil
}


func FetchHTML(rawURL string) (string, error) {
    if rawURL == "" {
        return "",nil
    }    


    responseBody, err := http.Get(rawURL)


    if err != nil {
        return "", err
    }
    defer responseBody.Body.Close()

    httpBody, err := io.ReadAll(responseBody.Body)
    
    

    return string(httpBody) ,nil
}


func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
    
}
 










