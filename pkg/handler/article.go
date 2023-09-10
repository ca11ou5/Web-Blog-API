package handler

import (
	"blog/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	article := &models.Article{}

	err = json.Unmarshal(body, &article)
	if err != nil {
		log.Fatal(err)
	}

	err = h.service.CreateArticle(article)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	
}
