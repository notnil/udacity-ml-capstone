package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	clipPath string
	dbPath   string
	port     string

	db *sql.DB
)

func init() {
	flag.StringVar(&clipPath, "clip_path", "../data/clips", "path to clips that need labeled")
	flag.StringVar(&dbPath, "db_path", "labels.db", "path to sqlite db")
	flag.StringVar(&port, "port", "8080", "port to listen")
}

func main() {
	flag.Parse()

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/video/", videoHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	viewID := r.URL.Path[len("/save/"):]
	parts := strings.Split(viewID, "/")
	videoID := parts[0]
	clipID := parts[1]

	body := r.FormValue("shot")
	i, err := strconv.ParseInt(body, 10, 64)
	if err != nil {
		panic(err)
	}

	clipInt, err := strconv.ParseInt(clipID, 10, 64)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO labels (video_id, clip_id, label) VALUES ($1, $2, $3)", videoID, clipInt, i)
	if err != nil {
		panic(err)
	}

	title := fmt.Sprintf("%s/%05d", videoID, clipInt+1)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	viewID := r.URL.Path[len("/video/"):]
	parts := strings.Split(viewID, "/")
	videoID := parts[0]
	clipID := parts[1]
	path := fmt.Sprintf("%s/%s/%s.mp4", clipPath, videoID, clipID)
	http.ServeFile(w, r, path)
}

// http://localhost:8080/view/dEtg6urGQBo/00000
func viewHandler(w http.ResponseWriter, r *http.Request) {
	viewID := r.URL.Path[len("/view/"):]
	parts := strings.Split(viewID, "/")
	videoID := parts[0]
	clipID := parts[1]
	type labelView struct {
		VideoID  string
		ClipID   string
		VideoURL string
	}
	v := &labelView{
		VideoID:  videoID,
		ClipID:   clipID,
		VideoURL: fmt.Sprintf("/video/%s", viewID),
	}
	t, err := template.ParseFiles("label.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, v)
}
