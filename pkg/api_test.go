package weatherbit

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMisc(t *testing.T) {
	ts := getTestServer("Hello")
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

func TestGetResponse(t *testing.T) {

	type args struct {
		ep Parameters
	}
	tests := []struct {
		name string
		args args
		want WbResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetResponse(tt.args.ep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTestServerWithJSONData(filename string) *httptest.Server {
	bytes, _ := ioutil.ReadFile("testdata/" + filename)
	return getTestServer(string(bytes))
}

func getTestServer(data string) *httptest.Server {
	httphandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, data)
	})
	testserver := httptest.NewServer(httphandler)
	return testserver
}
