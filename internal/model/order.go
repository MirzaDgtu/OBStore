package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderUid      int            `gorm:"column:orderuid" json:"order_uid"`
	UnicumNum     int            `gorm:"column:unicumnum" json:"unicum_num"`
	FolioNum      int            `gorm:"column:foliosum" json:"folio_num"`
	OrderDate     string         `gorm:"column:orderdate" json:"order_date"`
	FolioDate     string         `gorm:"column:foliodate" json:"folio_date"`
	OrderSum      float64        `gorm:"column:ordersum" json:"order_sum"`
	FolioSum      float64        `gorm:"column:foliosum" json:"folio_sum"`
	Driver        string         `gorm:"column:driver" json:"driver"`
	Agent         string         `gorm:"column:agent" json:"agent"`
	Brieforg      string         `gorm:"column:brieforg" json:"brieforg"`
	ClientId      int            `gorm:"column:clientId" json:"client_Id"`
	ClientName    string         `gorm:"column:clientName" json:"client_name"`
	ClientAddress string         `gorm:"column:clientaddress" json:"client_address"`
	VidDoc        string         `gorm:"column:viddoc" json:"vid_doc"`
	UserId        int            `gorm:"column:userid" json:"user_id"`
	StartAt       *time.Time     `gorm:"column:startAt" json:"start_at"`
	FinishAt      *time.Time     `gorm:"column:finishAt" json:"finish_at"`
	Done          bool           `gorm:"column:done" json:"done"`
	Status        int            `gorm:"column:status" json:"status"`
	OrderDetails  []OrderDetails `gorm:"foreignKey:OrderUID" json:"order_details"`
}

func (Order) TableName() string {
	return "orders"
}
