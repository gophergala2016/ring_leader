package services

import (
	"errors"
	db "github.com/dancannon/gorethink"
	"github.com/gophergala2016/ring_leader/resources"
	//"log"
)

type ResourceService struct {
}

func (s *ResourceService) InsertResource(DB *db.Session, src resources.Resource) error {
	if src == nil {
		return errors.New("Nil Service")
	}
	res, err := db.Table("resources").Filter(db.Row.Field("name").Eq(src.GetName())).Count().Run(DB)
	if err != nil {
		return err
	}
	var cnt int
	err = res.One(&cnt)
	defer res.Close()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("resource exist with name already")
	}
	_, err = db.Table("resources").Insert(src).RunWrite(DB)
	if err != nil {
		return err
	}
	return nil
}

func (s *ResourceService) GetResources(DB *db.Session, resType string) ([]resources.Resource, error) {
	result, err := db.Table("resources").Filter(db.Row.Field("type").Eq(resType)).Run(DB)
	if err != nil {
		return nil, err
	}
	var ret []resources.Resource
	switch resType {
	case "software":
		var res []*resources.Software
		err = result.All(&res)
		if err != nil {
			return nil, err
		}
		ret = make([]resources.Resource, len(res))
		for i, v := range res {
			ret[i] = resources.Resource(v)
		}
	default:
		return nil, errors.New("Unknown resource type: " + resType)
	}

	return ret, nil
}
