package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var tmpl *template.Template

// ─── Models ───────────────────────────────────────────────────────────────────

type Subject struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	CreatedAt   string `json:"created_at"`
}

type Grade struct {
	ID          int    `json:"id"`
	SubjectID   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Score       int    `json:"score"`
	MaxScore    int    `json:"max_score"`
	Topic       string `json:"topic"`
	Details     string `json:"details"`
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
}

type Topic struct {
	ID          int    `json:"id"`
	SubjectID   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Status      string `json:"status"` // planned, in_progress, done
	Date        string `json:"date"`
	CreatedAt   string `json:"created_at"`
}

type PageData struct {
	Subjects []Subject
	Grades   []Grade
	Topics   []Topic
	Stats    Stats
}

type Stats struct {
	TotalSubjects int
	TotalTopics   int
	DoneTopics    int
	AvgScore      float64
}

// ─── Database ─────────────────────────────────────────────────────────────────

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./studylog.db")
	if err != nil {
		log.Fatal(err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS subjects (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		name        TEXT NOT NULL,
		description TEXT DEFAULT '',
		color       TEXT DEFAULT '#6366f1',
		created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS grades (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		subject_id  INTEGER NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
		score       INTEGER NOT NULL DEFAULT 0,
		max_score   INTEGER NOT NULL DEFAULT 10,
		topic       TEXT DEFAULT '',
		details     TEXT DEFAULT '',
		date        DATE NOT NULL,
		created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS topics (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		subject_id  INTEGER NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
		title       TEXT NOT NULL,
		content     TEXT DEFAULT '',
		status      TEXT DEFAULT 'planned',
		date        DATE NOT NULL,
		created_at  DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err = db.Exec(schema); err != nil {
		log.Fatal(err)
	}

}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func jsonErr(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func jsonOK(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func getSubjects() ([]Subject, error) {
	rows, err := db.Query("SELECT id,name,description,color,created_at FROM subjects ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Subject
	for rows.Next() {
		var s Subject
		rows.Scan(&s.ID, &s.Name, &s.Description, &s.Color, &s.CreatedAt)
		list = append(list, s)
	}
	return list, nil
}

func getGrades() ([]Grade, error) {
	rows, err := db.Query(`
		SELECT g.id, g.subject_id, s.name, g.score, g.max_score, g.topic, g.details, g.date, g.created_at
		FROM grades g JOIN subjects s ON s.id=g.subject_id
		ORDER BY g.date DESC, g.created_at DESC LIMIT 200`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Grade
	for rows.Next() {
		var g Grade
		rows.Scan(&g.ID, &g.SubjectID, &g.SubjectName, &g.Score, &g.MaxScore, &g.Topic, &g.Details, &g.Date, &g.CreatedAt)
		list = append(list, g)
	}
	return list, nil
}

func getTopics() ([]Topic, error) {
	rows, err := db.Query(`
		SELECT t.id, t.subject_id, s.name, t.title, t.content, t.status, t.date, t.created_at
		FROM topics t JOIN subjects s ON s.id=t.subject_id
		ORDER BY t.date DESC, t.created_at DESC LIMIT 300`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []Topic
	for rows.Next() {
		var t Topic
		rows.Scan(&t.ID, &t.SubjectID, &t.SubjectName, &t.Title, &t.Content, &t.Status, &t.Date, &t.CreatedAt)
		list = append(list, t)
	}
	return list, nil
}

func getStats() Stats {
	var s Stats
	db.QueryRow("SELECT COUNT(*) FROM subjects").Scan(&s.TotalSubjects)
	db.QueryRow("SELECT COUNT(*) FROM topics").Scan(&s.TotalTopics)
	db.QueryRow("SELECT COUNT(*) FROM topics WHERE status='done'").Scan(&s.DoneTopics)
	db.QueryRow("SELECT COALESCE(AVG(CAST(score AS REAL)/max_score*10),0) FROM grades").Scan(&s.AvgScore)
	return s
}

// ─── Handlers ─────────────────────────────────────────────────────────────────

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// --- Subjects ---

func handleSubjects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list, err := getSubjects()
		if err != nil {
			jsonErr(w, err.Error(), 500)
			return
		}
		jsonOK(w, list)

	case http.MethodPost:
		var s Subject
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			jsonErr(w, "bad json", 400)
			return
		}
		if s.Name == "" {
			jsonErr(w, "name required", 400)
			return
		}
		if s.Color == "" {
			s.Color = "#6366f1"
		}
		res, err := db.Exec("INSERT INTO subjects(name,description,color) VALUES(?,?,?)", s.Name, s.Description, s.Color)
		if err != nil {
			jsonErr(w, err.Error(), 500)
			return
		}
		id, _ := res.LastInsertId()
		s.ID = int(id)
		jsonOK(w, s)

	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			jsonErr(w, "id required", 400)
			return
		}
		db.Exec("DELETE FROM subjects WHERE id=?", id)
		jsonOK(w, map[string]bool{"ok": true})

	case http.MethodPut:
		var s Subject
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			jsonErr(w, "bad json", 400)
			return
		}
		db.Exec("UPDATE subjects SET name=?,description=?,color=? WHERE id=?", s.Name, s.Description, s.Color, s.ID)
		jsonOK(w, s)
	}
}

// --- Grades ---

func handleGrades(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list, err := getGrades()
		if err != nil {
			jsonErr(w, err.Error(), 500)
			return
		}
		jsonOK(w, list)

	case http.MethodPost:
		var g Grade
		if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
			jsonErr(w, "bad json", 400)
			return
		}
		if g.Date == "" {
			g.Date = time.Now().Format("2006-01-02")
		}
		if g.MaxScore == 0 {
			g.MaxScore = 10
		}
		res, err := db.Exec("INSERT INTO grades(subject_id,score,max_score,topic,details,date) VALUES(?,?,?,?,?,?)",
			g.SubjectID, g.Score, g.MaxScore, g.Topic, g.Details, g.Date)
		if err != nil {
			jsonErr(w, err.Error(), 500)
			return
		}
		id, _ := res.LastInsertId()
		g.ID = int(id)
		jsonOK(w, g)

	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		db.Exec("DELETE FROM grades WHERE id=?", id)
		jsonOK(w, map[string]bool{"ok": true})

	case http.MethodPut:
		var g Grade
		if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
			jsonErr(w, "bad json", 400)
			return
		}
		db.Exec("UPDATE grades SET subject_id=?,score=?,max_score=?,topic=?,details=?,date=? WHERE id=?",
			g.SubjectID, g.Score, g.MaxScore, g.Topic, g.Details, g.Date, g.ID)
		jsonOK(w, g)
	}
}

// --- Topics ---

func handleTopics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		list, err := getTopics()
		if err != nil {
			jsonErr(w, err.Error(), 500)
			return
		}
		jsonOK(w, list)

	case http.MethodPost:
		var t Topic
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			jsonErr(w, "bad json", 400)
			return
		}
		if t.Date == "" {
			t.Date = time.Now().Format("2006-01-02")
		}
		if t.Status == "" {
			t.Status = "planned"
		}
		res, err := db.Exec("INSERT INTO topics(subject_id,title,content,status,date) VALUES(?,?,?,?,?)",
			t.SubjectID, t.Title, t.Content, t.Status, t.Date)
		if err != nil {
			jsonErr(w, err.Error(), 500)
			return
		}
		id, _ := res.LastInsertId()
		t.ID = int(id)
		jsonOK(w, t)

	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		db.Exec("DELETE FROM topics WHERE id=?", id)
		jsonOK(w, map[string]bool{"ok": true})

	case http.MethodPut:
		var t Topic
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			jsonErr(w, "bad json", 400)
			return
		}
		db.Exec("UPDATE topics SET subject_id=?,title=?,content=?,status=?,date=? WHERE id=?",
			t.SubjectID, t.Title, t.Content, t.Status, t.Date, t.ID)
		jsonOK(w, t)
	}
}

