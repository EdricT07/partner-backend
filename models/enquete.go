package models

type EnqueteMaandag struct {
	Bedrijf string
	Quota   uint
	Week    int
	Jaar    int
}

type EnqueteVrijdag struct {
	Bedrijf string
	Quota   uint
	Waarom  string
	Steden  string
	Week    int
	Jaar    int
}
