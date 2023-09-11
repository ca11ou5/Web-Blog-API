package handler

import (
	"blog/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	article := &models.Article{}

	err = json.Unmarshal(body, &article)
	if err != nil {
		log.Fatal(err)
	}

	err = h.service.CreateArticle(article)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(204)
}

func (h *Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	article, err := h.service.GetArticleByID(id)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(article)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *Handler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(articles)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *Handler) DeleteArticleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = h.service.DeleteArticleByID(id)
}
