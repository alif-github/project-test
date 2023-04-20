package dao

import (
	"database/sql"
	"fmt"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/repository"
	"net/http"
)

type authUserDAO struct {
	AbstractDAO
}

var AuthUserDAO = authUserDAO{}.initiate()

func (input authUserDAO) initiate() (output authUserDAO) {
	output.FileName = "AuthUserDAO.go"
	output.TableName = "auth_user"
	return
}

func (input authUserDAO) RegisterNewUser(db *sql.DB, userModel repository.AuthUserModel) (id int64, err model.ErrorModel) {
	var (
		funcName = "RegisterNewUser"
		query    string
	)

	query = fmt.Sprintf(`INSERT INTO %s 
		(full_name, username, password, created_by, updated_by) 
		VALUES 
		($1, $2, $3, $4, $5)
		RETURNING id`,
		input.TableName)

	param := []interface{}{
		userModel.FullName.String,
		userModel.Username.String,
		userModel.Password.String,
		userModel.CreatedBy.Int64,
		userModel.UpdatedBy.Int64,
	}

	result := db.QueryRow(query, param...)
	dbError := result.Scan(&id)
	if dbError != nil && dbError.Error() != sql.ErrNoRows.Error() {
		err = model.GenerateErrorModel(http.StatusInternalServerError, dbError.Error(), input.FileName, funcName)
		return
	}

	err = model.GenerateNonErrorModel()
	return
}

func (input authUserDAO) GetUserByUsername(db *sql.DB, userModel repository.AuthUserModel) (resultDB repository.AuthUserModel, err model.ErrorModel) {
	var (
		funcName = "GetUserByUsername"
		query    string
	)

	query = fmt.Sprintf(`SELECT 
		id, username, password 
		FROM %s 
		WHERE username = $1 AND deleted = FALSE `,
		input.TableName)

	param := []interface{}{userModel.Username.String}
	_, dbError := db.Prepare(query)
	if dbError != nil {
		err = model.GenerateErrorModel(http.StatusInternalServerError, dbError.Error(), input.FileName, funcName)
		return
	}

	result := db.QueryRow(query, param...)
	dbError = result.Scan(&resultDB.ID, &resultDB.Username, &resultDB.Password)
	if dbError != nil && dbError.Error() != sql.ErrNoRows.Error() {
		err = model.GenerateErrorModel(http.StatusInternalServerError, dbError.Error(), input.FileName, funcName)
		return
	}

	err = model.GenerateNonErrorModel()
	return
}
