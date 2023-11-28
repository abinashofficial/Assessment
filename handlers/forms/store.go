package forms

import (
	"Assessment/log"
	"Assessment/model"
	"Assessment/services/forms"
	"Assessment/tapcontext"
	"Assessment/utils"
	"encoding/json"
	"net/http"
)

func New(service forms.Store) FormHandlers {
	return &formHandler{
		formServ: service.FormServ,
	}
}

type formHandler struct {
	formServ forms.Service
}

func (h formHandler) HandlerCreate(w http.ResponseWriter, r *http.Request) {
	functionDesc := "Create Form"

	var req map[string]string
	var form model.ConvertedRequest

	ctx := tapcontext.UpgradeCtx(r.Context())
	//if err := utils.ValidateUserEmail(ctx.UserEmail); err != nil {
	//	utils.ErrorResponse(ctx, w, err.Error(), http.StatusBadRequest, err, nil)
	//	return
	//}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.GenericInfo(ctx, functionDesc, log.FieldsMap{"errMsg": err})
		utils.ErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError, err, nil)
		return
	}

	form, err = h.formServ.Create(req)
	if err != nil {
		log.GenericInfo(ctx, functionDesc, log.FieldsMap{"errMsg": err})
		utils.ErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError, err, nil)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, form)
	log.GenericInfo(ctx, functionDesc, log.FieldsMap{"message": "Success", "forms": ""})
}
