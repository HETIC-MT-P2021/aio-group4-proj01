package models

import (
	"github.com/HETIC-MT-P2021/aio-group4-proj01/Back/lib/common"
	"github.com/jinzhu/gorm"
)

type Category struct {
	Name   string
	gorm.Model

}

func (p Category) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"name":       p.Name,
		"date_created": p.CreatedAt,
	}
}