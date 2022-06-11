package models

type Company struct {
	Bedrijf string
}

type MaandagInput struct {
	Bedrijf string `form:"bedrijf"`
	Quota   uint   `form:"quota"`
}
