package controller

import (
	"encoding/json"
	"fmt"
	"mvcmodel/model"
	"mvcmodel/views"
	"net/http"
)

func crud() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some Error Creating the DB record"))
				return

			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")

			//data, err := model.ReadAll()
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte("error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete { // DELETE localhost:3000/Angad5

			//var name string
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			fmt.Println(name)
			json.NewEncoder(w).Encode("Row Deleted")

		}
	}
}
