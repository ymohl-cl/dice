package loto

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"sort"

	"github.com/jszwec/csvutil"
)

var (
	recentFiles = []string{
		"../../pkg/loto/resources/loto_201911.csv",
		"../../pkg/loto/resources/loto_201902.csv",
		"../../pkg/loto/resources/loto2017.csv",
		"../../pkg/loto/resources/nouveau_loto.csv",
	}
	oldFiles = []string{
		"../../pkg/loto/resources/loto.csv",
	}
)

// History data type
type History struct {
	OldDraws    []OldDraw
	RecentDraws []Draw
}

// OldDraw description data
type OldDraw struct {
	ID        string            `csv:"annee_numero_de_tirage"`
	Date      string            `csv:"date_de_tirage"`
	Day       string            `csv:"jour_de_tirage"`
	B1        int32             `csv:"boule_1"`
	B2        int32             `csv:"boule_2"`
	B3        int32             `csv:"boule_3"`
	B4        int32             `csv:"boule_4"`
	B5        int32             `csv:"boule_5"`
	B6        int32             `csv:"boule_6"`
	Joker     int32             `csv:"boule_complementaire"`
	Tirage    int32             `csv:"1er_ou_2eme_tirage"`
	OtherData map[string]string `csv:"-"`
}

// Draw description data
type Draw struct {
	ID        string            `csv:"annee_numero_de_tirage"`
	Date      string            `csv:"date_de_tirage"`
	Day       string            `csv:"jour_de_tirage"`
	B1        int32             `csv:"boule_1"`
	B2        int32             `csv:"boule_2"`
	B3        int32             `csv:"boule_3"`
	B4        int32             `csv:"boule_4"`
	B5        int32             `csv:"boule_5"`
	Joker     int32             `csv:"numero_chance"`
	OtherData map[string]string `csv:"-"`
}

// Parse the csv file to loto result history
// newData provide a great type to decode the csv file parameter and addData function save the decoded data
func Parse(file string, newData func() interface{}, addData func(data interface{})) error {
	var err error
	var f *os.File

	if f, err = os.Open(file); err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return err
	}

	for {
		d := newData()
		if err := dec.Decode(d); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		addData(d)
	}
	return nil
}

// NewHistory read the files parameter to extract the draw history in date order
func NewHistory() (History, error) {
	var err error
	var h History

	for _, f := range oldFiles {
		if err = Parse(f, func() interface{} {
			return &OldDraw{}
		}, func(data interface{}) {
			h.OldDraws = append(h.OldDraws, *(data.(*OldDraw)))
		}); err != nil {
			return History{}, err
		}
	}

	for _, f := range recentFiles {
		if err = Parse(f, func() interface{} {
			return &Draw{}
		}, func(data interface{}) {
			h.RecentDraws = append(h.RecentDraws, *(data.(*Draw)))
		}); err != nil {
			return History{}, err
		}
	}
	return h, nil
}

// OrderRecent the recent list
func (h *History) OrderRecent() {
	sort.Slice(h.RecentDraws, func(i, j int) bool {
		return h.RecentDraws[i].ID < h.RecentDraws[j].ID
	})
}

// OrderOld the old list
func (h *History) OrderOld() {
	sort.Slice(h.OldDraws, func(i, j int) bool {
		return h.OldDraws[i].ID < h.OldDraws[j].ID ||
			(h.OldDraws[i].ID == h.OldDraws[j].ID && h.OldDraws[i].Tirage < h.OldDraws[j].Tirage)
	})
}

// DrawByIndex getter not available on old history
func (h History) DrawByIndex(index int) (Draw, error) {
	if index < 0 {
		return Draw{}, errors.New("invalid index < 0")
	}
	if len(h.RecentDraws) <= index {
		return Draw{}, errors.New("invalid index to get the draw by index")
	}
	return h.RecentDraws[index], nil
}

// Len getter not available on old history
func (h History) Len() int {
	return len(h.RecentDraws)
}
