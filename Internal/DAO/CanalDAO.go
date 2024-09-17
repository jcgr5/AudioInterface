package DAO

import (
	"database/sql"
	"errors"
	"github.com/jcgr5/AudioInterface/Internal/VO"
)

type CanalDAO interface {
	GetCanal(id int) (*VO.CanalVO, error)
	CreateCanal(canal *VO.CanalVO) error
	UpdateCanal(canal *VO.CanalVO) error
	DeleteCanal(id int) error
	GetFuente(canalID int) (*VO.FuenteVO, error)
	SetFuente(canalID int, fuente *VO.FuenteVO) error
	GetAll() ([]*VO.CanalVO, error)
}

type canalDAO struct {
	db *sql.DB
}

func NewCanalDAO(db *sql.DB) *canalDAO {
	return &canalDAO{db: db}
}

func (dao *canalDAO) GetCanal(id int) (*VO.CanalVO, error) {
	query := "SELECT id, codigoCanal, etiqueta, volumen, solo, mute FROM canal WHERE id = ?"
	row := dao.db.QueryRow(query, id)

	var canal VO.CanalVO
	err := row.Scan(&canal.ID, &canal.CodigoCanal, &canal.Etiqueta, &canal.Volumen, &canal.Solo, &canal.Mute)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("canal no encontrado")
		}
		return nil, err
	}
	return &canal, nil
}

func (dao *canalDAO) CreateCanal(canal *VO.CanalVO) error {
	query := "INSERT INTO canal (codigoCanal, etiqueta, volumen, solo, mute) VALUES (?, ?, ?, ?, ?)"
	_, err := dao.db.Exec(query, canal.CodigoCanal, canal.Etiqueta, canal.Volumen, canal.Solo, canal.Mute)
	if err != nil {
		return err
	}
	return nil
}

func (dao *canalDAO) UpdateCanal(canal *VO.CanalVO) error {
	query := "UPDATE canal SET codigoCanal = ?, etiqueta = ?, volumen = ?, solo = ?, mute = ? WHERE id = ?"
	_, err := dao.db.Exec(query, canal.CodigoCanal, canal.Etiqueta, canal.Volumen, canal.Solo, canal.Mute, canal.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *canalDAO) DeleteCanal(id int) error {
	query := "DELETE FROM canal WHERE id = ?"
	_, err := dao.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dao *canalDAO) GetFuente(canalID int) (*VO.FuenteVO, error) {
	query := "SELECT fuente_id, Tipo FROM fuente WHERE canal_id = ?"
	row := dao.db.QueryRow(query, canalID)

	var fuente VO.FuenteVO
	err := row.Scan(&fuente.ID, &fuente.Tipo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("fuente no encontrada")
		}
		return nil, err
	}
	return &fuente, nil
}

func (dao *canalDAO) SetFuente(canalID int, fuente *VO.FuenteVO) error {
	query := "UPDATE fuente SET Tipo = ? WHERE canal_id = ? AND fuente_id = ?"
	_, err := dao.db.Exec(query, fuente.Tipo, canalID, fuente.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *canalDAO) GetAll() ([]*VO.CanalVO, error) {
	query := "SELECT id, codigoCanal, etiqueta, volumen, solo, mute FROM canal"
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var canales []*VO.CanalVO
	for rows.Next() {
		var canal VO.CanalVO
		err := rows.Scan(&canal.ID, &canal.CodigoCanal, &canal.Etiqueta, &canal.Volumen, &canal.Solo, &canal.Mute)
		if err != nil {
			return nil, err
		}
		canales = append(canales, &canal)
	}
	return canales, nil
}
