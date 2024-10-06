package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Создаем структуру для ответа
	type response struct {
		Message string `json:"message"`
	}

	// Заполняем структуру
	resp := response{
		Message: "Hello, World!",
	}

	// Устанавливаем Content-Type в заголовках
	w.Header().Set("Content-Type", "application/json")

	// Преобразуем структуру в JSON и отправляем ответ
	json.NewEncoder(w).Encode(resp)
}
