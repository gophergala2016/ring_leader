package resources

import (
	//"github.com/gophergala2016/ring_leader/credentials"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
)

type Resource interface {
	//Requirements() map[string]string
	//Grant(credentials.Credential)
	//Revoke(credentials.Credential)
	GetName() string
}

type ResourceInfo struct {
	Id   string `gorethink:"id,omitempty" json:"id"`
	Name string `json:"name" gorethink:"name" binding:"required"`
	Type string `json:"type" gorethink:"type" binding:"required"`
}

func (i ResourceInfo) GetName() string {
	return i.Name
}

func UnmarshalJSON(req *http.Request) (Resource, error) {
	if req.Body == nil {
		return nil, errors.New("No body found")
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var t json.RawMessage
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	log.Println("1", t)
	sw := new(Software)
	if sw.Unmarshal(t) {
		log.Println("2", sw)
		return sw, nil
	}
	return nil, errors.New("Could not find type of resource")
}
