package main

import "time"

type Customer struct {
	ID           uint      `gorm:"primarykey" json:"cID"`
	RegisterDate time.Time `gorm:"autoCreateTime" json:"cRegisterDate"`
	Name         string    `json:"cName" schema:"cName"`
	Tel          int64     `json:"cTel"`
	Address      string    `json:"cAddress"`
}
