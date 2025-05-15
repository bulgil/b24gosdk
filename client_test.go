package b24gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func Test_ClientCall(t *testing.T) {
	var cases = []struct {
		name          string
		method        string
		webhook       string
		query         url.Values
		body          any
		expectedResp  any
		expectedError error
		errInfo       *errInfo
	}{
		{
			name:    "Success GET query",
			method:  http.MethodGet,
			webhook: "some.webhook",
			query:   url.Values{"q_key1": {"q_val1"}, "q_key2": {"q_val2"}},
			body:    nil,
			expectedResp: struct {
				Field1 string `json:"field1"`
				Field2 int    `json:"field2"`
				Field3 bool   `json:"field3"`
			}{
				Field1: "fourty two",
				Field2: 42,
				Field3: true,
			},
			errInfo: nil,
		},

		{
			name:    "API error response",
			method:  http.MethodGet,
			webhook: "some.webhook",
			query:   url.Values{"q_key": {"q_val"}},
			body:    nil,
			expectedResp: struct {
				Field1 string `json:"field1"`
			}{},
			errInfo:       &errInfo{Err: "INVALID_TOKEN", ErrDesc: "Access token expired"},
			expectedError: fmt.Errorf("Client.Call: bitrix api response error: INVALID_TOKEN: Access token expired"),
		},

		{
			name:    "Success POST body",
			method:  http.MethodPost,
			webhook: "deal.update",
			query:   nil,
			body: struct {
				ID     int `json:"id"`
				Fields struct {
					Title  string `json:"title"`
					Active bool   `json:"active"`
				} `json:"fields"`
			}{
				ID: 42,
				Fields: struct {
					Title  string "json:\"title\""
					Active bool   "json:\"active\""
				}{
					Title:  "test deal",
					Active: true,
				},
			},
			expectedResp: struct {
				Result struct {
					Updated bool `json:"updated"`
				} `json:"result"`
			}{},
			expectedError: nil,
			errInfo:       &errInfo{},
		},
		{
			name:    "Success GET no query parameters",
			method:  http.MethodGet,
			webhook: "no.query.webhook",
			query:   nil,
			body:    nil,
			expectedResp: struct {
				Field1 string `json:"field1"`
				Field2 int    `json:"field2"`
			}{
				Field1: "no query",
				Field2: 0,
			},
			errInfo: nil,
		},
		{
			name:    "API error response with missing field",
			method:  http.MethodGet,
			webhook: "some.webhook",
			query:   url.Values{"q_key": {"missing"}},
			body:    nil,
			expectedResp: struct {
				Field1 string `json:"field1"`
			}{},
			errInfo:       &errInfo{Err: "MISSING_FIELD", ErrDesc: "Required field is missing"},
			expectedError: fmt.Errorf("Client.Call: bitrix api response error: MISSING_FIELD: Required field is missing"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name,
			func(t *testing.T) {
				ts := initTestServer(t,
					tt.method,
					tt.webhook,
					tt.query,
					tt.body,
					tt.expectedResp,
					tt.errInfo)
				defer ts.Close()

				// u, err := url.Parse(ts.URL)
				// if err != nil {
				// 	t.Fatalf("test server url parse error: %v", err)
				// }

				client := NewClient(ts.Client(), ts.URL)

				resultPtr := reflect.New(reflect.TypeOf(tt.expectedResp))
				result := resultPtr.Interface()

				err := client.Call(tt.method, tt.webhook, tt.query, tt.body, result)

				if tt.errInfo != nil {
					if err == nil {
						t.Fatalf("expected error, got nil")
					}
					if !containsError(err.Error(), tt.errInfo.Err) || !containsError(err.Error(), tt.errInfo.ErrDesc) {
						t.Fatalf("expected error to contain %q and %q, got %q", tt.errInfo.Err, tt.errInfo.ErrDesc, err.Error())
					}
				} else {
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
					if !reflect.DeepEqual(resultPtr.Elem().Interface(), tt.expectedResp) {
						t.Fatalf("expected result %+v, got %+v", tt.expectedResp, resultPtr.Elem().Interface())
					}
				}
			})
	}
}

func containsError(msg string, want string) bool {
	return want == "" || (want != "" && msg != "" && (bytes.Contains([]byte(msg), []byte(want))))
}

func initTestServer(t *testing.T,
	expectedMethod,
	expectedPath string,
	expectedQuery url.Values,
	expectedBody,
	response any,
	errInfo *errInfo) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != expectedMethod {
				t.Fatalf("expected \"%s\" method, got %v", expectedMethod, r.Method)
			}

			if r.URL.Path != "/"+expectedPath {
				t.Fatalf("expected path %+v, got %+v", "/"+expectedPath, r.URL.Path)
			}

			if expectedQuery != nil {
				if !reflect.DeepEqual(r.URL.Query(), expectedQuery) {
					t.Fatalf("expected query %+v, got %+v", expectedQuery, r.URL.Query())
				}
			}

			if expectedBody != nil {
				reqBody := r.Body
				defer reqBody.Close()

				reqBodyMap := make(map[string]any)
				err := json.NewDecoder(reqBody).Decode(&reqBodyMap)
				if err != nil {
					t.Fatalf("test server request body unmarshal in map error: %v", err)
				}

				expectedBodyBuffer := &bytes.Buffer{}
				err = json.NewEncoder(expectedBodyBuffer).Encode(expectedBody)
				if err != nil {
					t.Fatalf("test server expected body buffer marshal error: %v", err)
				}

				expectedBodyMap := make(map[string]any)
				err = json.NewDecoder(expectedBodyBuffer).Decode(&expectedBodyMap)
				if err != nil {
					t.Fatalf("test server expected body map unmarshal error: %v", err)
				}

				if !reflect.DeepEqual(reqBodyMap, expectedBodyMap) {
					t.Fatalf("expected body %+v, got %+v", expectedBodyMap, reqBodyMap)
				}
			}

			if errInfo != nil {
				resp := map[string]any{
					"error":             errInfo.Err,
					"error_description": errInfo.ErrDesc,
					"time":              timeInfo{},
				}
				w.WriteHeader(http.StatusBadRequest)
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					t.Fatalf("error encoding error response: %v", err)
				}
				return
			}

			resp := struct {
				Result any `json:"result"`
			}{
				Result: response,
			}

			err := json.NewEncoder(w).Encode(resp)
			if err != nil {
				t.Fatalf("test server response marshal error %+v", err)
			}
		},
	))
}
