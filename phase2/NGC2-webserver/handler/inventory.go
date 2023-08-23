package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"ngc2-webserver/entity"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// NGC 3
// task 1
func (h MarvelHandler)HandleInventories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	query := `SELECT id,name,codeitem,stock,description,status
	FROM inventories;`
	rows,err := h.Handler.QueryContext(ctx,query)
	if err != nil {
		log.Fatal(err)
	}

	var inventories []entity.Inventory
	for rows.Next(){
		var inventory entity.Inventory

		err := rows.Scan(&inventory.ID,&inventory.Name,&inventory.CodeItem,&inventory.Stock,&inventory.Description,&inventory.Status)
		if err != nil {
			log.Fatal(err)
		}
		inventories = append(inventories, inventory)
	}

	// json
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(inventories)

	if err != nil {
		log.Fatal(err)
	}
}

// task 2
func (h MarvelHandler)HandleInventoryById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idInventoryStr := p.ByName("id")
	idInventory,err := strconv.Atoi(idInventoryStr)
	if err != nil {
		log.Fatal(err)
	}


	ctx := context.Background()
	query := `
	SELECT id,name,codeitem,stock,description,status
	FROM inventories
	WHERE id = ?;`
	rows,err := h.Handler.QueryContext(ctx,query,idInventory)
	if err != nil {
		log.Fatal(err)
	}

	var inventory entity.Inventory
	if rows.Next(){

		err := rows.Scan(&inventory.ID,&inventory.Name,&inventory.CodeItem,&inventory.Stock,&inventory.Description,&inventory.Status)
		if err != nil {
			log.Fatal(err)
		}
	}

	// json
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(inventory)

	if err != nil {
		log.Fatal(err)
	}
}

// task 3 create
func (h MarvelHandler)HandleCreateInventories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newInventories []entity.Inventory
	
	err := decoder.Decode(&newInventories)
	if err != nil {
		log.Fatal(err)
	}


	ctx := context.Background()
	query := `
	INSERT INTO Inventories(Name,CodeItem,Stock,Description,Status)
	VALUES
	(?,?,?,?,?);`

	for _,inventory := range newInventories{
		_,err = h.Handler.ExecContext(ctx,query,inventory.Name,inventory.CodeItem,inventory.Stock,inventory.Description,inventory.Status)
	if err != nil {
		log.Fatal(err)
	}
	}
	

	// json
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(newInventories)

	if err != nil {
		log.Fatal(err)
	}
}

// task 4 PUT /inventories/:id


func (h MarvelHandler)HandleEditInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idInventoryStr := p.ByName("id")
	idInventory,err := strconv.Atoi(idInventoryStr)
	if err != nil {
		log.Fatal(err)
	}


	ctx := context.Background()
	query1 := `
	SELECT id,name,codeitem,stock,description,status
	FROM inventories
	WHERE id = ?;`
	rows,err := h.Handler.QueryContext(ctx,query1,idInventory)
	if err != nil {
		log.Fatal(err)
	}

	var oldInventory entity.Inventory
	if rows.Next(){
		err := rows.Scan(&oldInventory.ID,&oldInventory.Name,&oldInventory.CodeItem,&oldInventory.Stock,&oldInventory.Description,&oldInventory.Status)
		if err != nil {
			log.Fatal(err)
		}
	}else{
		log.Fatal("id not found")
	}

	// replace
	decoder := json.NewDecoder(r.Body)
	var newInventory entity.Inventory
	
	err = decoder.Decode(&newInventory)
	if err != nil {
		log.Fatal(err)
	}

	query2 := `
	UPDATE inventories
	SET name = ?, codeitem = ?, stock = ?, description = ?, status = ?
	WHERE id = ?;`

	
	_,err = h.Handler.ExecContext(ctx,query2,newInventory.Name,newInventory.CodeItem,newInventory.Stock,newInventory.Description,newInventory.Status,idInventory)
	if err != nil {
		log.Fatal(err)
	}

	// json
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(map[string]entity.Inventory{
		"before" : oldInventory,
		"after"	 : newInventory,})

	if err != nil {
		log.Fatal(err)
	}
}

// task 5 DELETE /inventories/:id


func (h MarvelHandler)HandleDeleteInventory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idInventoryStr := p.ByName("id")
	idInventory,err := strconv.Atoi(idInventoryStr)
	if err != nil {
		log.Fatal(err)
	}


	ctx := context.Background()
	query1 := `
	DELETE FROM inventories WHERE id = ?;
	`
	_,err = h.Handler.ExecContext(ctx,query1,idInventory)
	if err != nil {
		log.Fatal(err)
	}

	// json
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(map[string]int{
		"deleted" : idInventory})

	if err != nil {
		log.Fatal(err)
	}
}