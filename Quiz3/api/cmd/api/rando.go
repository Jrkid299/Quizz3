package main

import (
	"crypto/rand"
	"net/http"

	"quiz2.jalenlamb.net/internals/data"
)

func generateRandomString(length int64) string {

	randomStringSource := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}

func (app *application) showRandonString(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	rString := generateRandomString(id)

	// //create a new instance of the data struct containing the id we extracted from our url and some sample data

	data := data.RandoString{
		Data: rString,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
