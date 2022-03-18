package domain

import "gopkg.in/guregu/null.v4"

const (
	Door       = "Door"
	Window     = "Window"
	QtyTin0_5L = 0.5
	QtyTin2_5L = 2.5
	QtyTin3_6L = 3.6
	QtyTin18L  = 18.0
)

type Wall struct {
	Height     null.Float `json:"height"`
	Width      null.Float `json:"width"`
	QtyDoors   null.Int   `json:"qty_doors"`
	QtyWindows null.Int   `json:"qty_windows"`
}

type WallItems struct {
	ID           null.Int    `db:"id" json:"id"`
	UID          null.String `db:"uid" json:"uid"`
	Name         null.String `db:"name" json:"name"`
	Height       null.Float  `db:"height" json:"height"`
	Width        null.Float  `db:"width" json:"width"`
	SquareMeters null.Float  `db:"square_meters" json:"square_meters"`
}

type Can struct {
	QtyCan0_5l null.Int   `json:"qty_can_0_5l"`
	QtyCan2_5l null.Int   `json:"qty_can_2_5l"`
	QtyCan3_6l null.Int   `json:"qty_can_3_6l"`
	QtyCan18l  null.Int   `json:"qty_can_18l"`
	Liters     null.Float `json:"liters"`
}
