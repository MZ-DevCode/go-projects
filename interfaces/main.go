package main

import "fmt"

// Интерфейс Notifier
type Notifier interface {
    Send(message string) string
}

// EmailNotifier
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) string {
    return "Email: " + message
}

// TelegramNotifier
type TelegramNotifier struct{}

func (t TelegramNotifier) Send(message string) string {
    return "Telegram: " + message
}

// ConsoleNotifier
type ConsoleNotifier struct{}

func (c ConsoleNotifier) Send(message string) string {
    return "Console: " + message
}

// Функция для отправки всем
func NotifyAll(notifiers []Notifier, message string) {
    for _, n := range notifiers {
        fmt.Println(n.Send(message))
    }
}

func main() {
    notifiers := []Notifier{
        EmailNotifier{},
        TelegramNotifier{},
        ConsoleNotifier{},
    }
    NotifyAll(notifiers, "Привет, мир!")
}
