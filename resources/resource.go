package resources

import (
	"github.com/gophergala2016/ring_leader/credentials"
)

type Resource interface {
	Requirements() map[string]string
	Grant(credentials.Credential)
	Revoke(credentials.Credential)
}
