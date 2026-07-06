package main

import "testing"

func TestNotifiers(t *testing.T) {
    // Тестируем EmailNotifier
    email := EmailNotifier{}
    if res := email.Send("test"); res != "Email: test" {
        t.Errorf("Ожидалось 'Email: test', получено '%s'", res)
    }

    // Тестируем TelegramNotifier
    tele := TelegramNotifier{}
    if res := tele.Send("test"); res != "Telegram: test" {
        t.Errorf("Ожидалось 'Telegram: test', получено '%s'", res)
    }

    // Тестируем ConsoleNotifier
    console := ConsoleNotifier{}
    if res := console.Send("test"); res != "Console: test" {
        t.Errorf("Ожидалось 'Console: test', получено '%s'", res)
    }
}
