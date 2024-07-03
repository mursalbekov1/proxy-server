package proxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type RequestProxy struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ResponseProxy struct {
	ID      int               `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var (
	responseID   int
	mu           sync.Mutex
	requestData  sync.Map
	responseData sync.Map
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func Proxy(w http.ResponseWriter, r *http.Request) {
	var proxyRequest RequestProxy

	if err := json.NewDecoder(r.Body).Decode(&proxyRequest); err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(proxyRequest.URL)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	status := resp.StatusCode

	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = values[0]
	}

	mu.Lock()
	responseID++
	mu.Unlock()

	response := ResponseProxy{
		ID:      responseID,
		Status:  status,
		Headers: headers,
		Length:  len(body),
	}

	proxyResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	// Save request and response data in sync.Map
	requestData.Store(responseID, proxyRequest)
	responseData.Store(responseID, proxyResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(proxyResponse)
}
