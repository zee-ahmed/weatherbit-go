package weatherbit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_syncHTTPGets(t *testing.T) {
	type args struct {
		p parameters
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
			if got := syncHTTPGets(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("syncHTTPGets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_syncHTTPGets2(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}

}
