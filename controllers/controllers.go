package controllers

import (
	"errors"
	"loker-app/helpers"
	"loker-app/loker"
	"regexp"
	"strconv"
)

func Init(capacity string) {
	// Fungsi untuk menangani perintah `init`

	var res string
	var err error
	var lokerCap int

	if capacity == "" {
		err = errors.New("perintah kurang tepat")
		helpers.ShowResult(res, err)
		return
	}

	lokerCap, err = strconv.Atoi(capacity)
	if err != nil || lokerCap < 1 {
		err = errors.New("nilai yang Anda masukkan salah")
		helpers.ShowResult(res, err)
		return
	}

	res, err = loker.InitLoker(lokerCap)
	helpers.ShowResult(res, err)
}

func Status() {
	// Fungsi untuk menangani perintah `status`
	res, err := loker.LokerStatus()
	helpers.ShowResult(res, err)
}

func Input(cardType, cardNo string) {
	// Fungsi untuk menangani perintah `input`

	var res string
	var err error
	var cardNumber int

	if cardType == "" || cardNo == "" {
		err = errors.New("perintah kurang tepat")
		helpers.ShowResult(res, err)
		return
	}

	regex, _ := regexp.Compile(`^[a-zA-Z]+$`)
	cardTypeIsAlphabet := regex.MatchString(cardType)

	cardNumber, err = strconv.Atoi(cardNo)

	if err != nil || cardNumber < 1 || !cardTypeIsAlphabet {
		err = errors.New("nilai yang Anda masukkan salah")
		helpers.ShowResult(res, err)
		return
	}

	res, err = loker.InputIdCard(cardType, cardNumber)
	helpers.ShowResult(res, err)
}

func Leave(lokerNo string) {
	// Fungsi untuk menangani perintah `leave`

	var res string
	var err error
	var lokerId int

	if lokerNo == "" {
		err = errors.New("perintah kurang tepat")
		helpers.ShowResult(res, err)
		return
	}

	lokerId, err = strconv.Atoi(lokerNo)
	if err != nil || lokerId < 1 {
		err = errors.New("nilai yang Anda masukkan salah")
		helpers.ShowResult(res, err)
		return
	}

	res, err = loker.LeaveIdCard(lokerId)
	helpers.ShowResult(res, err)
}

func Find(cardNo string) {
	// Fungsi untuk menangani perintah `find`

	var res string
	var err error
	var cardNumber int

	if cardNo == "" {
		err = errors.New("perintah kurang tepat")
		helpers.ShowResult(res, err)
		return
	}

	cardNumber, err = strconv.Atoi(cardNo)
	if err != nil || cardNumber < 1 {
		err = errors.New("nilai yang Anda masukkan salah")
		helpers.ShowResult(res, err)
		return
	}

	res, err = loker.FindIdCard(cardNumber)
	helpers.ShowResult(res, err)
}

func Search(cardType string) {
	// Fungsi untuk menangani perintah `search`

	var res string
	var err error

	if cardType == "" {
		err = errors.New("perintah kurang tepat")
		helpers.ShowResult(res, err)
		return
	}

	res, err = loker.SearchIdCard(cardType)
	helpers.ShowResult(res, err)
}
