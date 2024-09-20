package DAO

import (
	"database/sql"
	"fmt"
	vo "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type UsuarioDAOImpl struct {
	db *sql.DB
}

// NewUsuarioDAO creates a new instance of UsuarioDAOImpl
func NewUsuarioDAO(db *sql.DB) *UsuarioDAOImpl {
	return &UsuarioDAOImpl{db: db}
}

// GetUsuario retrieves a UsuarioVO by its ID
func (dao *UsuarioDAOImpl) GetUsuario(id int) (*vo.UsuarioVO, error) {
	query := "SELECT id, email, password FROM usuario WHERE id = ?"
	row := dao.db.QueryRow(query, id)

	var usuario vo.UsuarioVO
	err := row.Scan(&usuario.ID, &usuario.Email, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario not found")
		}
		return nil, fmt.Errorf("error fetching usuario by ID: %v", err)
	}
	return &usuario, nil
}

// CreateUsuario inserts a new UsuarioVO into the database
func (dao *UsuarioDAOImpl) CreateUsuario(usuario *vo.UsuarioVO) error {
	query := "INSERT INTO usuario (email, password) VALUES (?, ?)"
	_, err := dao.db.Exec(query, usuario.Email, usuario.Password)
	if err != nil {
		return fmt.Errorf("error creating usuario: %v", err)
	}
	return nil
}

// UpdateUsuario updates an existing UsuarioVO in the database
func (dao *UsuarioDAOImpl) UpdateUsuario(usuario *vo.UsuarioVO) error {
	query := "UPDATE usuario SET email = ?, password = ? WHERE id = ?"
	_, err := dao.db.Exec(query, usuario.Email, usuario.Password, usuario.ID)
	if err != nil {
		return fmt.Errorf("error updating usuario: %v", err)
	}
	return nil
}

// DeleteUsuario deletes a UsuarioVO from the database by its ID
func (dao *UsuarioDAOImpl) DeleteUsuario(id int) error {
	query := "DELETE FROM usuario WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting usuario: %v", err)
	}
	return nil
}

// GetUsuarioByEmail retrieves a UsuarioVO by its email
func (dao *UsuarioDAOImpl) GetUsuarioByEmail(email string) (*vo.UsuarioVO, error) {
	query := "SELECT id, email, password FROM usuario WHERE email = ?"
	row := dao.db.QueryRow(query, email)

	var usuario vo.UsuarioVO
	err := row.Scan(&usuario.ID, &usuario.Email, &usuario.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario not found")
		}
		return nil, fmt.Errorf("error fetching usuario by email: %v", err)
	}
	return &usuario, nil
}

// GetAll retrieves all UsuarioVOs from the database
func (dao *UsuarioDAOImpl) GetAll() ([]*vo.UsuarioVO, error) {
	query := "SELECT id, email, password FROM usuario"
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all usuarios: %v", err)
	}
	defer rows.Close()

	var usuarios []*vo.UsuarioVO
	for rows.Next() {
		var usuario vo.UsuarioVO
		err := rows.Scan(&usuario.ID, &usuario.Email, &usuario.Password)
		if err != nil {
			return nil, fmt.Errorf("error scanning usuario: %v", err)
		}
		usuarios = append(usuarios, &usuario)
	}

	return usuarios, nil
}
