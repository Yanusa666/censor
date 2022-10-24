package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"net/http"
	"sf-censor/internal/censor"
	"sf-censor/internal/config"
	"sf-censor/internal/constants"
)

type Handler struct {
	cfg    *config.Config
	lgr    zerolog.Logger
	censor *censor.Censor
}

func NewHandler(cfg *config.Config, lgr zerolog.Logger, censor *censor.Censor) *Handler {
	return &Handler{
		cfg:    cfg,
		lgr:    lgr,
		censor: censor,
	}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	request := new(CheckReq)
	err := decoder.Decode(request)
	if err != nil {
		resp, _ := json.Marshal(ErrorResp{Error: fmt.Sprintf("incorrect request: %s", err.Error())})
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, string(resp))
		return
	}

	ctx := r.Context()
	requestId, _ := ctx.Value(constants.RequestIdKey).(string)

	lgr := h.lgr.With().
		Str("handler", "Check").
		Str(constants.RequestIdKey, requestId).
		Dict("request", zerolog.Dict().
			Str("text", request.Text)).
		Logger()

	isApproved := h.censor.Check(request.Text)
	if err != nil {
		resp, _ := json.Marshal(ErrorResp{Error: fmt.Sprintf("internal error: %s", err.Error())})
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, string(resp))
		return
	}

	lgr.Debug().Msg("executed")

	resp, _ := json.Marshal(CheckResp{Status: isApproved})
	fmt.Fprintf(w, string(resp))
}
