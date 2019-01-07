package weatherbit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func getMockServerWithFileData(filename string) *httptest.Server {
	bytes, _ := ioutil.ReadFile("testdata/" + filename)
	return serveMockServer(string(bytes))
}

func serveMockServer(data string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, data)
	}))

	return server
}

func Test_syncHTTPGets(t *testing.T) {
	type args struct {
		p Parameters
	}

	mockserver := getMockServerWithFileData("current.json")
	defer mockserver.Close()

	tests := []struct {
		name string
		args args
		want WbResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := syncHTTPGets(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("syncHTTPGets() = %v, want %v", got, tt.want)
			}
		})
	}
}
