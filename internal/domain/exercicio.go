package domain

import "gorm.io/gorm"

type Exercicio struct {
	gorm.Model
	Nome          string `json:"nome" gorm:"not null"`
	Descricao     string `json:"descricao"`
	GrupoMuscular string `json:"grupo_muscular"`
}
