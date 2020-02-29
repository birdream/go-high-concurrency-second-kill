package repositories

import (
	"database/sql"
	"errors"
	"product-manager/common"
	"product-manager/datamodels"
	"strconv"
)

type IUserRepository interface {
	Conn() error
	Select(userName string) (user *datamodels.User, err error)
	Insert(user *datamodels.User) (userId int64, err error)
}

func NewUserResposity(table string, db *sql.DB) IUserRepository {
	return &UserManagerRepository{table, db}
}

type UserManagerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func (u *UserManagerRepository) Conn() error {
	if u.mysqlConn == nil {
		mysql, errMysql := common.NewMysqlConn()
		if errMysql != nil {
			return errMysql
		}
		u.mysqlConn = mysql
	}
	if u.table == "" {
		u.table = "user"
	}
	return nil
}

func (u *UserManagerRepository) Select(userName string) (user *datamodels.User, err error) {
	if userName == "" {
		return &datamodels.User{}, errors.New("username can not be empty")
	}

	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}

	sql := "Select * from " + u.table + " where username=?"
	rows, errRows := u.mysqlConn.Query(sql, userName)
	defer rows.Close()
	if errRows != nil {
		return &datamodels.User{}, errRows
	}

	result := common.GetResultRow(rows)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("user not exist")
	}

	user = &datamodels.User{}
	common.DataToStructByTagSQL(result, user)

	return
}

func (u *UserManagerRepository) Insert(user *datamodels.User) (userId int64, err error) {
	if err = u.Conn(); err != nil {
		return
	}

	sql := "INSERT " + u.table + " SET nickName=?,userName=?,passWord=?"
	stmt, errStmt := u.mysqlConn.Prepare(sql)
	if errStmt != nil {
		return userId, errStmt
	}
	result, errResult := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if errResult != nil {
		return userId, errResult
	}
	return result.LastInsertId()
}

func (u *UserManagerRepository) SelectByID(userId int64) (user *datamodels.User, err error) {
	if err = u.Conn(); err != nil {
		return &datamodels.User{}, err
	}
	sql := "select * from " + u.table + " where ID=" + strconv.FormatInt(userId, 10)
	row, errRow := u.mysqlConn.Query(sql)
	if errRow != nil {
		return &datamodels.User{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.User{}, errors.New("user not exist!")
	}
	user = &datamodels.User{}
	common.DataToStructByTagSQL(result, user)
	return
}
