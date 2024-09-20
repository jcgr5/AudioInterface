package DAO

import (
	"database/sql"
	"fmt"
	vo2 "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type EntradaDAOImpl struct {
	db *sql.DB
}

// NewEntradaDAO creates a new instance of EntradaDAOImpl
func NewEntradaDAO(db *sql.DB) *EntradaDAOImpl {
	return &EntradaDAOImpl{db: db}
}

// GetEntrada fetches an entry (Entrada) by its ID
func (dao *EntradaDAOImpl) GetEntrada(id int) (*vo2.EntradaVO, error) {
	query := `SELECT e.id, e.etiqueta, e.descripcion, d.id, d.nombre, d.descripcion 
			  FROM entradas e 
			  JOIN dispositivos d ON e.dispositivo_id = d.id 
			  WHERE e.id = ?`

	row := dao.db.QueryRow(query, id)

	var entrada vo2.EntradaVO
	err := row.Scan(&entrada.ID, &entrada.Etiqueta, &entrada.Descripcion, &entrada.Dispositivo.ID, &entrada.Dispositivo.Nombre, &entrada.Dispositivo.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("entrada not found")
		}
		return nil, fmt.Errorf("error fetching entrada by ID: %v", err)
	}

	return &entrada, nil
}

// CreateEntrada inserts a new entry (Entrada) into the database
func (dao *EntradaDAOImpl) CreateEntrada(entrada *vo2.EntradaVO) error {
	query := "INSERT INTO entradas (dispositivo_id, etiqueta, descripcion) VALUES (?, ?, ?)"
	_, err := dao.db.Exec(query, entrada.Dispositivo.ID, entrada.Etiqueta, entrada.Descripcion)
	if err != nil {
		return fmt.Errorf("error creating entrada: %v", err)
	}
	return nil
}

// UpdateEntrada modifies an existing entry (Entrada)
func (dao *EntradaDAOImpl) UpdateEntrada(entrada *vo2.EntradaVO) error {
	query := "UPDATE entradas SET dispositivo_id = ?, etiqueta = ?, descripcion = ? WHERE id = ?"
	_, err := dao.db.Exec(query, entrada.Dispositivo.ID, entrada.Etiqueta, entrada.Descripcion, entrada.ID)
	if err != nil {
		return fmt.Errorf("error updating entrada: %v", err)
	}
	return nil
}

// DeleteEntrada removes an entry (Entrada) by its ID
func (dao *EntradaDAOImpl) DeleteEntrada(id int) error {
	query := "DELETE FROM entradas WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting entrada: %v", err)
	}
	return nil
}

// GetDispositivo fetches the associated device (Dispositivo) for a specific entry (Entrada)
func (dao *EntradaDAOImpl) GetDispositivo(entradaID int) (*vo2.DispositivoVO, error) {
	query := `SELECT d.id, d.nombre, d.descripcion 
			  FROM dispositivos d 
			  JOIN entradas e ON e.dispositivo_id = d.id 
			  WHERE e.id = ?`

	row := dao.db.QueryRow(query, entradaID)

	var dispositivo vo2.DispositivoVO
	err := row.Scan(&dispositivo.ID, &dispositivo.Nombre, &dispositivo.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("dispositivo not found")
		}
		return nil, fmt.Errorf("error fetching dispositivo: %v", err)
	}

	return &dispositivo, nil
}

// SetDispositivo sets a new device (Dispositivo) for an entry (Entrada)
func (dao *EntradaDAOImpl) SetDispositivo(entradaID int, dispositivo *vo2.DispositivoVO) error {
	query := "UPDATE entradas SET dispositivo_id = ? WHERE id = ?"
	_, err := dao.db.Exec(query, dispositivo.ID, entradaID)
	if err != nil {
		return fmt.Errorf("error setting dispositivo: %v", err)
	}
	return nil
}

// GetAll retrieves all entries (Entradas) from the database
func (dao *EntradaDAOImpl) GetAll() ([]*vo2.EntradaVO, error) {
	query := `SELECT e.id, e.etiqueta, e.descripcion, d.id, d.nombre, d.descripcion 
			  FROM entradas e 
			  JOIN dispositivos d ON e.dispositivo_id = d.id`
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all entradas: %v", err)
	}
	defer rows.Close()

	var entradas []*vo2.EntradaVO
	for rows.Next() {
		var entrada vo2.EntradaVO
		err := rows.Scan(&entrada.ID, &entrada.Etiqueta, &entrada.Descripcion, &entrada.Dispositivo.ID, &entrada.Dispositivo.Nombre, &entrada.Dispositivo.Descripcion)
		if err != nil {
			return nil, fmt.Errorf("error scanning entrada: %v", err)
		}
		entradas = append(entradas, &entrada)
	}

	return entradas, nil
}

// GetAllByDispositivo retrieves all entries (Entradas) associated with a specific device (Dispositivo)
func (dao *EntradaDAOImpl) GetAllByDispositivo(dispositivoID int) ([]*vo2.EntradaVO, error) {
	query := `SELECT e.id, e.etiqueta, e.descripcion, d.id, d.nombre, d.descripcion 
			  FROM entradas e 
			  JOIN dispositivos d ON e.dispositivo_id = d.id 
			  WHERE d.id = ?`
	rows, err := dao.db.Query(query, dispositivoID)
	if err != nil {
		return nil, fmt.Errorf("error fetching entradas by dispositivo: %v", err)
	}
	defer rows.Close()

	var entradas []*vo2.EntradaVO
	for rows.Next() {
		var entrada vo2.EntradaVO
		err := rows.Scan(&entrada.ID, &entrada.Etiqueta, &entrada.Descripcion, &entrada.Dispositivo.ID, &entrada.Dispositivo.Nombre, &entrada.Dispositivo.Descripcion)
		if err != nil {
			return nil, fmt.Errorf("error scanning entrada: %v", err)
		}
		entradas = append(entradas, &entrada)
	}

	return entradas, nil
}
