package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*
var content embed.FS

type Link struct {
	Label string
	URL   string
}

type Stat struct {
	Value string
	Label string
	Trend string // e.g., "+12.4%"
}

type Project struct {
	Title       string
	Description string
	Tags        []string
	Link        string
}

type PageData struct {
	Name        string
	Role        string
	Tagline     string
	Socials     []Link
	Stats       []Stat
	Projects    []Project
	ResumeLink  string
	GeneratedAt string
}

func main() {
	tmpl, err := template.ParseFS(content, "templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Name:       "DAVID YAMUSA IDRIS",
			Role:       "BACKEND & AI ENGINEER",
			Tagline:    "Architecting autonomous systems and scalable infrastructure that hunt latency and capture value.",
			ResumeLink: "https://www.linkedin.com/in/david-idris-174222289/", // Or link to actual PDF if hosted
			Socials: []Link{
				{"GITHUB", "https://github.com/mrYamusa"},
				{"LINKEDIN", "https://www.linkedin.com/in/david-idris-174222289/"},
				{"TWITTER", "https://x.com/mr_yamusa"},
			},
			Stats: []Stat{
				{Value: "4+", Label: "YEARS EXP", Trend: "LEVEL UP"},
				{Value: "99.9%", Label: "UPTIME", Trend: "STABLE"},
				{Value: "12", Label: "ACTIVE DEPLOYS", Trend: "+3"},
				{Value: "K8s", Label: "ORCHESTRATION", Trend: "NATIVE"},
			},
			Projects: []Project{
				{
					Title:       "NEXTWORK RAG API",
					Description: "Production-ready Retrieval-Augmented Generation API using FastAPI, ChromaDB, and Ollama for local LLM inference.",
					Tags:        []string{"GO", "PYTHON", "K8S"},
					Link:        "https://github.com/mrYamusa",
				},
				{
					Title:       "SHIPMENT MICROSERVICE",
					Description: "Async shipment operations handling concurrent requests with FastAPI, Redis, and PostgreSQL.",
					Tags:        []string{"REDIS", "ASYNC", "POSTGRES"},
					Link:        "https://github.com/mrYamusa",
				},
				{
					Title:       "SERVERLESS AI MOD",
					Description: "Real-time streaming chat platform using Azure Functions and Google Gemini 2.0 for automated content moderation.",
					Tags:        []string{"AZURE", "GEMINI", "NoSQL"},
					Link:        "https://github.com/mrYamusa",
				},
			},
			GeneratedAt: time.Now().Format("2006-01-02 15:04:05 MST"),
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("System online at port %s", port)
	http.ListenAndServe(":"+port, nil)
}
