package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
)

// Функция генерации SVG (та самая логика)
func generateAvatar(username string) string {
	hash := md5.Sum([]byte(username))
	// Берем цвета из хеша, чтобы они были яркими (не слишком темными)
	color := fmt.Sprintf("#%02x%02x%02x", hash[0]|0x40, hash[1]|0x40, hash[2]|0x40)

	rects := ""
	for i := 0; i < 15; i++ {
		// Читаем биты хеша для определения, рисовать ли квадрат
		if hash[i%len(hash)]&(1<<(i%8)) != 0 {
			x := i / 5
			y := i % 5
			rects += fmt.Sprintf("<rect x='%d' y='%d' width='1' height='1' fill='%s' />", x, y, color)
			if x < 2 { // Зеркалим
				rects += fmt.Sprintf("<rect x='%d' y='%d' width='1' height='1' fill='%s' />", 4-x, y, color)
			}
		}
	}
	return fmt.Sprintf(`<svg viewBox="0 0 5 5" xmlns="http://w3.org" style="width:200px;height:200px;" shape-rendering="crispEdges">%s</svg>`, rects)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Достаем ник из параметра запроса /?name=твой_ник
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "anonymous" // Дефолтный ник
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Аватарка для: %s</h1>", name)
	fmt.Fprintf(w, "<div>%s</div>", generateAvatar(name))
	fmt.Fprintf(w, "<br><p>Попробуй изменить ник в ссылке: <code>/?name=user123</code></p>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Сервер Teneon запущен на http://localhost:8080")
	fmt.Println("Открой в браузере: http://localhost:8080/?name=teneon")
	http.ListenAndServe(":8080", nil)
}
