package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PiiData struct {
	gorm.Model
	UserName   string `json:"user_name"`
	Approver   string `json:"approver"`
	POD        string `gorm:"default:null" json:"pod"`
	Status     string `json:"status"`
	Recipients string `json:"recipients"`
	TechFamily string `json:"tech_family"`
	Data       string `json:"data"`
	Consent    string   `json:"consent"`
	Q1         string   `json:"q1"`
	Q2         string   `json:"q2"`
	Q3         string   `json:"q3"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&PiiData{})
	return db
}