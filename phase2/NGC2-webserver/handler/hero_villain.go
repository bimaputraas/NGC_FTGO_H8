package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"ngc2-webserver/entity"

	"github.com/julienschmidt/httprouter"
)

func (h MarvelHandler) HandleHeroes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	query := `SELECT id,name,universe,skill,imageURL
	FROM heroes;`
	rows, err := h.Handler.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	var heroes []entity.Hero
	for rows.Next() {
		var hero entity.Hero

		err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
		if err != nil {
			log.Fatal(err)
		}
		heroes = append(heroes, hero)
	}

	// json
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(heroes)

	if err != nil {
		log.Fatal(err)
	}
}

func (h MarvelHandler) HandleVillains(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	query := `SELECT id,name,universe,imageURL
	FROM villains;`
	rows, err := h.Handler.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	var villains []entity.Villain
	for rows.Next() {
		var villain entity.Villain

		err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
		if err != nil {
			log.Fatal(err)
		}
		villains = append(villains, villain)
	}

	// json
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(villains)

	if err != nil {
		log.Fatal(err)
	}
}