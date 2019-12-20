package users

import (
	"context"
	"database/sql"
	"github.com/nik/JWTDemo/model"
	repsitory "github.com/nik/JWTDemo/repository"
	"github.com/nik/go-mysql-crud/models"
)

//NewSQLUsersRepo implement login repository interface
func NewLoginPostGresRepo(Conn *sql.DB) repsitory.LoginRepo {
	return &loginPostGresRepo{
		Conn: Conn,
	}
}

type loginPostGresRepo struct {
	Conn *sql.DB
}

func (m *loginPostGresRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Login, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*model.Login, 0)
	for rows.Next() {
		data := new(model.Login)

		err := rows.Scan(
			&data.UserID,
			&data.Email,
			&data.Password,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *loginPostGresRepo) Fetch(ctx context.Context, num int64) ([]*model.Login, error) {
	query := "Select userid, email, password From users limit ?"

	return m.fetch(ctx, query, num)
}

func (m *loginPostGresRepo) GetByUserID(ctx context.Context, id int64) (*model.Login, error) {
	query := "Select userid, email, password From users where userid=$1"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &model.Login{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

//Creates a new login session and returns token
func (m *loginPostGresRepo) Create(ctx context.Context, p *model.Login) (*model.Login, error) {
	query := "insert into users (email, password) values($1, $2) RETURNING userid;"
	err := m.Conn.QueryRow(query, p.Email, p.Password).Scan(&p.UserID)

	if err != nil {
		return nil, err
	}

	return p,nil
}

func (m *loginPostGresRepo) Update(ctx context.Context, p *model.Login) (*model.Login, error) {
	query := "Update users set password=?, email=? where userid=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Password,
		p.Email,
		p.UserID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *loginPostGresRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From users Where userid=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

//retrieves the user from the data store by using email as the lookup key
func (m *loginPostGresRepo) GetUserByEmail(email string) (*model.Login, error) {
	query := "Select userid, email, password From users where email=$1"

	rows, err := m.Conn.Query(query, email)
	if err != nil {
		return nil, err
	}

	//instantiate a login instance and populate the attributes of it with the values fetched
	payload := &model.Login{}
	defer rows.Close()
	for rows.Next() {
		//in case there are more than one record then raise the panic as email is supposed to be the unique per record
		err = rows.Scan(&payload.UserID, &payload.Email,&payload.Password)
		if(err!=nil) {
			panic(err)
		} else if (rows.Next()==true) {
			panic("Violation of unique record")
		}
	}

	return payload, nil
}