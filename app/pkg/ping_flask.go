package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Payload представляет структуру данных, которую вы отправляете.
type Payload struct {
	Data string `json:"data"`
}

// sendRequest отправляет HTTP-запрос на указанный URL с данными в JSON-формате.
func sendRequest(url string, payload Payload) (string, error) {
	// Преобразуем данные в JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("ошибка кодирования JSON: %v", err)
	}

	// Создаём новый HTTP-запрос
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("ошибка создания HTTP-запроса: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос с помощью HTTP-клиента
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка отправки HTTP-запроса: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	// Возвращаем тело ответа как строку
	return string(body), nil
}

type MessageResponse struct {
	Context string `json:"context"`
}

// sendMessage отправляет сообщение и получает ответ от сервера
func SendMessage(url string, message string) (MessageResponse, error) {
	// Создаем тело запроса
	payload := map[string]string{
		"api_key": "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",
		"prompt":  message,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return MessageResponse{}, fmt.Errorf("ошибка при сериализации JSON: %v", err)
	}

	// Выполняем POST-запрос
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return MessageResponse{}, fmt.Errorf("ошибка при выполнении POST-запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return MessageResponse{}, fmt.Errorf("получен некорректный код ответа: %d", resp.StatusCode)
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MessageResponse{}, fmt.Errorf("ошибка при чтении тела ответа: %v", err)
	}

	// Парсим JSON-ответ
	var response MessageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return MessageResponse{}, fmt.Errorf("ошибка при десериализации JSON-ответа: %v", err)
	}

	// Возвращаем ответ
	return response, nil
}
