package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Logger personalizado
type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.logWithLevel("INFO", message, fields)
}

func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.logWithLevel("ERROR", message, fields)
}

func (l *Logger) logWithLevel(level, message string, fields map[string]interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	logMsg := fmt.Sprintf(`{"timestamp":"%s","level":"%s","message":"%s"`, timestamp, level, message)
	
	for key, value := range fields {
		logMsg += fmt.Sprintf(`,"%s":"%v"`, key, value)
	}
	logMsg += "}"
	
	l.Println(logMsg)
}

var logger = NewLogger()

func main() {
	http.HandleFunc("/", handleMainRoute)
	http.HandleFunc("/health", handleHealthCheck)
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	logger.Info("Server starting", map[string]interface{}{
		"port": "8080",
		"env":  "production",
	})
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error("Server failed to start", map[string]interface{}{
			"error": err.Error(),
		})
		os.Exit(1)
	}
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	logger.Info("Request received", map[string]interface{}{
		"method": r.Method,
		"path":   r.URL.Path,
		"ip":     r.RemoteAddr,
	})
	
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Proyecto DevSecOps</title>
	</head>
	<body>
		<h1>Proyecto DevSecOps - Actividad 4</h1>
		<img src="/static/logo.png" alt="Go Logo">
		<p>Pipeline CI/CD funcionando correctamente</p>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Info("Health check", map[string]interface{}{
		"status": "healthy",
	})
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
