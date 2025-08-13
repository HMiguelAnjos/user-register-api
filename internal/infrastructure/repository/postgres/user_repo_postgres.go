package postgres

import (
	"database/sql"
	"userregisterapi/internal/app/ports"
	"userregisterapi/internal/common"
	"userregisterapi/internal/domain"
)

var _ ports.UserRepository = (*UserRepoPostgres)(nil)

type UserRepoPostgres struct {
	db *sql.DB
}

func NewUserRepoPostgres(db *sql.DB) *UserRepoPostgres {
	return &UserRepoPostgres{db: db}
}

func (r *UserRepoPostgres) Save(User *domain.User) error {
	_, err := r.db.Exec(`INSERT INTO Users (id, title, description, done, created_at, updated_at)
                         VALUES ($1, $2, $3, $4, $5, $6)`,
		User.ID, User.Title, User.Description, User.Done, User.CreatedAt, User.UpdatedAt)
	return err
}

func (r *UserRepoPostgres) GetByID(id string) (*domain.User, error) {
	row := r.db.QueryRow(`SELECT id, title, description, done, created_at, updated_at FROM Users WHERE id=$1`, id)
	t := &domain.User{}
	if err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CreatedAt, &t.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrNotFound
		}
		return nil, err
	}
	return t, nil
}

func (r *UserRepoPostgres) List() ([]*domain.User, error) {
	rows, err := r.db.Query(`SELECT id, title, description, done, created_at, updated_at FROM Users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Users []*domain.User
	for rows.Next() {
		t := &domain.User{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Done, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		Users = append(Users, t)
	}
	return Users, nil
}

func (r *UserRepoPostgres) Update(User *domain.User) error {
	res, err := r.db.Exec(`UPDATE Users SET title=$1, description=$2, done=$3, updated_at=$4 WHERE id=$5`,
		User.Title, User.Description, User.Done, User.UpdatedAt, User.ID)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return common.ErrNotFound
	}
	return nil
}

func (r *UserRepoPostgres) Delete(id string) error {
	res, err := r.db.Exec(`DELETE FROM Users WHERE id=$1`, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return common.ErrNotFound
	}
	return nil
}
