package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"ngc2-webserver/entity"

	"github.com/julienschmidt/httprouter"
)

// read
func (h *MarvelHandler) HandleViewReports(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	w.Header().Set("Content-Type","application/json")
	ctx := context.Background()
	query := `SELECT id,hero_id,villain_id,description,incident_time
	FROM criminal_reports;`

	rows,err := h.Handler.QueryContext(ctx,query)
	if err != nil {
		InternalError(w,err)
	}

	var criminal_reports []entity.Criminal_report

	for rows.Next() {
		var criminal_report entity.Criminal_report

		err := rows.Scan(&criminal_report.ID,&criminal_report.Hero_id,&criminal_report.Villain_id,&criminal_report.Description,&criminal_report.Incident_time)
		if err != nil {
			InternalError(w,err)
		}

		criminal_reports = append(criminal_reports, criminal_report)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(criminal_reports)
	if err != nil {
		log.Fatal(err)
	}
}

// create
func (h *MarvelHandler) HandleCreateReports(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	w.Header().Set("Content-Type","application/json")
	var newReports []entity.Criminal_report
	
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newReports)
	if err != nil {
		log.Fatal(err)
	}

	for _,report := range newReports{
		ctx := context.Background()
		query := `INSERT INTO criminal_reports(hero_id,villain_id,description,incident_time)
					VALUES
						(?,?,?,?);`

		_,err := h.Handler.ExecContext(ctx,query,report.Hero_id,report.Villain_id,report.Description,report.Incident_time)
		if err != nil {
			InternalError(w,err)
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(HandleMessage{Message: "Success create reports"})
}

// Edit
func (h *MarvelHandler) HandleEditReport(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	w.Header().Set("Content-Type","application/json")
	paramsId := p.ByName("id")

	ctx := context.Background()
	query := `SELECT *
	FROM criminal_reports WHERE id = ?;`

	rows,err := h.Handler.QueryContext(ctx,query,paramsId)
	if err != nil {
		InternalError(w,err)
	}

	if !rows.Next(){
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(HandleMessage{Message: "Id not found"})
		return
	}

	var newReport entity.Criminal_report
	
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newReport)
	if err != nil {
		log.Fatal(err)
	}

	query2 := `UPDATE criminal_reports
	SET hero_id = ?, villain_id = ?, description = ?, incident_time = ?
	WHERE id = ?;`

	_,err = h.Handler.ExecContext(ctx,query2,newReport.Hero_id,newReport.Villain_id,newReport.Description,newReport.Incident_time,paramsId)
	if err != nil {
		InternalError(w,err)
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HandleMessage{Message: "Success update report"})
}

// Delete
func (h *MarvelHandler) HandleDeleteReport(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	w.Header().Set("Content-Type","application/json")
	paramsId := p.ByName("id")

	ctx := context.Background()

	query1 := `SELECT *
	FROM criminal_reports WHERE id = ?;`

	rows,err := h.Handler.QueryContext(ctx,query1,paramsId)
	if err != nil {
		InternalError(w,err)
	}

	if !rows.Next(){
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(HandleMessage{Message: "Id not found"})
		return
	}

	query2 := `DELETE FROM criminal_reports WHERE id = ?;`

	_,err = h.Handler.QueryContext(ctx,query2,paramsId)
	if err != nil {
		InternalError(w,err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HandleMessage{Message: "Success delete report"})
}