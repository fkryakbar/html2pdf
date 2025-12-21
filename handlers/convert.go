package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"html2pdf/services"
)

// ConvertRequest represents the request body for HTML to PDF conversion
type ConvertRequest struct {
	HTML string                   `json:"html"`
	Var  []map[string]interface{} `json:"var,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ConvertHTMLToPDF handles the HTML to PDF conversion
func ConvertHTMLToPDF(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "bad_request",
			Message: "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	// Validate HTML content
	if strings.TrimSpace(req.HTML) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "bad_request",
			Message: "HTML content is required.",
		})
		return
	}

	// Process variable substitution
	html := processVariables(req.HTML, req.Var)

	// Generate PDF
	pdfBytes, err := services.GeneratePDF(html)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:   "pdf_generation_failed",
			Message: "Failed to generate PDF: " + err.Error(),
		})
		return
	}

	// Return PDF file
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"document.pdf\"")
	w.Header().Set("Content-Length", strconv.Itoa(len(pdfBytes)))
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}

// processVariables replaces <<variable>> placeholders with values from var array
func processVariables(html string, vars []map[string]interface{}) string {
	if len(vars) == 0 {
		return html
	}

	result := html

	// Iterate through each variable object in the array
	for _, varMap := range vars {
		for key, value := range varMap {
			// Convert value to string
			var strValue string
			switch v := value.(type) {
			case string:
				strValue = v
			case float64:
				// Format number without trailing zeros
				strValue = fmt.Sprintf("%v", v)
			default:
				// For other types, use JSON encoding
				if jsonBytes, err := json.Marshal(v); err == nil {
					strValue = string(jsonBytes)
				}
			}

			// Replace <<key>> with value
			placeholder := "<<" + key + ">>"
			result = strings.ReplaceAll(result, placeholder, strValue)
		}
	}

	return result
}
