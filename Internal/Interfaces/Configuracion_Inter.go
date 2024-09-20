package Interfaces

import (
	"database/sql"
	"fmt"
	vo2 "github.com/jcgr5/AudioInterface/Internal/MODELs/VO"
)

type ConfiguracionDAO interface {
	GetConfiguracion(id int) (*vo2.ConfiguracionVO, error)
	CreateConfiguracion(config *vo2.ConfiguracionVO) error
	UpdateConfiguracion(config *vo2.ConfiguracionVO) error
	DeleteConfiguracion(id int) error
	GetCanales(configID int) ([]vo2.CanalVO, error)
	GetEntradas(configID int) ([]vo2.EntradaVO, error)
	GetAll() ([]*vo2.ConfiguracionVO, error)
	GetAllByUser(usuarioID int) ([]*vo2.ConfiguracionVO, error)
}

type ConfiguracionDAOImpl struct {
	db *sql.DB
}

// NewConfiguracionDAO creates a new instance of ConfiguracionDAOImpl
func NewConfiguracionDAO(db *sql.DB) *ConfiguracionDAOImpl {
	return &ConfiguracionDAOImpl{db: db}
}

func (dao *ConfiguracionDAOImpl) GetConfiguracion(id int) (*vo2.ConfiguracionVO, error) {
	query := `
		SELECT c.id, c.fecha, u.id, u.nombre, i.id, i.nombre
		FROM configuraciones c
		JOIN usuarios u ON c.usuario_id = u.id
		JOIN interfaces_audio i ON c.interfaz_id = i.id
		WHERE c.id = ?`

	row := dao.db.QueryRow(query, id)

	var config vo2.ConfiguracionVO
	var usuario vo2.UsuarioVO
	var interfaz vo2.InterfazAudioVO

	err := row.Scan(&config.ID, &config.Fecha, &usuario.ID, &usuario.Email, &interfaz.ID, &interfaz.NombreCorto)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("configuration not found")
		}
		return nil, err
	}

	// Assign Usuario and Interfaz to ConfiguracionVO
	config.Usuario = usuario
	config.Interfaz = interfaz

	// Fetch related Canales and Entradas
	config.Canales, err = dao.GetCanales(config.ID)
	if err != nil {
		return nil, err
	}

	config.Entradas, err = dao.GetEntradas(config.ID)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// CreateConfiguracion inserts a new configuration into the database
func (dao *ConfiguracionDAOImpl) CreateConfiguracion(config *vo2.ConfiguracionVO) error {
	query := "INSERT INTO configuraciones (fecha, usuario_id, interfaz_id) VALUES (?, ?, ?)"
	_, err := dao.db.Exec(query, config.Fecha, config.Usuario.ID, config.Interfaz.ID)
	if err != nil {
		return err
	}
	// Insert related Canales and Entradas if necessary
	return nil
}

// UpdateConfiguracion updates an existing configuration
func (dao *ConfiguracionDAOImpl) UpdateConfiguracion(config *vo2.ConfiguracionVO) error {
	query := "UPDATE configuraciones SET fecha = ?, usuario_id = ?, interfaz_id = ? WHERE id = ?"
	_, err := dao.db.Exec(query, config.Fecha, config.Usuario.ID, config.Interfaz.ID, config.ID)
	if err != nil {
		return err
	}
	// Update related Canales and Entradas if necessary
	return nil
}

// DeleteConfiguracion deletes a configuration by its ID
func (dao *ConfiguracionDAOImpl) DeleteConfiguracion(id int) error {
	query := "DELETE FROM configuraciones WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetCanales fetches all channels related to a specific configuration
func (dao *ConfiguracionDAOImpl) GetCanales(configID int) ([]vo2.CanalVO, error) {
	query := "SELECT id, codigo_canal, etiqueta, volumen, solo, mute FROM canales WHERE configuracion_id = ?"
	rows, err := dao.db.Query(query, configID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var canales []vo2.CanalVO
	for rows.Next() {
		var canal vo2.CanalVO
		if err := rows.Scan(&canal.ID, &canal.CodigoCanal, &canal.Etiqueta, &canal.Volumen, &canal.Solo, &canal.Mute); err != nil {
			return nil, err
		}
		canales = append(canales, canal)
	}

	return canales, nil
}

// GetEntradas fetches all inputs related to a specific configuration
func (dao *ConfiguracionDAOImpl) GetEntradas(configID int) ([]vo2.EntradaVO, error) {
	query := `SELECT e.id, e.idDispositivo, e.Descripcion, e.Etiqueta  FROM Configuraciones
                 JOIN Entradas e ON Configuraciones.idEntrada = e.id`
	rows, err := dao.db.Query(query, configID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entradas []vo2.EntradaVO
	for rows.Next() {
		var entrada vo2.EntradaVO
		if err := rows.Scan(&entrada.ID, &entrada.Dispositivo.ID, &entrada.Descripcion, &entrada.Etiqueta); err != nil {
			return nil, err
		}
		entradas = append(entradas, entrada)
	}

	return entradas, nil
}

// GetAll fetches all configurations from the database
func (dao *ConfiguracionDAOImpl) GetAll() ([]*vo2.ConfiguracionVO, error) {
	query := `
		SELECT c.id, c.fecha, u.id, u.email, i.id, i.nombreCorto
		FROM configuraciones c
		JOIN usuarios u ON c.usuario_id = u.id
		JOIN interfaces_audio i ON c.interfaz_id = i.id`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []*vo2.ConfiguracionVO
	for rows.Next() {
		var config vo2.ConfiguracionVO
		var usuario vo2.UsuarioVO
		var interfaz vo2.InterfazAudioVO

		err := rows.Scan(&config.ID, &config.Fecha, &usuario.ID, &usuario.Email, &interfaz.ID, &interfaz.NombreCorto)
		if err != nil {
			return nil, err
		}

		config.Usuario = usuario
		config.Interfaz = interfaz

		config.Canales, err = dao.GetCanales(config.ID)
		if err != nil {
			return nil, err
		}

		config.Entradas, err = dao.GetEntradas(config.ID)
		if err != nil {
			return nil, err
		}

		configs = append(configs, &config)
	}

	return configs, nil
}

// GetAllByUser fetches all configurations for a specific user
func (dao *ConfiguracionDAOImpl) GetAllByUser(usuarioID int) ([]*vo2.ConfiguracionVO, error) {
	query := `
		SELECT c.id, c.fecha, u.id, u.email, i.id, i.nombreCorto
		FROM configuraciones c
		JOIN usuarios u ON c.usuario_id = u.id
		JOIN interfaces_audio i ON c.interfaz_id = i.id
		WHERE c.usuario_id = ?`

	rows, err := dao.db.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []*vo2.ConfiguracionVO
	for rows.Next() {
		var config vo2.ConfiguracionVO
		var usuario vo2.UsuarioVO
		var interfaz vo2.InterfazAudioVO

		err := rows.Scan(&config.ID, &config.Fecha, &usuario.ID, &usuario.Email, &interfaz.ID, &interfaz.NombreCorto)
		if err != nil {
			return nil, err
		}

		config.Usuario = usuario
		config.Interfaz = interfaz

		config.Canales, err = dao.GetCanales(config.ID)
		if err != nil {
			return nil, err
		}

		config.Entradas, err = dao.GetEntradas(config.ID)
		if err != nil {
			return nil, err
		}

		configs = append(configs, &config)
	}

	return configs, nil
}
