package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

type registrationHandler struct {
	db              sqlExecutor
	publisher       publisher
	regExchangeName string
}

func (h registrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.ping(w, r)
		return

	case "/api/v1/registration":
		h.handleRegistration(w, r)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (h registrationHandler) ping(w http.ResponseWriter, r *http.Request) {
	if _, err := h.db.Exec("SELECT 1"); err != nil {
		log.Printf("select 1 from db err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.publisher.Publish(h.regExchangeName, "", false, false, amqp.Publishing{
		ContentType: "plain/text",
		Body:        []byte("HealthCheck"),
	}); err != nil {
		log.Printf("publish helath check to reg exhange err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("OK"))
}

func (h registrationHandler) handleRegistration(w http.ResponseWriter, r *http.Request) {
	user := user{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("decode user err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := h.db.NamedQuery(`
		INSERT INTO users (first_name, email, age)
		VALUES (:first_name, :email, :age)
	`, user); err != nil {
		log.Printf("save user to db err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userData, _ := json.Marshal(user)
	if err := h.publisher.Publish(h.regExchangeName, "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        userData,
	}); err != nil {
		log.Printf("publish user data to reg exhange err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
