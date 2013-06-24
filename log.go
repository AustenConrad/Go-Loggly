package gologgly

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

type entryRoot struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Tags      []string  `json:"tags"`
	File      entryFile `json:"file"`
	Request   entryUser `json:"request,omitempty"`
}
type entryUser struct {
	Host  string `json:"host,omitempty"`
	Agent string `json:"user_agent,omitempty"`
}
type entryFile struct {
	Name   string `json:"name"`
	Line   int    `json:"line"`
	Path   string `json:"path"`
	Memory string `json:"memory_address"`
}

// Creates a log entry on Loggly using the specified input within the input map and setting the specified tags,
// specified message, file name, file line number, memory address, time, user agent, and user ip address/host.
func Log(input string, tags []string, message string, rw http.ResponseWriter, req *http.Request) (err error) {

	// Initialize an appengine context.
	c := appengine.NewContext(req)

	// Get the error time right away so that it's as accurate as possible.
	timestamp := time.Now()

	// Initialize a new entry.
	var entry entryRoot

	// Lookup the file line, file name, and memory address of where Loggly was called.
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???.go"
		line = -1
	}

	// Remove the directory from the file name.
	short := file
	path := ""
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			path = file[:i+1]
			break
		}
	}
	file = short

	// Assemble output.
	entry.Message = message
	entry.Timestamp = timestamp
	entry.Tags = tags
	entry.Request.Host = req.RemoteAddr
	entry.Request.Agent = req.UserAgent()
	entry.File.Name = file
	entry.File.Path = path
	entry.File.Line = line
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "0x%x", pc)
	entry.File.Memory = buf.String()

	// Construct the body of the request.
	entry_json, err := json.Marshal(entry)
	if err != nil {
		c.Errorf(err.Error())
		return err
	}
	body := bytes.NewReader(entry_json)

	// Retrieve the input's url from the globally defined input map.
	url := "https://logs.loggly.com/inputs/" + loggly[input]
	c.Infof("url: %s", url)
	// Initial a http.Client with an appengine/urlfetch transport.
	client := urlfetch.Client(c)

	// Construct the request.
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		c.Errorf(err.Error())
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	// Send log message to Loggly.
	response, err := client.Do(request)
	if err != nil {
		c.Errorf(err.Error())
		return err
	}

	// Handle response.
	if response.StatusCode != 201 {
		c.Errorf("Loggly log did not send: '" + string(entry_json) + "'" + " Loggly HTTP response was code: " + strconv.Itoa(response.StatusCode))
		return errors.New("Loggly log did not send: '" + string(entry_json) + "'" + " Loggly HTTP response was code: " + strconv.Itoa(response.StatusCode))
	}

	// Success!
	return nil
}
