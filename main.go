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

// --- STRUCTS ---
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

// Data for Home Page
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

// Data for Resume Page
type ExperienceItem struct {
	Title    string
	Company  string
	Duration string
	Points   []string
}

type ResumeData struct {
	Skills     []string
	Experience []ExperienceItem
}

func main() {
	// Parse all templates
	tmpl, err := template.ParseFS(content, "templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	// 1. HOME PAGE HANDLER
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		data := PageData{
			Name:       "DAVID YAMUSA IDRIS",
			Role:       "BACKEND & AI ENGINEER",
			Tagline:    "Architecting autonomous systems and scalable infrastructure that hunt latency and capture value.",
			ResumeLink: "https://www.linkedin.com/in/david-idris-174222289/", // DIRECT LINKEDIN LINK
			Socials: []Link{
				{"GITHUB", "https://github.com/mrYamusa"},
				{"LINKEDIN", "https://www.linkedin.com/in/david-idris-174222289/"},
				{"TWITTER", "https://x.com/mr_yamusa"},
			},
			Stats: []Stat{
				{Value: "4+", Label: "YEARS EXP", Trend: "LEVEL UP"},
				{Value: "99.9%", Label: "UPTIME", Trend: "STABLE"},
				{Value: "8", Label: "ACTIVE DEPLOYS", Trend: "+3"},
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

	// 2. RESUME PAGE HANDLER (/resume)
	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		data := ResumeData{
			Skills: []string{"Python", "Go", "FastAPI", "Django", "Docker", "Kubernetes", "Azure", "PostgreSQL", "Redis", "TensorFlow", "Git/Linux"},
			Experience: []ExperienceItem{
				{
					Title:    "Backend Developer",
					Company:  "NACS Landmark University",
					Duration: "10/2024 - 09/2025",
					Points: []string{
						"Architected secure REST API backend for a voting system using Django/DRF.",
						"Designed user authentication, MySQL integration, and cloud storage for media.",
						"Deployed and managed the application on Render with CI/CD integration.",
					},
				},
				{
					Title:    "Software Engineer Intern",
					Company:  "NCAIR Nigeria",
					Duration: "03/2024 - 09/2024",
					Points: []string{
						"Developed and maintained Python-based applications in a Linux environment.",
						"Collaborated with the development team to enhance AI model integration.",
						"Utilized Git for version control and CI/CD pipelines.",
					},
				},
			},
		}
		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, "resume.html", data)
	})

	// Health Check
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
