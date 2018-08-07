package handlers

import (
	"net/http"
	"strconv"

	"github.com/wlwanpan/wifFee/services/models"
	mgo "gopkg.in/mgo.v2"
)

func LogWifi(w http.ResponseWriter, r *http.Request) (int, error) {
	db := r.Context().Value("db").(*mgo.Session)

	err := r.ParseForm()
	if err != nil {
		return http.StatusBadRequest, err
	}

	wif := models.Wifi{
		PlaceID:   r.Form.Get("placeID"),
		Ping:      ToUint64(r.Form.Get("ping")),
		UpSpeed:   ToUint32(r.Form.Get("downSpeed")),
		DownSpeed: ToUint32(r.Form.Get("downSpeed")),
	}

	err = wif.Create(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func ToUint64(s string) uint64 {

	parsedUint, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return parsedUint
}

func ToUint32(s string) uint64 {

	parsedUint, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return parsedUint
}
