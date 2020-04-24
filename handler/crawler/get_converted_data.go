package crawler

import (
	"net/http"
	"net/url"

	"github.com/dexterorion/smartmei/errors"
	"github.com/dexterorion/smartmei/lib"
	"github.com/dexterorion/smartmei/model"
)

func (handler methods) getConvertedDataHandler(w http.ResponseWriter, r *http.Request) {
	paramURL := r.URL.Query().Get("url")
	u, err := url.Parse(paramURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		lib.JSONError(w, errors.URLCannotBeEmpty(), http.StatusInternalServerError)
		return
	}

	result, err := handler.business.GetConvertedData(r.Context(), paramURL, model.CurrencyBRL, []string{model.CurrencyUSD, model.CurrencyEUR})
	if err != nil {
		lib.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	lib.JSONReturn(w, http.StatusOK, result)
}
