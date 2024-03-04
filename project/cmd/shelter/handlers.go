package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/snowiiiish/Golang/project/pkg/shelter/model"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createAnimalHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID             string `json:"id"`
		Kind_Of_Animal string `json:"kind_of_animal"`
		Kind_Of_Breed  string `json:"kind_of_breed"`
		Name           string `json:"name"`
		Age            string `json:"age"`
		Description    string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	animal := &model.Animal{
		ID:             input.ID,
		Kind_Of_Animal: input.Kind_Of_Animal,
		Kind_Of_Breed:  input.Kind_Of_Breed,
		Name:           input.Name,
		Age:            input.Age,
		Description:    input.Description,
	}

	err = app.models.Animals.Insert(animal)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, animal)
}

func (app *application) getAnimalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["animalID"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid animal ID")
		return
	}

	animal, err := app.models.Animals.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, animal)
}

/*
	 func (app *application) updateAnimalHandler(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		param := vars["animalID"] // CHECK HERE FOR ERRORS

		id, err := strconv.Atoi(param)
		if err != nil || id < 1 {
			app.respondWithError(w, http.StatusBadRequest, "Invalid animal ID")
			return
		}

		animal, err := app.models.Animals.Get(id)
		if err != nil {
			app.respondWithError(w, http.StatusNotFound, "404 Not Found")
			return
		}

		var input struct {
			ID             *string `json:"id"`
			Kind_Of_Animal *string `json:"kind_of_animal"`
			Kind_Of_Breed  *string `json:"kind_of_breed"`
			Name           *string `json:"name"`
			Age            *string `json:"age"`
			Description    *string `json:"description"`
		}

		err = app.readJSON(w, r, &input)
		if err != nil {
			app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if input.Name != nil {
			animal.Name = *input.Name
		}

		if input.Kind_Of_Animal != nil {
			animal.Kind_Of_Animal = *input.Kind_Of_Animal
		}

		if input.Kind_Of_Breed != nil {
			animal.Kind_Of_Breed = *input.Kind_Of_Breed
		}

		if input.ID != nil {
			animal.ID = *input.ID
		}

		if input.Description != nil {
			animal.Description = *input.Description
		}

		if input.Age != nil {
			animal.Age = *input.Age
		}
	}
*/
func (app *application) deleteAnimalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["menuId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	err = app.models.Animals.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}