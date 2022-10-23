package main

import (
	"net/http"

	"quiz2.jalenlamb.net/internals/data"
)

func (app *application) showMyInfo(w http.ResponseWriter, r *http.Request) {

	data := data.Me{
		Name:       "Jalen",
		Age:        21,
		Location:   "Belmopan, Belize",
		Email:      "2018118881@ub.edu.bz",
		Interest:   []string{"Music", "Piano", "Bass", "Drums", "Computer Builds"},
		Occupation: "Website designer for People Diabetes Association-Belize",
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
