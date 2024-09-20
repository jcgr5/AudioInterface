package DAO

import (
	"database/sql"
	"fmt"
	vo "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type DispositivoDAOImpl struct {
	db *sql.DB
}

// NewDispositivoDAO creates a new instance of DispositivoDAOImpl
func NewDispositivoDAO(db *sql.DB) *DispositivoDAOImpl {
	return &DispositivoDAOImpl{db: db}
}

// Create inserts a new device (Dispositivo) into the database
func (dao *DispositivoDAOImpl) Create(d *vo.DispositivoVO) error {
	query := "INSERT INTO dispositivos (nombre, descripcion) VALUES (?, ?)"
	_, err := dao.db.Exec(query, d.Nombre, d.Descripcion)
	if err != nil {
		return fmt.Errorf("error creating dispositivo: %v", err)
	}
	return nil
}

// GetByID fetches a device (Dispositivo) by its ID
func (dao *DispositivoDAOImpl) GetByID(id int) (*vo.DispositivoVO, error) {
	query := "SELECT id, nombre, descripcion FROM dispositivos WHERE id = ?"
	row := dao.db.QueryRow(query, id)

	var dispositivo vo.DispositivoVO
	err := row.Scan(&dispositivo.ID, &dispositivo.Nombre, &dispositivo.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dispositivo not found")
		}
		return nil, fmt.Errorf("error fetching dispositivo by ID: %v", err)
	}

	return &dispositivo, nil
}

// Update modifies an existing device (Dispositivo)
func (dao *DispositivoDAOImpl) Update(d *vo.DispositivoVO) error {
	query := "UPDATE dispositivos SET nombre = ?, descripcion = ? WHERE id = ?"
	_, err := dao.db.Exec(query, d.Nombre, d.Descripcion, d.ID)
	if err != nil {
		return fmt.Errorf("error updating dispositivo: %v", err)
	}
	return nil
}

// Delete removes a device (Dispositivo) by its ID
func (dao *DispositivoDAOImpl) Delete(id int) error {
	query := "DELETE FROM dispositivos WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting dispositivo: %v", err)
	}
	return nil
}

// GetAll retrieves all devices (Dispositivos) from the database
func (dao *DispositivoDAOImpl) GetAll() ([]*vo.DispositivoVO, error) {
	query := "SELECT id, nombre, descripcion FROM dispositivos"
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all dispositivos: %v", err)
	}
	defer rows.Close()

	var dispositivos []*vo.DispositivoVO
	for rows.Next() {
		var dispositivo vo.DispositivoVO
		err := rows.Scan(&dispositivo.ID, &dispositivo.Nombre, &dispositivo.Descripcion)
		if err != nil {
			return nil, fmt.Errorf("error scanning dispositivo: %v", err)
		}
		dispositivos = append(dispositivos, &dispositivo)
	}

	return dispositivos, nil
}
