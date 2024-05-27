package model

import "mvcmodel/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * from TODO")
	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * from TODO WHERE name=?", name) // localhost:3000/?name=Angad10
	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
