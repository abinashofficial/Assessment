package api

import (
	"Assessment/log"
	_ "Assessment/log"
	"Assessment/tapcontext"
	"Assessment/utils"
	"bufio"
	"context"
	"fmt"
	"go.elastic.co/apm"
	"net/http"
	"os"
)

// Ping
//
//	@Summary	Ping API
//	@Schemes
//	@Tags			Health Check
//	@Description	API to check the service availability
//	@Success		200	{string}	string	"alive"
//	@Router			/ping [get]
//	@Router			/public/ping [get]
//	@Router			/private/ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive")
}

func RenderAPIDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/swagger.json")
}

// @Summary	Get Git Version
// @Schemes
// @Tags			Git Version
// @Description	API to fetch the current git version
// @Success		200	{string}	string	"tap-release-2023.6.1.0"
// @Failure		500	{object}	utils.ErrResponse
// @Router			/public/version [get]

func GetGitVersion(w http.ResponseWriter, r *http.Request) {
	ctx := tapcontext.UpgradeCtx(r.Context())
	functionDesc := "GetGitVersion"
	var version string

	span, _ := apm.StartSpan(ctx, functionDesc, "handler")
	defer span.End()

	//version, err := exec.Command("git branch --show-current").Output()

	file, err := os.Open("./.git/version")
	if err != nil {
		log.GenericInfo(ctx, functionDesc, log.FieldsMap{"err": err.Error()})
		utils.ErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError, err, nil)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		version = scanner.Text()
		break
	}
	if err = scanner.Err(); err != nil {
		log.GenericInfo(ctx, functionDesc, log.FieldsMap{"err": err.Error()})
		utils.ErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError, err, nil)
		return
	}
	utils.ReturnResponse(w, http.StatusOK, version)
	log.GenericInfo(ctx, functionDesc, log.FieldsMap{"msg": "Success", "version": version})
}

func MqttKey(w http.ResponseWriter, r *http.Request) {
	ctx := tapcontext.UpgradeCtx(r.Context())
	functionDesc := "MqttKey"
	type key struct {
		Zzz string `json:"zzz"`
	}

	response := key{
		Zzz: os.Getenv("MQTT_ENCRYPT"),
	}

	utils.ReturnResponse(w, http.StatusOK, response)
	log.GenericInfo(ctx, functionDesc, log.FieldsMap{"msg": "Success"})
	context.Background()
}
