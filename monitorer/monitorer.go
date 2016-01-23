package monitorer

import (
	"github.com/gophergala2016/ring_leader/policy"
	"github.com/gophergala2016/ring_leader/resources"
)

type Monitorer interface {
	Monitor(resources.Resource, policy.Policy)
	Forget(resources.Resource)
}
