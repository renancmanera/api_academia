package domain

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome  string `json:"nome" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Senha string `json:"senha" gorm:"not null"`
	Papel string `json:"papel" gorm:"not null"` // aluno ou instrutor
}
