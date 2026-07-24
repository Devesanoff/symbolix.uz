package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Devesanoff/symbolix.uz/internal/usecase"
)

type ConvertRequest struct {
	Text string `json:"text"`
	Mode string `json:"mode"`
}

type ConvertResponse struct {
	Result string `json:"result"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}

func CyrillicToLatinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	mode := req.Mode
	if mode == "" {
		mode = "standard"
	}
	res := usecase.CyrillicToLatin(req.Text, mode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConvertResponse{Result: res})
}

func LatinToCyrillicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	res := usecase.LatinToCyrillic(req.Text)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConvertResponse{Result: res})
}

func PDFToDocxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(5 << 20) // 5MB max memory
	if err != nil {
		http.Error(w, "Unable to parse form or file too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	translateStr := r.URL.Query().Get("translate")
	translate, _ := strconv.ParseBool(translateStr)

	docxBytes, err := usecase.PDFToDocx(file, header.Size, translate)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting file: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Set("Content-Disposition", `attachment; filename="converted.docx"`)
	w.WriteHeader(http.StatusOK)
	w.Write(docxBytes)
}
