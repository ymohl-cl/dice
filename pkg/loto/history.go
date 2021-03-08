package loto

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/jszwec/csvutil"
)

// History data type
type History struct {
	Draws []Draw
}

// Draw description data
type Draw struct {
	ID        string            `csv:"annee_numero_de_tirage"`
	Date      string            `csv:"date_de_tirage"`
	B1        int32             `csv:"boule_1"`
	B2        int32             `csv:"boule_2"`
	B3        int32             `csv:"boule_3"`
	B4        int32             `csv:"boule_4"`
	B5        int32             `csv:"boule_5"`
	Joker     int32             `csv:"numero_chance"`
	Joker2    int32             `csv:"boule_complementaire"`
	NTirage   int32             `csv:"1er_ou_2eme_tirage"`
	OtherData map[string]string `csv:"-"`
}

// NewHistory read the file parameter to extract the draw history
func NewHistory(file string) ([]Draw, error) {
	var err error
	var Draws []Draw

	var f *os.File
	if f, err = os.Open(file); err != nil {
		return []Draw{}, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return []Draw{}, err
	}

	for {
		var d Draw
		if err := dec.Decode(&d); err == io.EOF {
			break
		} else if err != nil {
			return []Draw{}, err
		}
		if d.Joker == 0 {
			if d.NTirage == 1 {
				d.Joker = d.Joker2
			} else {
				continue
			}
		}
		Draws = append(Draws, d)
	}
	return Draws, nil
}
