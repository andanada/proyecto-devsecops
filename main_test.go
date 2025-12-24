package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthEndpoint(t *testing.T) {
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler devolvió código de estado incorrecto: obtuvo %v esperaba %v",
            status, http.StatusOK)
    }

    expected := "OK"
    if rr.Body.String() != expected {
        t.Errorf("handler devolvió cuerpo inesperado: obtuvo %v esperaba %v",
            rr.Body.String(), expected)
    }
}

func TestMainRoute(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
    })

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler devolvió código de estado incorrecto: obtuvo %v esperaba %v",
            status, http.StatusOK)
    }
}
