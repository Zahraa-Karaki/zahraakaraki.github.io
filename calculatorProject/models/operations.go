package models

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

type Operation struct {
	Id           string `json:"id"`
	FirstNumber  string `json:"firstnumber"`
	SecondNumber string `json:"secondnumber"`
	Operator     string `json:"operator"`
	Answer       string `json:"answer"`
}

func AddOperation(db *sql.DB, firstnumber string, secondnumber string, operator string) (string, error) {
	sqlStatement := "INSERT INTO records(first_number,second_number,operator,answer) values ($1,$2,$3,$4)"

	fn, err1 := strconv.ParseFloat(firstnumber, 64)
	sn, err2 := strconv.ParseFloat(secondnumber, 64)

	answer := ""

	if err1 != nil {
		answer = "Only numbers are permitted!"
		return answer, err1
	} else if err2 != nil {
		answer = "Only numbers are permitted!"
		return answer, err2
	}

	switch operator {
	case "+":
		answer = fmt.Sprintf("%v", fn+sn)
		break
	case "-":
		answer = fmt.Sprintf("%v", fn-sn)
		break
	case "*":
		answer = fmt.Sprintf("%v", fn*sn)
		break
	case "/":
		if sn == 0 {
			answer = "you cannot divide by zero!"
		} else {
			answer = fmt.Sprintf("%v", fn/sn)
		}
		break
	default:
		answer = "Invalid Operator!"
	}

	answer1, err1 := strconv.ParseFloat(answer, 64)

	if err1 == nil {

		res, err := db.Query(sqlStatement, firstnumber, secondnumber, operator, answer1)
		if err != nil {

			fmt.Println(err)

		} else {

			fmt.Println("added to database...", res)
		}

		return answer, err

	} else {

		return answer, err1

	}

}
