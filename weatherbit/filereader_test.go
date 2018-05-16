package weatherbit

// https://golang.org/src/io/ioutil/ioutil_test.go

import (
	"os"
	"testing"
)

func checkSize(t *testing.T, path string, size int64) {
	dir, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat %q (looking for size %d): %s", path, size, err)
	}
	if dir.Size() != size {
		t.Errorf("Stat %q: size %d want %d", path, dir.Size(), size)
	}
}

func TestReadFile(t *testing.T) {
	filename := "api_key.txt"
	contents, err := Readapikey(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: error expected, none found", filename)
	}

	checkSize(t, filename, int64(len(contents)))

	filename = "filereader_test.go"
	contents, err = Readapikey(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", filename, err)
	}

}

// func TestReadapikey(t *testing.T) {
// 	type args struct {
// 		apikeyfile string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add more test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got, _ := Readapikey(tt.args.apikeyfile); got != tt.want {
// 				t.Errorf("Readapikey() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
