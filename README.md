### gologgly
Sample Go project demonstrating how to send JSON-formatted log data to Loggly from an app deployed on the Google App Engine.

Creates a log entry on Loggly using the specified input within the input map. Logs the specified tags, specified message, file name, file line number, memory address, time, user agent, and user ip address/host.

Licensed under the MIT License. See License.txt file.


###### Usage:
```go
// Import library.
import ("github.com/AustenConrad/gologgly")

// Add named input(s). Typically done in the main() function.
loggly.AddInput("input name", "SHA-2 key from Loggly dashboard")

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
	"github.com/AustenConrad/gologgly"
)

func main() {
	
	// Store a few Loggly inputs for reference by other parts of the application.
	loggly.AddInput("errors", "SHA-2 key from Loggly dashboard")
	loggly.AddInput("uploads", "SHA-2 key from Loggly dashboard")
	loggly.AddInput("catsFTW", "SHA-2 key from Loggly dashboard")
}

func someUploadFunc(user string) (err error) {
	
	// Log an upload starting.
	// EXAMPLE: For example, to find all of the uploads that have started using the Loggly console: 'search json.tags:started'
	err := loggly.Log("uploads", []string{"Started", user, "Another tag"}, "{'some':'json', 'more': 'json stuff'}", rw, req)
	if err != nil {
		return err
	}

	return nil
}

func someErrorHandler() {

	// Log an error.
	err := loggly.Log("errors", []string{"CRITICAL"}, "OMG where is the cat!?!", rw, req)
	if err != nil {
		return err
	}

	err := loggly.Log("errors", []string{"EMERGENCY"}, "(to cat) Seriously, *how* did you get up here?", rw, req)
	if err != nil {
		return err
	}

	return nil
}

```
