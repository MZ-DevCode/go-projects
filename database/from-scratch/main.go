package main

import (
	"bufio"        // Импорт пакета для удобного чтения строк из стандартного ввода (консоли)
	"database/sql" // Импорт пакета для работы с SQL-базами данных
	"fmt"          // Импорт пакета для форматированного вывода текста в консоль
	"log"          // Импорт пакета для ведения логов и вывода ошибок
	"os"           // Импорт пакета для взаимодействия с операционной системой
	"strings"      // Импорт пакета для манипуляции строками (например, перевод в нижний регистр)

	_ "modernc.org/sqlite" // Импорт драйвера SQLite: символ '_' означает, что мы подключаем драйвер,
	// но не используем его напрямую в коде, он сам регистрируется в пакете sql
)

func main() {
	// Открываем соединение с файлом базы данных. Если файла нет, SQLite создаст его сам.
	db, err := sql.Open("sqlite", "./library.db")
	if err != nil {
		log.Fatal(err) // Если не удалось открыть БД, завершаем программу с ошибкой
	}
	defer db.Close() // Гарантируем, что БД закроется автоматически, когда завершится функция main

	// Создаем таблицу авторов, если она еще не существует.
	// id — это первичный ключ, который автоматически увеличивается при каждой записи.
	db.Exec(`CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);`)

	// Создаем таблицу книг.
	// author_id — это связь (внешний ключ), которая указывает, какому автору принадлежит книга.
	db.Exec(`CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author_id INTEGER,
		FOREIGN KEY(author_id) REFERENCES authors(id)
	);`)

	// Инициализируем сканер для чтения ввода от пользователя через консоль
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Библиотека готова. Команды: add, list, exit")

	// Запускаем бесконечный цикл для работы меню
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		scanner.Scan()                           // Ждем ввода строки от пользователя
		input := strings.ToLower(scanner.Text()) // Приводим ввод к нижнему регистру для сравнения

		// Используем переключатель для выбора действия пользователя
		switch input {
		case "add":
			// Считываем имя автора
			fmt.Print("Имя автора: ")
			scanner.Scan()
			author := scanner.Text()

			// Считываем название книги
			fmt.Print("Название книги: ")
			scanner.Scan()
			title := scanner.Text()

			// Добавляем автора в базу. Символ '?' — это плейсхолдер для защиты от SQL-инъекций.
			res, err := db.Exec("INSERT INTO authors (name) VALUES (?)", author)
			if err != nil {
				log.Println("Ошибка при добавлении автора:", err)
				continue // Возвращаемся в начало цикла, если произошла ошибка
			}

			// Получаем уникальный ID, который база данных только что создала для этого автора
			authorID, _ := res.LastInsertId()

			// Добавляем книгу, привязывая её к полученному authorID
			_, err = db.Exec("INSERT INTO books (title, author_id) VALUES (?, ?)", title, authorID)
			if err != nil {
				log.Println("Ошибка при добавлении книги:", err)
				continue
			}
			fmt.Println("Записано в базу!")

		case "list":
			fmt.Println("--- Список всех книг ---")
			// Выполняем SQL-запрос, объединяющий таблицы книг и авторов по полю связи (author_id = id)
			rows, err := db.Query(`SELECT books.title, authors.name FROM books JOIN authors ON books.author_id = authors.id`)
			if err != nil {
				log.Println("Ошибка при запросе:", err)
				continue
			}
			defer rows.Close() // Закрываем результат запроса, когда закончим работу с ним

			// Перебираем все полученные строки по одной
			for rows.Next() {
				var title, author string
				// Извлекаем данные из текущей строки в переменные
				err := rows.Scan(&title, &author)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Книга: %s | Автор: %s\n", title, author)
			}

		case "exit":
			fmt.Println("Выход из программы...")
			return // Завершаем выполнение функции main и выходим из программы

		default:
			fmt.Println("Неизвестная команда. Доступны: add, list, exit")
		}
		if err := scanner.Err(); err != nil {
			log.Println("Ошибка чтения из консоли:", err)
		}
	}
}
