package DAO

import (
	"database/sql"
	"fmt"
	vo "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type FuenteDAOImpl struct {
	db *sql.DB
}

// NewFuenteDAO creates a new instance of FuenteDAOImpl
func NewFuenteDAO(db *sql.DB) *FuenteDAOImpl {
	return &FuenteDAOImpl{db: db}
}

// Create inserts a new FuenteVO into the database
func (dao *FuenteDAOImpl) Create(f *vo.FuenteVO) error {
	query := "INSERT INTO fuentes (tipo_id) VALUES (?)"
	_, err := dao.db.Exec(query, f.Tipo.ID)
	if err != nil {
		return fmt.Errorf("error creating fuente: %v", err)
	}
	return nil
}

// GetByID fetches a FuenteVO by its ID from the database
func (dao *FuenteDAOImpl) GetByID(id int) (*vo.FuenteVO, error) {
	query := "SELECT id, tipo_id FROM fuentes WHERE id = ?"
	row := dao.db.QueryRow(query, id)

	var fuente vo.FuenteVO
	var tipoID int
	err := row.Scan(&fuente.ID, &tipoID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("fuente not found")
		}
		return nil, fmt.Errorf("error fetching fuente by ID: %v", err)
	}
	tipoDaoTemp := &TipoDAOImpl{db: dao.db}
	// Fetch the TipoVO related to the FuenteVO (assuming you have a GetTipoByID function)
	tipo, err := tipoDaoTemp.GetTipoByID(tipoID) // This function should be implemented elsewhere
	if err != nil {
		return nil, fmt.Errorf("error fetching tipo for fuente: %v", err)
	}

	fuente.Tipo = *tipo
	return &fuente, nil
}

// Update modifies an existing FuenteVO in the database
func (dao *FuenteDAOImpl) Update(f *vo.FuenteVO) error {
	query := "UPDATE fuentes SET tipo_id = ? WHERE id = ?"
	_, err := dao.db.Exec(query, f.Tipo.ID, f.ID)
	if err != nil {
		return fmt.Errorf("error updating fuente: %v", err)
	}
	return nil
}

// Delete removes a FuenteVO from the database by its ID
func (dao *FuenteDAOImpl) Delete(id int) error {
	query := "DELETE FROM fuentes WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting fuente: %v", err)
	}
	return nil
}

// GetAllByTipo retrieves all FuenteVOs based on their tipoID
func (dao *FuenteDAOImpl) GetAllByTipo(tipoID int) ([]*vo.FuenteVO, error) {
	query := "SELECT id, tipo_id FROM fuentes WHERE tipo_id = ?"
	rows, err := dao.db.Query(query, tipoID)
	if err != nil {
		return nil, fmt.Errorf("error fetching fuentes by tipo: %v", err)
	}
	defer rows.Close()

	var fuentes []*vo.FuenteVO
	for rows.Next() {
		var fuente vo.FuenteVO
		err := rows.Scan(&fuente.ID, &fuente.Tipo.ID)
		if err != nil {
			return nil, fmt.Errorf("error scanning fuente: %v", err)
		}

		tipoDaoTemp := &TipoDAOImpl{db: dao.db}

		// Fetch the TipoVO related to the FuenteVO (assuming you have a GetTipoByID function)
		tipo, err := tipoDaoTemp.GetTipoByID(fuente.Tipo.ID) // This function should be implemented elsewhere
		if err != nil {
			return nil, fmt.Errorf("error fetching tipo for fuente: %v", err)
		}

		fuente.Tipo = *tipo
		fuentes = append(fuentes, &fuente)
	}

	return fuentes, nil
}
