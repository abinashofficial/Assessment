package handlers

import (
	"Assessment/loged"
	"Assessment/model"
	"Assessment/services"
	"Assessment/tapcontext"
	"Assessment/utils"
	"encoding/json"
	"net/http"
)

func New(service services.Store) FormHandlers {
	return &formHandler{
		formServ: service.FormServ,
	}
}

type formHandler struct {
	formServ services.Service
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
		loged.GenericInfo(ctx, functionDesc, loged.FieldsMap{"errMsg": err})
		utils.ErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError, err, nil)
		return
	}

	form, err = h.formServ.Create(req)
	if err != nil {
		loged.GenericInfo(ctx, functionDesc, loged.FieldsMap{"errMsg": err})
		utils.ErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError, err, nil)
		return
	}

	utils.ReturnResponse(w, http.StatusOK, form)
	loged.GenericInfo(ctx, functionDesc, loged.FieldsMap{"message": "Success", "forms": ""})
}
