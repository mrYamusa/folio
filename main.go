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

type ArchNode struct {
	Label       string
	Description string
}

type APIEndpoint struct {
	Method      string
	Path        string
	Description string
}

type TutorialStep struct {
	Title    string
	Code     string
	Note     string
	Response string
	Language string
}

type ResponseField struct {
	Field       string
	Type        string
	Description string
}

type CurrentlyBuilding struct {
	Name           string
	Tagline        string
	Description    string
	Status         string
	Architecture   []ArchNode
	Endpoints      []APIEndpoint
	TechStack      []string
	Tutorials      []TutorialStep
	ResponseFields []ResponseField
	RepoLink       string
	LiveLink       string
	ScalarLink     string
	DocsLink       string
}

// Data for Home Page
type PageData struct {
	Name        string
	Role        string
	Tagline     string
	Socials     []Link
	Stats       []Stat
	Projects    []Project
	Building    CurrentlyBuilding
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
			ResumeLink: "https://www.linkedin.com/in/david-idris-174222289/",
			Socials: []Link{
				{"GITHUB", "https://github.com/mrYamusa"},
				{"LINKEDIN", "https://www.linkedin.com/in/david-idris-174222289/"},
				{"TWITTER", "https://x.com/mr_yamusa"},
			},
			Stats: []Stat{
				{Value: "2+", Label: "YEARS EXP", Trend: "LEVEL UP"},
				{Value: "99.9%", Label: "UPTIME", Trend: "STABLE"},
				{Value: "5", Label: "ACTIVE DEPLOYS", Trend: "+3"},
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
			Building: CurrentlyBuilding{
				Name:        "ORPHEUS",
				Tagline:     "Voice-activated music discovery engine",
				Description: "Ingests songs from Spotify & YouTube, extracts audio features using deep learning (VibeNet + librosa), stores them as vectors in Qdrant, and serves semantic playlist search with Shazam-style audio recognition.",
				Status:      "IN DEVELOPMENT",
				Architecture: []ArchNode{
					{"Orpheus API", "Heroku — FastAPI handling search, ingestion, YouTube downloads, and Qdrant storage"},
					{"Orpheus Extractor", "HuggingFace Space (16 GB) — Heavy ML inference with VibeNet + librosa"},
					{"Qdrant Cloud", "Vector database storing 30-dim song embeddings and 32-dim frame embeddings"},
					{"Spotify API", "Song discovery and metadata enrichment across genres"},
					{"YouTube / yt-dlp", "Audio download pipeline with cookie auth and Deno cipher solving"},
				},
				Endpoints: []APIEndpoint{
					{"GET", "/health", "Liveness check"},
					{"GET", "/stats", "Qdrant collection statistics (song count, frame count)"},
					{"GET", "/scheduler/status", "Next scheduled run, song source, config"},
					{"POST", "/scheduler/trigger", "Manually fire an ingestion cycle"},
					{"POST", "/ingest/song", "Ingest a specific YouTube video by ID"},
					{"POST", "/search", "Semantic song search from structured JSON query"},
					{"POST", "/match/snippet", "Upload an audio clip for Shazam-style matching"},
					{"GET", "/scalar", "Interactive API docs (Scalar UI)"},
				},
				TechStack: []string{"FastAPI", "Qdrant", "librosa", "VibeNet", "yt-dlp", "Spotify API", "APScheduler", "Docker", "Heroku", "HuggingFace"},
				Tutorials: []TutorialStep{
					{
						Title:    "Health Check — Verify the API is alive",
						Language: "bash",
						Code:     "curl https://orpheus-api-0bc4904911a6.herokuapp.com/health",
						Note:     "Returns {\"status\": \"ok\"} when the service is up. Use this to verify connectivity before making search requests.",
						Response: "{\n  \"status\": \"ok\"\n}",
					},
					{
						Title:    "Semantic Search — Find songs by vibe",
						Language: "bash",
						Code:     "curl -X POST https://orpheus-api-0bc4904911a6.herokuapp.com/search \\\n  -H \"Content-Type: application/json\" \\\n  -d '{\n    \"energy\": 0.8,\n    \"valence\": 0.6,\n    \"danceability\": 0.7,\n    \"tempo_min\": 120,\n    \"tempo_max\": 150,\n    \"mode\": \"major\",\n    \"limit\": 5\n  }'",
						Note:     "All vibe fields (energy, valence, danceability, acousticness, instrumentalness, speechiness, liveness) are floats in [0, 1]. Omit any field to default to 0.5. tempo_min/tempo_max are BPM hard filters. mode accepts \"major\" or \"minor\".",
						Response: "{\n  \"count\": 2,\n  \"playlist\": [\n    {\n      \"rank\": 1,\n      \"score\": 0.9288,\n      \"title\": \"Calm Down\",\n      \"artist\": \"Rema\",\n      \"album\": \"Rave & Roses\",\n      \"genres\": [\"Afrobeats\"],\n      \"tempo_bpm\": 106.99,\n      \"key\": \"C# minor\",\n      \"mood\": {\n        \"energy\": 0.62,\n        \"valence\": 0.74,\n        \"danceability\": 0.80\n      },\n      \"links\": {\n        \"youtube\": \"https://youtube.com/watch?v=...\",\n        \"spotify\": \"https://open.spotify.com/track/...\"\n      }\n    }\n  ]\n}",
					},
					{
						Title:    "Match Snippet — Shazam-style recognition",
						Language: "bash",
						Code:     "curl -X POST \\\n  \"https://orpheus-api-0bc4904911a6.herokuapp.com/match/snippet?similar_limit=5\" \\\n  -F \"file=@recording.mp3\"",
						Note:     "Upload any audio clip (mp3, wav, m4a, ogg, flac — anything ffmpeg can decode). The API extracts 32-dim frame embeddings from 5-second windows with 2.5s hop and matches against the orpheus_frames collection in Qdrant. The similar_limit query param controls how many similar songs to return (default: 5).",
						Response: "{\n  \"match\": {\n    \"title\": \"Calm Down\",\n    \"artist\": \"Rema\",\n    \"youtube_id\": \"WcIcVapfqXw\",\n    \"heard_at\": {\n      \"start_s\": 45.0,\n      \"end_s\": 50.0\n    },\n    \"confidence\": 0.94,\n    \"links\": {\n      \"youtube\": \"https://youtube.com/watch?v=WcIcVapfqXw\",\n      \"spotify\": \"https://open.spotify.com/track/...\"\n    }\n  },\n  \"similar_songs\": [\n    {\n      \"title\": \"Essence\",\n      \"artist\": \"Wizkid\",\n      \"score\": 0.87,\n      \"links\": { \"youtube\": \"...\" }\n    }\n  ]\n}",
					},
					{
						Title:    "Library Stats — Check collection health",
						Language: "bash",
						Code:     "curl https://orpheus-api-0bc4904911a6.herokuapp.com/stats",
						Note:     "Returns Qdrant collection statistics including total songs indexed, frame vectors stored, and index health status. Use this to monitor library growth after ingestion cycles.",
						Response: "{\n  \"total_songs\": 142,\n  \"total_frames\": 9840,\n  \"status\": \"green\",\n  \"vectors_count\": 142,\n  \"indexed_vectors_count\": 142\n}",
					},
					{
						Title:    "Ingest Song — Add a YouTube video",
						Language: "bash",
						Code:     "curl -X POST https://orpheus-api-0bc4904911a6.herokuapp.com/ingest/song \\\n  -H \"Content-Type: application/json\" \\\n  -d '{\"youtube_id\": \"dQw4w9WgXcQ\"}'",
						Note:     "Pass an 11-character YouTube video ID. Orpheus will download the audio via yt-dlp, extract features through the HuggingFace Space extractor (VibeNet + librosa), and store the 30-dim song embedding + 32-dim frame embeddings in Qdrant. Skips if already in the library.",
						Response: "{\n  \"message\": \"Ingestion queued.\",\n  \"youtube_id\": \"dQw4w9WgXcQ\",\n  \"title\": \"Rick Astley - Never Gonna Give You Up\",\n  \"artist\": \"Rick Astley\"\n}",
					},
					{
						Title:    "Python Client — Search example",
						Language: "python",
						Code:     "import httpx\n\nBASE = \"https://orpheus-api-0bc4904911a6.herokuapp.com\"\n\n# Semantic search — find chill acoustic tracks\nresp = httpx.post(f\"{BASE}/search\", json={\n    \"energy\": 0.3,\n    \"acousticness\": 0.9,\n    \"valence\": 0.7,\n    \"limit\": 5\n})\n\nfor song in resp.json()[\"playlist\"]:\n    print(f\"{song['rank']}. {song['title']} \"\n          f\"by {song['artist']} \"\n          f\"(score: {song['score']:.2f})\")",
						Note:     "Use httpx (async-capable HTTP client) for Python integration. All mood fields are optional — omit any to default to neutral (0.5). The response includes full metadata: mood scores, instrument profile, spectral info, and streaming links.",
						Response: "1. Essence by Wizkid (score: 0.91)\n2. Calm Down by Rema (score: 0.88)\n3. Love Nwantiti by CKay (score: 0.85)",
					},
				},
				ResponseFields: []ResponseField{
					{"rank", "integer", "1-based position in result list"},
					{"score", "float", "Cosine similarity (0–1). Higher = better match"},
					{"title", "string", "Song title from metadata"},
					{"artist", "string", "Primary artist name"},
					{"mood", "object", "VibeNet scores: energy, valence, danceability, acousticness, instrumentalness, speechiness, liveness"},
					{"tempo_bpm", "float", "Detected tempo in BPM"},
					{"key", "string", "Musical key, e.g. 'C# minor'"},
					{"links", "object", "YouTube and Spotify URLs"},
					{"instrument_profile", "object", "Harmonic ratio, brightness, tonal strength, estimated instrument"},
					{"cover_url", "string", "Album art URL from Spotify"},
				},
				RepoLink:   "https://github.com/mrYamusa/Orpheus",
				LiveLink:   "https://orpheus-api-0bc4904911a6.herokuapp.com",
				ScalarLink: "https://orpheus-api-0bc4904911a6.herokuapp.com/scalar",
				DocsLink:   "https://orpheus-api-0bc4904911a6.herokuapp.com/docs",
			},
			GeneratedAt: time.Now().Format("2006-01-02 15:04:05 MST"),
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	// 2. RESUME PAGE HANDLER (/resume)
	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		data := ResumeData{
			Skills: []string{"Python", "Go", "FastAPI", "Django", "Docker", "Kubernetes", "Azure", "PostgreSQL", "Redis", "TensorFlow", "Git/Linux", "Ollama"},
			Experience: []ExperienceItem{
				{
					Title:    "Backend Developer",
					Company:  "NACOS Landmark University",
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
