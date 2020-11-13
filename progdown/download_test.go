package progdown

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type client struct {
	client *http.Client
}

func (c *client) Do(req *http.Request) (*http.Response, error) {
	fmt.Println(req.Method, req.URL.String(), req.Header)
	return c.client.Do(req)
}

func TestDownload(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./testdata/100bytes")
	}))
	defer svr.Close()

	client := client{
		client: http.DefaultClient,
	}
	r, err := Download(&client, "ua", svr.URL+"/testdata/100bytes", 25, 7)
	if err != nil {
		t.Fatal("download failed:", err)
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal("read error:", err)
	}
	if want, got := 100, len(b); want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}

	want, err := ioutil.ReadFile("./testdata/100bytes")
	if diff := cmp.Diff(string(want), string(b)); diff != "" {
		t.Errorf("download content mismatch (-want +got):\n%s", diff)
	}
}
