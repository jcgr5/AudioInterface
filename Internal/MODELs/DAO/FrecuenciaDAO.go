package DAO

import (
	"database/sql"
	"fmt"
	vo "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type FrecuenciaDAOImpl struct {
	db *sql.DB
}

// NewFrecuenciaDAO creates a new instance of FrecuenciaDAOImpl
func NewFrecuenciaDAO(db *sql.DB) *FrecuenciaDAOImpl {
	return &FrecuenciaDAOImpl{db: db}
}

// GetFrecuencia fetches a FrecuenciaVO by its ID
func (dao *FrecuenciaDAOImpl) GetFrecuencia(id int) (*vo.FrecuenciaVO, error) {
	query := "SELECT id, valor FROM frecuencias WHERE id = ?"
	row := dao.db.QueryRow(query, id)

	var frecuencia vo.FrecuenciaVO
	err := row.Scan(&frecuencia.ID, &frecuencia.Valor)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("frecuencia not found")
		}
		return nil, fmt.Errorf("error fetching frecuencia by ID: %v", err)
	}

	return &frecuencia, nil
}

// CreateFrecuencia inserts a new FrecuenciaVO into the database
func (dao *FrecuenciaDAOImpl) CreateFrecuencia(frecuencia *vo.FrecuenciaVO) error {
	query := "INSERT INTO frecuencias (valor) VALUES (?)"
	_, err := dao.db.Exec(query, frecuencia.Valor)
	if err != nil {
		return fmt.Errorf("error creating frecuencia: %v", err)
	}
	return nil
}

// UpdateFrecuencia updates an existing FrecuenciaVO in the database
func (dao *FrecuenciaDAOImpl) UpdateFrecuencia(frecuencia *vo.FrecuenciaVO) error {
	query := "UPDATE frecuencias SET valor = ? WHERE id = ?"
	_, err := dao.db.Exec(query, frecuencia.Valor, frecuencia.ID)
	if err != nil {
		return fmt.Errorf("error updating frecuencia: %v", err)
	}
	return nil
}

// DeleteFrecuencia deletes a FrecuenciaVO by its ID
func (dao *FrecuenciaDAOImpl) DeleteFrecuencia(id int) error {
	query := "DELETE FROM frecuencias WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting frecuencia: %v", err)
	}
	return nil
}

// GetAll retrieves all FrecuenciaVO records from the database
func (dao *FrecuenciaDAOImpl) GetAll() ([]*vo.FrecuenciaVO, error) {
	query := "SELECT id, valor FROM frecuencias"
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all frecuencias: %v", err)
	}
	defer rows.Close()

	var frecuencias []*vo.FrecuenciaVO
	for rows.Next() {
		var frecuencia vo.FrecuenciaVO
		err := rows.Scan(&frecuencia.ID, &frecuencia.Valor)
		if err != nil {
			return nil, fmt.Errorf("error scanning frecuencia: %v", err)
		}
		frecuencias = append(frecuencias, &frecuencia)
	}

	return frecuencias, nil
}
