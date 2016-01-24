package resources

import (
	"encoding/json"
	"time"
	"log"
)

// Software License
type License struct {
	Key  string `json:"key" gorethink:"key"`
	Expiration time.Time `json:"expiration" gorethink:"expiration"`
	NeverExpires bool `json:"neverExpires" gorethink:"never_expires"`
	Capacity int `json:"capacity" gorethink:"capacity"`
	Users []string `json:"users" gorethink:"users"`
}

type Software struct {
	ResourceInfo
	Licenses []License `json:"licenses" gorethink:"licenses"`
}

func (sw *Software) Unmarshal(raw json.RawMessage) bool {
	log.Println("3")
	if err := json.Unmarshal(raw, sw); err != nil {
		log.Println("55", err.Error())
		return false
	}
	log.Printf("4 %v\n", sw)

	if  sw.Name != "" && sw.Type == "software" {
		return true
	}
	return false
}
