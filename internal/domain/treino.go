package domain

import "gorm.io/gorm"

type Treino struct {
	gorm.Model
	Nome      string `json:"nome" gorm:"not null"`
	Descricao string `json:"descricao"`
	UsuarioID uint   `json:"usuario_id"` // ID do aluno dono do treino
}
