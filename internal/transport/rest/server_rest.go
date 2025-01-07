package rest

import (
	"fmt"
	"net/http"
)

// Регистрация обработчиков
func registerHandlers() {

}

// Запуск сервера
func StartServer() {
	registerHandlers()

	port := ":8080"
	fmt.Printf("Сервер запущен на http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}