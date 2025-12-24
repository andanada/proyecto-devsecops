package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Servir archivos estáticos desde la carpeta "static"
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Ruta principal que muestra la imagen
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        html := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>DevSecOps Pipeline - Actividad 4</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    text-align: center;
                    background-color: #f0f0f0;
                    padding: 50px;
                }
                img {
                    max-width: 600px;
                    border: 3px solid #333;
                    border-radius: 10px;
                }
                h1 {
                    color: #333;
                }
            </style>
        </head>
        <body>
            <h1>🚀 Aplicación DevSecOps - Pipeline CI/CD Seguro</h1>
            <p>Esta aplicación ha sido desplegada mediante un pipeline automatizado</p>
            <img src="/static/logo.png" alt="Logo DevSecOps">
            <p><strong>Estudiante:</strong> Actividad 4 Individual</p>
        </body>
        </html>
        `
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprint(w, html)
    })

    // Ruta de health check
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "OK")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Servidor iniciado en http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
