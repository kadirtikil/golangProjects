package webcrawler 

import( 
    "testing"
    "reflect"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
        // add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}


func TestGetUrlsFromHTML(t *testing.T) {
    tests := []struct{
        name string
        inputURL string
        inputHTMLBody string
        expected []string
    }{
        {
            name: "empty html body and url",
            inputURL: "",
            inputHTMLBody: "",
            expected: nil, 
        },
        {
            name:     "absolute and relative URLs",
            inputURL: "https://blog.boot.dev",
            inputHTMLBody: `
            <html>
	            <body>
		            <a href="/path/one">
			            <span>Boot.dev</span>
		            </a>
		            <a href="https://other.com/path/one">
			            <span>Boot.dev</span>
		            </a>
	            </body>
            </html>`,
            expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
        },
        {
            name:     "another unneccesary case",
            inputURL: "https://kadir.wants.to.be.a.dev",
            inputHTMLBody: `
            <html>
                <body>
                    <a href="/path/one">
                        <span>Link One</span>
                    </a>
                    <a href="https://example.com/path/two">
                        <span>Link Two</span>
                    </a>
                    <a href="/another/path/three">
                        <span>Link Three</span>
                    </a>
                </body>
            </html>`,
            expected: []string{"https://kadir.wants.to.be.a.dev/path/one", "https://example.com/path/two", "https://kadir.wants.to.be.a.dev/another/path/three"},
        },
    }
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLsFromHTML(tc.inputHTMLBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}


// Test cases for this are mad annoying because the html of websites is not always static and prone to change.
// static websites are far too big. furthermore, the functionality is kinda obvious anyway so we good.
func TestFetchURL(t *testing.T) {
    tests := []struct {
        name string
        inputURL string
        expected string
    }{
        {
            name: "empty object",
            inputURL: "",
            expected: "",
        },
        {
            name: "some website",
            inputURL: "",
            expected: "", 
            
        },
        {
            name: "some website that could change",
            inputURL: "",
            expected: "",
        },
    }
    
    for i, tc := range tests{
        t.Run(tc.name, func(t *testing.T) {
			actual, err := FetchHTML(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})

    }

}
