package DAO

import (
	"database/sql"
	"fmt"
	vo "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type TipoDAOImpl struct {
	db *sql.DB
}

// NewTipoDAO creates a new instance of TipoDAOImpl
func NewTipoDAO(db *sql.DB) *TipoDAOImpl {
	return &TipoDAOImpl{db: db}
}

// Create inserts a new TipoVO into the database
func (dao *TipoDAOImpl) Create(t *vo.TipoVO) error {
	query := "INSERT INTO tipo (nombre, descripcion) VALUES (?, ?)"
	_, err := dao.db.Exec(query, t.Nombre, t.Descripcion)
	if err != nil {
		return fmt.Errorf("error creating tipo: %v", err)
	}
	return nil
}

// GetByID retrieves a TipoVO by its ID
func (dao *TipoDAOImpl) GetTipoByID(id int) (*vo.TipoVO, error) {
	query := "SELECT id, nombre, descripcion FROM tipo WHERE id = ?"
	row := dao.db.QueryRow(query, id)

	var tipo vo.TipoVO
	err := row.Scan(&tipo.ID, &tipo.Nombre, &tipo.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tipo not found")
		}
		return nil, fmt.Errorf("error fetching tipo by ID: %v", err)
	}
	return &tipo, nil
}

// Update modifies an existing TipoVO in the database
func (dao *TipoDAOImpl) Update(t *vo.TipoVO) error {
	query := "UPDATE tipo SET nombre = ?, descripcion = ? WHERE id = ?"
	_, err := dao.db.Exec(query, t.Nombre, t.Descripcion, t.ID)
	if err != nil {
		return fmt.Errorf("error updating tipo: %v", err)
	}
	return nil
}

// Delete removes a TipoVO from the database by its ID
func (dao *TipoDAOImpl) Delete(id int) error {
	query := "DELETE FROM tipo WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting tipo: %v", err)
	}
	return nil
}

// GetAll retrieves all TipoVOs from the database
func (dao *TipoDAOImpl) GetAll() ([]*vo.TipoVO, error) {
	query := "SELECT id, nombre, descripcion FROM tipo"
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all tipos: %v", err)
	}
	defer rows.Close()

	var tipos []*vo.TipoVO
	for rows.Next() {
		var tipo vo.TipoVO
		err := rows.Scan(&tipo.ID, &tipo.Nombre, &tipo.Descripcion)
		if err != nil {
			return nil, fmt.Errorf("error scanning tipo: %v", err)
		}
		tipos = append(tipos, &tipo)
	}

	return tipos, nil
}
