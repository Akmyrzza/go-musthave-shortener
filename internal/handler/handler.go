package handler

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"io"
	"net/http"
	"strings"
)

type Handler struct {
	Service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) HandleURL(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		h.createID(res, req)
	} else if req.Method == http.MethodGet {
		h.getURL(res, req)
	} else {
		http.Error(res, "wrong method", http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) createID(res http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "bad body", http.StatusBadRequest)
		return
	}

	id := h.Service.CreateID(string(reqBody))

	res.Header().Add("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("http://" + req.Host + req.RequestURI + id))
}

func (h *Handler) getURL(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.RequestURI, "/")

	originalURL, ok := h.Service.GetURL(id)
	if !ok {
		http.Error(res, "id not found", http.StatusBadRequest)
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
