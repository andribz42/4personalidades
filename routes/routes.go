package routes

import (
	"estudo/4personalidades-go/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.CriarUmaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.AlterarUmaPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.ExcluirUmaPersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
