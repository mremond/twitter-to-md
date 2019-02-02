package httpmock

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Step is used to record one of the HTTP steps in a specific query.
type Step struct {
	RequestURL string
	Response   Response
	Err        string
}

// SaveBody writes the content of the body to a separate file during the recording of a query.
// TODO: Should be on a struct ?
func (step Step) SaveBody(content io.Reader, toFile string) (err error) {
	file, err := os.Create(toFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	return err
}

// Response is a simplified http.Response version tailored for serialization.
type Response struct {
	Status       string
	StatusCode   int
	Proto        string
	ProtoMajor   int
	ProtoMinor   int
	Header       http.Header
	BodyFilename string
}

// NewResponse generates a simple response from an http.Response.
func NewResponse(resp *http.Response) Response {
	var r Response

	if resp != nil {
		r.Status = resp.Status
		r.StatusCode = resp.StatusCode
		r.Proto = resp.Proto
		r.ProtoMajor = resp.ProtoMajor
		r.ProtoMinor = resp.ProtoMinor
		r.Header = resp.Header
	}
	return r
}

// ToHTTPResponse generates an http.Response with the data from the simple response.
func (r Response) ToHTTPResponse() (*http.Response, error) {
	var resp http.Response
	resp.Status = r.Status
	resp.StatusCode = r.StatusCode
	resp.Proto = r.Proto
	resp.ProtoMajor = r.ProtoMajor
	resp.ProtoMinor = r.ProtoMinor
	resp.Header = r.Header

	if r.BodyFilename != "" {
		// TODO(mr): Fixture dir should not be hardcoded
		file, err := os.Open("fixtures/" + r.BodyFilename) // TODO(mr): Use filename join for OS independance
		if err != nil {
			return &resp, err
		}
		resp.Body = file
	} else {
		reader := bytes.NewReader([]byte{})
		resp.Body = ioutil.NopCloser(reader)
	}
	return &resp, nil
}

/*

TODO: Add the ability to have multiple sequences and to index them.
Here are the possible data types:

Key {
  Method
  URL
}

Ref {
SequenceID int
StepID int
}

Scenario {
 Sequences []Sequence
 index map[key]sequenceId
}

*/

// Sequence is used to record all the redirect steps of a single HTTP Response.
type Sequence struct {
	Steps []Step
}

// SaveTo stores JSON data for sequence.
func (s Sequence) SaveTo(filePath string) (err error) {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	return encoder.Encode(s)
}

// ReadSequence reads a JSON representation from a given file and generate a valid request sequence.
func ReadSequence(filePath string) (seq Sequence, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return seq, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&seq)
	return seq, err
}