// --- Stats ---

func handleStats(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, getStats())
}

// --- Grades by subject ---

func handleGradesBySubject(w http.ResponseWriter, r *http.Request) {
	sid := r.URL.Query().Get("subject_id")
	if sid == "" {
		handleGrades(w, r)
		return
	}
	rows, err := db.Query(`
		SELECT g.id, g.subject_id, s.name, g.score, g.max_score, g.topic, g.details, g.date, g.created_at
		FROM grades g JOIN subjects s ON s.id=g.subject_id
		WHERE g.subject_id=? ORDER BY g.date DESC`, sid)
	if err != nil {
		jsonErr(w, err.Error(), 500)
		return
	}
	defer rows.Close()
	var list []Grade
	for rows.Next() {
		var g Grade
		rows.Scan(&g.ID, &g.SubjectID, &g.SubjectName, &g.Score, &g.MaxScore, &g.Topic, &g.Details, &g.Date, &g.CreatedAt)
		list = append(list, g)
	}
	jsonOK(w, list)
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(204)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func intParam(r *http.Request, key string) int {
	v, _ := strconv.Atoi(r.URL.Query().Get(key))
	return v
}

func main() {
	initDB()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/api/subjects", handleSubjects)
	mux.HandleFunc("/api/grades", handleGrades)
	mux.HandleFunc("/api/topics", handleTopics)
	mux.HandleFunc("/api/stats", handleStats)
	mux.HandleFunc("/api/grades/by-subject", handleGradesBySubject)

	log.Println("🚀  StudyLog запущен → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
