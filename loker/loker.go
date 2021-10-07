package loker

import (
	"errors"
	"fmt"
	"strings"
)

type idCard struct {
	name   string
	number int
}

var loker []idCard

func InitLoker(capacity int) (res string, err error) {

	if len(loker) > 0 {
		err = errors.New("Loker sudah ada, silahkan exit program untuk membuat loker baru")
		return
	}

	// slice untuk menampung idCard
	loker = make([]idCard, capacity)

	res = fmt.Sprintf("Berhasil membuat loker dengan jumlah %d", len(loker))
	return
}

func LokerStatus() (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("Belum membuat Loker")
		return
	}

	tableRows := []string{
		"========================================================",
		"No Loker\t| Tipe Identitas\t| No Identitas",
		"--------------------------------------------------------",
	}

	for i, card := range loker {
		if card != (idCard{}) {
			tableRows = append(tableRows, fmt.Sprintf("%d\t\t| %s\t\t\t| %d", i+1, card.name, card.number))
		}
	}
	tableRows = append(tableRows, "========================================================")
	res = strings.Join(tableRows, "\n")
	return
}

func InputIdCard(cardType string, cardNo int) (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("Belum membuat Loker")
		return
	}

	// check duplikasi idCard
	for i, v := range loker {
		if v.number == cardNo {
			err = errors.New(fmt.Sprintf("Kartu dengan no identitas tersebut sudah ada di loker nomor %d", i+1))
			return
		}
	}

	cardType = strings.ToUpper(cardType)

	lokerId := 0
	for i, v := range loker {
		if v == (idCard{}) {
			loker[i] = idCard{cardType, cardNo}
			lokerId = i + 1
			break
		}
	}

	if lokerId == 0 {
		err = errors.New("Maaf loker sudah penuh")
		return
	}

	res = fmt.Sprintf("Kartu identitas tersimpan di loker nomor %d", lokerId)
	return
}

func LeaveIdCard(lokerId int) (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("Loker belum dibuat")
		return
	}

	if lokerId > len(loker) {
		// nomor loker tidak tersedia
		err = errors.New(fmt.Sprintf("Loker nomor %d tidak ditemukan", lokerId))
		return
	}
	if loker[lokerId-1] == (idCard{}) {
		// nomor loker tersedia & kosong
		err = errors.New(fmt.Sprintf("Loker nomor %d masih kosong", lokerId))
		return
	}

	loker[lokerId-1] = idCard{}

	res = fmt.Sprintf("Loker nomor %d berhasil dikosongkan", lokerId)
	return
}

func FindIdCard(cardNo int) (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("Loker belum dibuat")
		return
	}

	for i, card := range loker {
		if card.number == cardNo {
			res = fmt.Sprintf("Kartu identitas tersebut berada di loker nomor %d", i+1)
			return
		}
	}
	res = "Nomor identitas tidak ditemukan"
	return
}

func SearchIdCard(cardType string) (res string, err error) {

	if len(loker) == 0 {
		err = errors.New("Loker belum dibuat")
		return
	}

	cardType = strings.ToUpper(cardType)

	var searchResults []string
	for _, card := range loker {
		if card.name == cardType {
			searchResults = append(searchResults, fmt.Sprintf("%d", card.number))
		}
	}

	if len(searchResults) == 0 {
		err = errors.New(fmt.Sprintf("Tidak ada kartu identitas dengan tipe %s", cardType))
		return
	}

	res = fmt.Sprintf("No Identitas: %s", strings.Join(searchResults, ", "))
	return
}
