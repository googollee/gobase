package progdown

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type reader struct {
	client      HTTPClient
	downloadURL string
	ua          string
	length      int64
	nextPos     int64
	firstSize   int
	chunkSize   int
	currentResp *http.Response
}

func Download(client HTTPClient, ua, url string, firstSize, chunkSize int) (io.ReadCloser, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}

	if ua != "" {
		req.Header.Set("User-Agent", ua)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't download %s, response: (%d) %s", url, resp.StatusCode, resp.Status)
	}
	if resp.ContentLength < 0 {
		return nil, fmt.Errorf("can't progressive download from %s", url)
	}
	acceptRange := resp.Header.Get("Accept-Ranges")
	if acceptRange != "bytes" {
		return nil, fmt.Errorf("can't progressive download from %s with bytes", url)
	}

	return &reader{
		client:      client,
		downloadURL: url,
		ua:          ua,
		length:      resp.ContentLength,
		nextPos:     0,
		firstSize:   firstSize,
		chunkSize:   chunkSize,
	}, nil
}

func (r *reader) Read(b []byte) (int, error) {
BEGIN:
	if r.currentResp == nil {
		var chunk int
		if r.nextPos == 0 {
			chunk = r.firstSize
		} else {
			chunk = r.chunkSize
		}
		if err := r.downloadChunk(chunk); err != nil {
			return 0, err
		}
	}

	n, err := r.currentResp.Body.Read(b)
	if err == io.EOF {
		r.currentResp.Body.Close()
		r.currentResp = nil
		if n == 0 {
			goto BEGIN
		}
		return n, nil
	}
	if err != nil {
		return n, err
	}

	return n, nil
}

func (r *reader) downloadChunk(size int) error {
	if r.nextPos > r.length {
		return io.EOF
	}

	req, err := http.NewRequest("GET", r.downloadURL, nil)
	if err != nil {
		return err
	}

	end := r.nextPos + int64(size)
	if end > r.length {
		end = r.length
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", r.nextPos, end))
	req.Header.Add("User-Agent", r.ua)

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}

	r.currentResp = resp
	r.nextPos = end + 1
	return nil
}

func (r *reader) Close() error {
	if r.currentResp == nil {
		return nil
	}
	return r.currentResp.Body.Close()
}
