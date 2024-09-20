package DAO

import (
	"database/sql"
	"fmt"
	vo "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type InterfazAudioDAOImpl struct {
	db *sql.DB
}

// NewInterfazAudioDAO creates a new instance of InterfazAudioDAOImpl
func NewInterfazAudioDAO(db *sql.DB) *InterfazAudioDAOImpl {
	return &InterfazAudioDAOImpl{db: db}
}

// GetInterfazAudio retrieves an InterfazAudioVO by its ID
func (dao *InterfazAudioDAOImpl) GetInterfazAudio(id int) (*vo.InterfazAudioVO, error) {
	query := "SELECT id, nombre_corto, modelo, nombre_comercial, precio, frecuencia_id FROM interfaz_audio WHERE id = ?"
	row := dao.db.QueryRow(query, id)

	var interfaz vo.InterfazAudioVO
	var frecuenciaID int
	err := row.Scan(&interfaz.ID, &interfaz.NombreCorto, &interfaz.Modelo, &interfaz.NombreComercial, &interfaz.Precio, &frecuenciaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("interfaz not found")
		}
		return nil, fmt.Errorf("error fetching interfaz_audio by ID: %v", err)
	}

	// Fetch the FrecuenciaVO related to the InterfazAudioVO
	frecuencia, err := dao.GetFrecuencia(frecuenciaID)
	if err != nil {
		return nil, fmt.Errorf("error fetching frecuencia for interfaz: %v", err)
	}
	interfaz.Frecuencia = *frecuencia

	return &interfaz, nil
}

// CreateInterfazAudio inserts a new InterfazAudioVO into the database
func (dao *InterfazAudioDAOImpl) CreateInterfazAudio(interfaz *vo.InterfazAudioVO) error {
	query := "INSERT INTO interfaz_audio (nombre_corto, modelo, nombre_comercial, precio, frecuencia_id) VALUES (?, ?, ?, ?, ?)"
	_, err := dao.db.Exec(query, interfaz.NombreCorto, interfaz.Modelo, interfaz.NombreComercial, interfaz.Precio, interfaz.Frecuencia.ID)
	if err != nil {
		return fmt.Errorf("error creating interfaz_audio: %v", err)
	}
	return nil
}

// UpdateInterfazAudio updates an existing InterfazAudioVO in the database
func (dao *InterfazAudioDAOImpl) UpdateInterfazAudio(interfaz *vo.InterfazAudioVO) error {
	query := "UPDATE interfaz_audio SET nombre_corto = ?, modelo = ?, nombre_comercial = ?, precio = ?, frecuencia_id = ? WHERE id = ?"
	_, err := dao.db.Exec(query, interfaz.NombreCorto, interfaz.Modelo, interfaz.NombreComercial, interfaz.Precio, interfaz.Frecuencia.ID, interfaz.ID)
	if err != nil {
		return fmt.Errorf("error updating interfaz_audio: %v", err)
	}
	return nil
}

// DeleteInterfazAudio removes an InterfazAudioVO from the database by its ID
func (dao *InterfazAudioDAOImpl) DeleteInterfazAudio(id int) error {
	query := "DELETE FROM interfaz_audio WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting interfaz_audio: %v", err)
	}
	return nil
}

// GetFrecuencia retrieves the FrecuenciaVO related to a given interfaz ID
func (dao *InterfazAudioDAOImpl) GetFrecuencia(id int) (*vo.FrecuenciaVO, error) {
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

// SetFrecuencia assigns a FrecuenciaVO to a given InterfazAudioVO
func (dao *InterfazAudioDAOImpl) SetFrecuencia(interfazID int, frecuencia *vo.FrecuenciaVO) error {
	query := "UPDATE interfaz_audio SET frecuencia_id = ? WHERE id = ?"
	_, err := dao.db.Exec(query, frecuencia.ID, interfazID)
	if err != nil {
		return fmt.Errorf("error setting frecuencia for interfaz_audio: %v", err)
	}
	return nil
}

// GetAll retrieves all InterfazAudioVOs from the database
func (dao *InterfazAudioDAOImpl) GetAll() ([]*vo.InterfazAudioVO, error) {
	query := "SELECT id, nombre_corto, modelo, nombre_comercial, precio, frecuencia_id FROM interfaz_audio"
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all interfaz_audio: %v", err)
	}
	defer rows.Close()

	var interfaces []*vo.InterfazAudioVO
	for rows.Next() {
		var interfaz vo.InterfazAudioVO
		var frecuenciaID int
		err := rows.Scan(&interfaz.ID, &interfaz.NombreCorto, &interfaz.Modelo, &interfaz.NombreComercial, &interfaz.Precio, &frecuenciaID)
		if err != nil {
			return nil, fmt.Errorf("error scanning interfaz_audio: %v", err)
		}

		// Fetch the FrecuenciaVO related to the InterfazAudioVO
		frecuencia, err := dao.GetFrecuencia(frecuenciaID)
		if err != nil {
			return nil, fmt.Errorf("error fetching frecuencia for interfaz_audio: %v", err)
		}
		interfaz.Frecuencia = *frecuencia

		interfaces = append(interfaces, &interfaz)
	}

	return interfaces, nil
}
