package b24gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client *Client

type HTTPClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient HTTPClient
	baseURL    *url.URL
}

func NewClient(httpClient HTTPClient, baseURL string) *Client {
	const op = "NewClient"

	if client != nil {
		return client
	}

	if strings.TrimSpace(baseURL) == "" {
		panic("client domain must not be empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 5,
		}
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		panic(fmt.Sprintf("%s: %v", op, err))
	}

	return &Client{
		httpClient: httpClient,
		baseURL:    u,
	}
}

func (c *Client) Call(method, webhook string, query url.Values, body, result any) error {
	const op = "Client.Call"

	url := url.URL{
		Scheme:   c.baseURL.Scheme,
		Host:     c.baseURL.Host,
		Path:     webhook,
		RawQuery: query.Encode(),
	}

	var reqBody io.Reader
	if body != nil {
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return fmt.Errorf("%s: body marshal error: %w", op, err)
		}
		reqBody = buf
	}

	req, err := http.NewRequest(method, url.String(), reqBody)
	if err != nil {
		return fmt.Errorf("%s: request create error: %w", op, err)
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s: response call error %w", op, err)
	}
	defer resp.Body.Close()

	var response response

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("%s: response body decode error: %w", op, err)
	}

	if response.ErrInfo != nil {
		return fmt.Errorf("%s: bitrix api response error: %w", op, B24Error{Err: response.ErrInfo.Err, ErrDescription: response.ErrInfo.ErrDesc})
	}

	err = json.NewDecoder(bytes.NewReader(response.Result)).Decode(&result)
	if err != nil {
		return fmt.Errorf("%s: result unmarshal error: %w", op, err)
	}

	return nil
}
