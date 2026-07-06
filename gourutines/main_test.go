package main

import (
	"testing"
)

func TestDownloadFileWithVerification(t *testing.T) {
	// Эталонный результат, который мы ОЖИДАЕМ получить
	expected := "Файл test.txt успешно скачан!"

	// Вызываем функцию напрямую (без "go"), передавая nil вместо WaitGroup,
	// так как внутри функции мы добавили проверку "if wg != nil"
	result := downloadFile("test.txt", 5, nil)

	// НАСТОЯЩАЯ ПРОВЕРКА (Утверждение / Assertion)
	if result != expected {
		// Если результат не совпал с ожиданиями, мы явно валим тест
		// и выводим красивую понятную ошибку разработчику
		t.Errorf("Ошибка проверки! \nОжидалось: %q\nПолучено:  %q", expected, result)
	}
}
