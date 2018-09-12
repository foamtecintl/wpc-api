package repository

import (
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//User type
type User struct {
	ID          int    `json:"id"`
	CtrateDate  string `json:"createDate"`
	EmployeeID  string `json:"employeeId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Status      string `json:"ststus"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Department  string `json:"department"`
	PhoneNumber string `json:"phoneNO"`
	RoleName    string `json:"roleName"`
	Other       string `json:"other"`
}

//FindAppUserByUsername database
func FindAppUserByUsername(username string) (User, error) {
	db := getConnection()
	defer db.Close()
	var user User
	sqlQuery := `SELECT AU_ID, AU_CREATE_DATE, COALESCE(AU_EMPLOYEE_ID,''),
		COALESCE(AU_FIRST_NAME,''), COALESCE(AU_LAST_NAME,''), COALESCE(AU_STATUS,''), COALESCE(AU_USERNAME,''),
		COALESCE(AU_PASSWORD,''), COALESCE(AU_EMAIL,''), COALESCE(AU_DEPARTMENT,''), COALESCE(AU_PHONE,''), COALESCE(AU_LEVEL,'')
		FROM APP_USER
		WHERE AU_USERNAME = ?`
	rowData, err := db.Query(sqlQuery, username)
	if err != nil {
		return user, err
	}
	for rowData.Next() {
		err = rowData.Scan(
			&user.ID,
			&user.CtrateDate,
			&user.EmployeeID,
			&user.FirstName,
			&user.LastName,
			&user.Status,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Department,
			&user.PhoneNumber,
			&user.RoleName,
		)
		return user, err
	}
	rowData.Close()
	return user, nil
}

//FindAppUserByEmployeeID database
func FindAppUserByEmployeeID(employeeID string) (User, error) {
	db := getConnection()
	defer db.Close()
	var user User
	sqlQuery := `SELECT AU_ID, AU_CREATE_DATE, COALESCE(AU_EMPLOYEE_ID,''),
		COALESCE(AU_FIRST_NAME,''), COALESCE(AU_LAST_NAME,''), COALESCE(AU_STATUS,''), COALESCE(AU_USERNAME,''),
		COALESCE(AU_PASSWORD,''), COALESCE(AU_EMAIL,''), COALESCE(AU_DEPARTMENT,''), COALESCE(AU_PHONE,''), COALESCE(AU_LEVEL,'')
		FROM APP_USER
		WHERE AU_EMPLOYEE_ID = ?`
	rowData, err := db.Query(sqlQuery, employeeID)
	if err != nil {
		return user, err
	}
	for rowData.Next() {
		err = rowData.Scan(
			&user.ID,
			&user.CtrateDate,
			&user.EmployeeID,
			&user.FirstName,
			&user.LastName,
			&user.Status,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Department,
			&user.PhoneNumber,
			&user.RoleName,
		)
		return user, err
	}
	rowData.Close()
	return user, nil
}

//UserRegister is register
func UserRegister(body map[string]string) error {
	db := getConnection()
	defer db.Close()
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	sqlQuery := `INSERT INTO APP_USER (AU_CREATE_DATE, AU_EMPLOYEE_ID,AU_USERNAME, AU_EMAIL, AU_PASSWORD, AU_DEPARTMENT,AU_STATUS, AU_LEVEL)
	VALUES (?,?,?,?,?,?,?,?)`
	result, err := db.Exec(
		sqlQuery,
		now,
		body["employeeId"],
		body["username"],
		body["email"],
		hashedPassword,
		"Please select",
		"REGISTER",
		"USER_REGISTER",
	)
	if err != nil {
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	sqlQuery = `INSERT INTO USER_LOG (LOG_CREATE_DATE, LOG_STATUS, LOG_IP, LOG_AU)
	VALUES (?,?,?,?)`
	_, err = db.Exec(sqlQuery, now, "Register", body["clientIP"], int(lastID))
	if err != nil {
		return err
	}
	return nil
}

//UpdateUser database
func UpdateUser(body map[string]string) error {
	db := getConnection()
	defer db.Close()
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	id, _ := strconv.Atoi(body["id"])
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body["password"]), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	sqlQuery := `UPDATE APP_USER SET 
	AU_EMPLOYEE_ID = ?, AU_FIRST_NAME = ?, AU_LAST_NAME = ?, AU_EMAIL = ?, AU_PASSWORD = ?, AU_DEPARTMENT = ?, AU_PHONE = ?,
	AU_STATUS = ?, AU_LEVEL = ?
	WHERE AU_ID = ?`
	_, err = db.Exec(
		sqlQuery,
		body["employeeId"],
		body["firstName"],
		body["lastName"],
		body["email"],
		hashedPassword,
		body["department"],
		body["telephone"],
		"APPROVE",
		"USER",
		id,
	)
	if err != nil {
		return err
	}
	sqlQuery = `INSERT INTO USER_LOG (LOG_CREATE_DATE, LOG_STATUS, LOG_IP, LOG_AU)
	VALUES (?,?,?,?)`
	_, err = db.Exec(sqlQuery, now, "Update", body["clientIP"], id)
	if err != nil {
		return err
	}
	return nil
}
