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
                    if len(n.Val) >= 4 && n.Val[0:4] == "http" {
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
	// crawl as long as the urls got the same domain. otherwise itll jump domains, which can potentially
    // lead to a whole lot of crawling
    currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}
    
    // get the normalized url to make it more readable and save in the map with a value showing how often said url appeared as an href
	normalizedURL, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
	}

	// if already in map increase counter
	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL]++
		return
	}

	// if not in map, then make a new index, and init counter with 1 
	pages[normalizedURL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

    // get the html
	htmlBody, err := FetchHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

    // get the hrefs from a tags out of the html body and save them in a slice
	nextURLs, err := GetURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
	}

    // iterate the slice and call the crawler again.
	for _, nextURL := range nextURLs {
		CrawlPage(rawBaseURL, nextURL, pages)
	}
} 










