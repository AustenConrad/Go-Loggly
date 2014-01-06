### Go-Loggly
Sample Go project demonstrating how to send JSON-formatted log data to Loggly from an app deployed on the Google App Engine.

Creates a log entry on Loggly using the specified input within the input map. Logs the specified tags, specified message, file name, file line number, memory address, time, user agent, and user ip address/host.

Licensed under the MIT License. See License.txt file.


###### Usage:
Set your GOPATH and then install via:
```bash
go get github.com/AustenConrad/Go-Loggly
```

Use in your application via:
```go
// Import library.
import (
	gologgly "github.com/AustenConrad/Go-Loggly"
)

// Add named input(s). Typically done in the main() function.
gologgly.AddInput("input name", "SHA-2 key from Loggly dashboard")

// Log anywhere within the application.
gologgly.Log("input name", tags []string{"a tag", "another tag"}, "message you would like to store", rw, req)
```
Stores in Loggy as:
```json
{
	"timestamp":"2013-06-24T03:57:03.982052Z",
	"message":"message you would like to store",
	"tags":[
		"a tag",
		"another tag"
	],
	"file":{
		"name":"filename.go",
		"line":218,
		"path":"current/path/to/file",
		"memory_address":"0xaba53"
	},
	"request":{
		"host":"8.8.8.8",
		"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0"
	}
} 
```

###### Sample Implementation:
```go

import (
	"github.com/AustenConrad/Go-Loggly/gologgly"
)

func main() {
	// Store a few Loggly inputs for reference by other parts of the application.
	gologgly.AddInput("errors", "SHA-2 key from Loggly dashboard")
	gologgly.AddInput("uploads", "SHA-2 key from Loggly dashboard")
	gologgly.AddInput("catsFTW", "SHA-2 key from Loggly dashboard")
}

func someUploadFunc(user string, rw http.ResponseWriter, req *http.Request) (err error) {
	
	// Log an upload starting.
	// EXAMPLE: For example, to find all of the uploads that have started using the Loggly console: 'search json.tags:started'
	err := gologgly.Log("uploads", []string{"Started", user, "Another tag"}, "{'some':'json', 'more': 'json stuff'}", rw, req)
	if err != nil {
		return err
	}

	return nil
}

func someErrorHandler(rw http.ResponseWriter, req *http.Request) (err error) {

	// Log an error.
	err := gologgly.Log("errors", []string{"CRITICAL"}, "OMG where is the cat!?!", rw, req)
	if err != nil {
		return err
	}

	err := gologgly.Log("errors", []string{"EMERGENCY"}, "(to cat) Seriously, *how* did you get up here?", rw, req)
	if err != nil {
		return err
	}

	return nil
}

```
