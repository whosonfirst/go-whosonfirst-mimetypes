package http

import (
	"encoding/json"
	"github.com/whosonfirst/go-whosonfirst-mimetypes"
	gohttp "net/http"
)

func ExtensionsHandler() (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		query := req.URL.Query()
		mimetype := query.Get("mimetype")

		if mimetype == "" {
			gohttp.Error(rsp, "Missing mimetype parameter", gohttp.StatusBadRequest)		
			return
		}

		extensions := mimetypes.ExtensionsByType(mimetype)

		enc, err := json.Marshal(extensions)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)		
			return
		}
		
		rsp.Header().Set("Content-Type", "application/json")
		rsp.Header().Set("Access-Control-Allow-Origin", "*")
								
		rsp.Write(enc)
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
