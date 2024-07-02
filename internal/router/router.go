package router

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"sync"
)

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      int               `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var (
	currentRequestID int
	mu               sync.Mutex
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthCheck", hello)
		r.Get("/external-api", proxy)
	})

	return router
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func proxy(w http.ResponseWriter, r *http.Request) {
	var proxyRequest ProxyRequest

	if err := json.NewDecoder(r.Body).Decode(&proxyRequest); err != nil {
		http.Error(w, "Ошибка декодирования тела запроса", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(proxyRequest.URL)
	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Ошибка в чтений тело запроса", http.StatusInternalServerError)
		return
	}

	status := resp.StatusCode

	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = values[0]
	}

	mu.Lock()
	currentRequestID++
	requestID := currentRequestID
	mu.Unlock()

	response := ProxyResponse{
		ID:      requestID,
		Status:  status,
		Headers: headers,
		Length:  len(body),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Не удалось вывести ответ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
