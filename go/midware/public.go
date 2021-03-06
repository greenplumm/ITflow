package midware

import (
	"encoding/json"
	"io/ioutil"
	"itflow/db"
	"itflow/internal/response"
	"net/http"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func JsonToStruct(w http.ResponseWriter, r *http.Request) bool {
	resp := &response.Response{}
	if goconfig.ReadBool("debug", false) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err)
			w.Write(resp.ErrorE(err))
			return true
		}
		golog.Info(string(b))
		err = json.Unmarshal(b, xmux.GetData(r).Data)
		if err != nil {
			golog.Error(err)
			w.Write(resp.ErrorE(err))
			return true
		}
	} else {
		err := json.NewDecoder(r.Body).Decode(xmux.GetData(r).Data)
		if err != nil {
			golog.Error(err)
			w.Write(resp.ErrorE(err))
			return true
		}

	}
	return false
}

func CheckToken(w http.ResponseWriter, r *http.Request) bool {
	errorcode := &response.Response{}
	a := r.Header.Get("X-Token")
	if a == "" {
		golog.Error("not found token")
		w.Write(errorcode.TokenNotFound())
		return true
	}
	filter, err := db.Table.Filter("Token", a)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.TokenNotFound())
		return true
	}
	var nickname string
	var uid int64
	err = filter.Get(db.NICKNAME, db.ID).Scan(&nickname, &uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.TokenNotFound())
		return true
	}
	xmux.GetData(r).Set("nickname", nickname)
	xmux.GetData(r).Set("uid", uid)

	return false
}
