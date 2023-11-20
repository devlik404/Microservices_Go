package services1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/form/v4"
)

// User contains user information
type User struct {
	Name   string
	Age    uint8
	Gender string
}

// Auth is a simple handler
type Auth struct {
	l *log.Logger
}

// Auth is a simple handler
type Form struct {
	e *form.Decoder
}

// NewHAuth creates a new Auth handler with the given logger
func NewAuth(l *log.Logger) *Auth {
	return &Auth{l}
}

// ServeHTTP implements the go http.Handler interface

func (h *Auth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Auth request")

	// write the response
	fmt.Fprintf(rw, "Auth")
}

func NewForm(e *form.Decoder) *Form {
	return &Form{e}
}

func (f *Form) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Membuat objek decoder hanya jika dibutuhkan
	decoder := form.NewDecoder()

	// Mengambil nilai dari inputan formulir dengan nama "name"
	name := r.FormValue("name")

	// Mengambil nilai dari inputan formulir dengan nama "age"
	age := r.FormValue("age")

	// Mengambil nilai dari inputan formulir dengan nama "gender"
	gender := r.FormValue("gender")

	// Melakukan sesuatu dengan nilai-nilai yang telah diambil
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Gender:", gender)

	user := User{}

	// Menguraikan data formulir ke dalam struktur User
	err := decoder.Decode(&user, r.Form)
	if err != nil {
		log.Panic(err)
	}

	// Sekarang 'user' berisi data yang diuraikan dari formulir

	// Tulis respons ke ResponseWriter jika diperlukan

	// Menanggapi dengan data yang diuraikan
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(user)
}
