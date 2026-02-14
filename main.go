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
	Trend string
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
			ResumeLink: "https://www.linkedin.com/in/david-idris-174222289/",
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
					Description: "Production-ready Document Q&A API using FastAPI, ChromaDB, and Ollama. Orchestrated with Kubernetes.",
					Tags:        []string{"PYTHON", "K8S", "OLLAMA"},
					Link:        "https://github.com/mrYamusa/nextwork-rag-api",
				},
				{
					Title:       "SHIPMENT MICROSERVICE",
					Description: "High-performance Async shipment operations with FastAPI, Redis, and PostgreSQL. Live Swagger Docs available.",
					Tags:        []string{"REDIS", "ASYNC", "POSTGRES"},
					Link:        "https://fastapi-shipmentservice.onrender.com/scalar#introduction",
				},
				{
					Title:       "SERVERLESS AI CHAT",
					Description: "Real-time streaming chat platform using Azure Functions and Google Gemini 2.0 for content moderation.",
					Tags:        []string{"AZURE", "GEMINI", "NOSQL"},
					Link:        "https://github.com/mrYamusa/nextwork-streaming-backend",
				},
				{
					Title:       "ULTRASOUND AI CLASSIFIER",
					Description: "CNN model for classifying medical images using Transfer Learning (EfficientNet). Hosted on HuggingFace.",
					Tags:        []string{"TENSORFLOW", "FASTAPI", "CV"},
					Link:        "https://huggingface.co/mryamusa",
				},
			},
			GeneratedAt: time.Now().Format("2006-01-02 15:04:05 MST"),
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	// Health Check for Render
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "operational", "system": "yamusa-tech"}`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("System online at port %s", port)
	http.ListenAndServe(":"+port, nil)
}
