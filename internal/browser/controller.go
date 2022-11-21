package browser

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"reflect"
	"ws/pkg/types"
)

type Controller struct {
	types.DatabaseController
}

func (c *Controller) GetImage(w http.ResponseWriter, r *http.Request) {
	var model Model

	params := mux.Vars(r)
	id, find := params["id"]

	if !find {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err := c.DB.Get(&model, "SELECT * FROM screen WHERE uuid=:uuid LIMIT 1", id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", http.DetectContentType(model.Data))
	w.Write(model.Data)
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	var resources []Resource
	var models []Model

	w.Header().Set("Content-Type", "application/json")

	if err := c.DB.Select(&models, "SELECT * FROM screen"); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	for _, model := range models {
		resources = append(resources, model.toResource())
	}

	if err := json.NewEncoder(w).Encode(resources); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

}

func (c *Controller) CreateScreenshot(w http.ResponseWriter, r *http.Request) {
	var img []byte
	var model Model

	br := new(Service)
	br.Init()
	defer br.Close()
	model.Init(c.DB)

	var request ScreenshotRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	u, err := url.Parse(request.Url)
	if err != nil {
		http.Error(w, "failed parse url", http.StatusBadRequest)
		return
	}

	img, err = br.Screenshot(request.Element, u)

	if err != nil {
		t := reflect.TypeOf(err)
		if t.String() == "context.deadlineExceededError" {
			err = errors.New("find element timeout")
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.UUID = uuid.NewString()
	model.Url = request.Url
	model.Element = request.Element
	model.Status = StatusWait
	model.Data = img

	_, err = model.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.toResource()); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
}
