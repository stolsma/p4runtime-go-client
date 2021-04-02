package p4info

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

type RecordInterface interface {
	GetName() string
	GetAlias() string
	GetId() uint32
}

type Collection struct {
	name  map[string]interface{}
	id    map[uint32]interface{}
	len   uint32
	names []string
}

func GetNewCollection() *Collection {
	return &Collection{
		name: map[string]interface{}{},
		id:   map[uint32]interface{}{},
		len:  0,
	}
}

func (c *Collection) AddRecord(record interface{}) {
	// TODO(stolsma) Check if given record supports GetName(), GetAlias and GetId()
	if reflect.ValueOf(record).Kind() != reflect.Ptr {
		log.Fatalf("Collection.AddRecord: Did not receive a pointer!")
	}

	name := record.(RecordInterface).GetName()
	c.names = append(c.names, name)
	c.name[name] = record

	alias := record.(RecordInterface).GetAlias()
	if alias != "" && name != alias {
		c.names = append(c.names, alias)
		c.name[alias] = record
	}

	c.id[record.(RecordInterface).GetId()] = record.(RecordInterface)
	c.len++
}

func (c *Collection) GetWithName(name string) interface{} {
	if record, ok := c.name[name]; ok {
		return record
	}
	return nil
}

func (c *Collection) GetWithId(id uint32) interface{} {
	if record, ok := c.id[id]; ok {
		return record
	}
	return nil
}

func (c *Collection) GetIdWithName(name string) uint32 {
	if record, ok := c.name[name]; ok {
		return record.(RecordInterface).GetId()
	}
	return 0
}

func (c *Collection) GetNameWithId(id uint32) string {
	if record, ok := c.id[id]; ok {
		return record.(RecordInterface).GetName()
	}
	return ""
}

func (c *Collection) GetNames() []string {
	return c.names
}
