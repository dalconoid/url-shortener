package server

import (
	"encoding/json"
	"fmt"
	"github.com/dalconoid/url-shortener/models"
	"github.com/dalconoid/url-shortener/storage"
	"github.com/dalconoid/url-shortener/urls"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func handleAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleRegisterShortenedURL(db storage.IURLStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		dec := json.NewDecoder(r.Body)
		registerURLRequest := &models.RegisterURLRequest{}
		if err = dec.Decode(registerURLRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		v := validator.New()
		errs := v.Struct(registerURLRequest)
		if errs != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Validation error(s):\n")
			for _, e := range errs.(validator.ValidationErrors) {
				w.Write([]byte(fmt.Sprintf("%v\n", e)))
			}
			return
		}

		url := strings.Trim(registerURLRequest.URL, "/ ")
		if !urls.ValidateURL(url){
			http.Error(w, "Invalid URL ", http.StatusBadRequest)
			return
		}

		slug := strings.Trim(registerURLRequest.Slug, " ")
		if slug == "" {
			slug = urls.GenerateSlug(7)
		} else {
			if !urls.ValidateSlug(slug) {
				http.Error(w, "Provided slug is not valid", http.StatusBadRequest)
			}
		}

		su, cErr := db.RegisterURL(url, slug)
		if cErr != nil {
			if cErr.Code == models.ErrDefaultCode {
				http.Error(w, cErr.Err.Error(), http.StatusInternalServerError)
				return
			}

			http.Error(w, cErr.Err.Error(), http.StatusBadRequest)
			return
		}

		resp := &models.RegisterURLResponse{Slug: su.Slug}
		data, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(data)
	}
}

func handleRedirectBySlug(db storage.IURLStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := mux.Vars(r)["slug"]
		redirectURL, cErr := db.GetRedirectURL(slug)
		if cErr != nil {
			if cErr.Code == models.ErrSlugNotRegistered {
				http.Error(w, cErr.Err.Error(), http.StatusNotFound)
				return
			} else {
				http.Error(w, cErr.Err.Error(), http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	}
}

