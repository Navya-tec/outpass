package models

import (
	"fmt"
	"goprojects/outpass/db"
	"log"
	"time"
)

type Request struct {
	Id           int64      `json:"id"`
	ProceedingTo string     `json:"proceedingTo"`
	Dov          *time.Time `json:"date_of_visit"`
	TimeToLeave  *time.Time `json:"time_to_leave"`
	Conveyence   string     `json:"conveyence"`
	Status       string     `json:"status"`
	UserId       int        `json:"user_id"`
}

var requests = []Request{}

func (r *Request) SaveRequest() error {

	status := "pending"

	query := `INSERT INTO requests(proceedingTo,date_of_visit,time_to_leave,conveyence,status,user_id) VALUES (?,?,?,?,?,?)`
	res, err := db.DB.Exec(query, r.ProceedingTo, r.Dov, r.TimeToLeave, r.Conveyence, status, r.UserId)
	if err != nil {
		return err
		// panic("error in inserting request")
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic("Error in getting last inserted id for request")
	}

	r.Id = id
	return err
}

func GetAllRequest() ([]Request, error) {
	query := `SELECT id,proceedingTo,date_of_visit,time_to_leave,conveyence,status,user_id FROM requests`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error in fetching all requests from DB: %v", err)
	}

	defer rows.Close()

	var requests []Request
	for rows.Next() {
		var req Request
		err := rows.Scan(&req.Id, &req.ProceedingTo, &req.Dov, &req.TimeToLeave, &req.Conveyence, &req.Status, &req.UserId)
		if err != nil {
			panic("Error Scanning rows")
		}
		requests = append(requests, req)
	}

	return requests, err
}

func GetReqByUserId(user_id int) ([]Request, error) {
	query := `SELECT * FROM requests WHERE user_id=?`
	rows, err := db.DB.Query(query, user_id)
	if err != nil {
		log.Println("Error in Get all requests by user_id", err)
		return nil, err
	}

	var reqs []Request
	for rows.Next() {
		var req Request
		err := rows.Scan(&req.Id, &req.ProceedingTo, &req.Dov, &req.TimeToLeave, &req.Conveyence, &req.Status, &req.UserId)
		if err != nil {
			log.Println("Error Scanning Rows!", err)
			return reqs, err
		}

		reqs = append(reqs, req)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user rows: %v", err)
	}

	return reqs, nil
}

func UpdateRequestStatus(rId int, status string) error {

	query := `UPDATE requests SET status=? WHERE id=?`
	res, err := db.DB.Exec(query, status, rId)
	if err != nil {
		log.Println("Error in Updating Request Status", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no request found with id %d", rId)
	}

	return nil
}

func GetRequestByStatus(status string)([]Request,error){
	query := `SELECT * FROM requests WHERE status=?`
	rows, err := db.DB.Query(query,status)
	if err != nil {
		log.Println("Error in Get all requests by status", err)
		return nil, err
	}

	var reqs []Request
	for rows.Next() {
		var req Request
		err := rows.Scan(&req.Id, &req.ProceedingTo, &req.Dov, &req.TimeToLeave, &req.Conveyence, &req.Status, &req.UserId)
		if err != nil {
			log.Println("Error Scanning Rows!", err)
			return reqs, err
		}

		reqs = append(reqs, req)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user rows: %v", err)
	}

	return reqs, nil
}
