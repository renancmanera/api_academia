package domain

import "gorm.io/gorm"

type Treino struct {
	gorm.Model
	Nome       string      `json:"nome" gorm:"not null"`
	Descricao  string      `json:"descricao"`
	UsuarioID  uint        `json:"usuario_id"`
	Exercicios []Exercicio `gorm:"many2many:treino_exercicios;"`
}
