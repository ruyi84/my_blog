package models

import (
	"my_blog/db/mysql"
	"testing"
)

func Test_Creat(t *testing.T) {
	db := mysql.GetDB()

	table := Users{}

	db.CreateTable(&table)
}
