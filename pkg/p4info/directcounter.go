package p4info

import v1 "github.com/p4lang/p4runtime/go/p4/config/v1"

type CounterSpec_Unit v1.CounterSpec_Unit

const (
	CounterSpec_UNSPECIFIED CounterSpec_Unit = CounterSpec_Unit(v1.CounterSpec_UNSPECIFIED)
	CounterSpec_BYTES       CounterSpec_Unit = CounterSpec_Unit(v1.CounterSpec_BYTES)
	CounterSpec_PACKETS     CounterSpec_Unit = CounterSpec_Unit(v1.CounterSpec_PACKETS)
	CounterSpec_BOTH        CounterSpec_Unit = CounterSpec_Unit(v1.CounterSpec_BOTH)
)

var CounterSpec_Unit_name = v1.CounterSpec_Unit_name
var CounterSpec_Unit_value = v1.CounterSpec_Unit_value

type DirectCounter struct {
	preamble    *v1.Preamble
	counterSpec CounterSpec_Unit
	directTable *Table
}

func getDirectCounters(rawDirectCounters []*v1.DirectCounter, tables *Collection) (*Collection, error) {
	directCounters := GetNewCollection()

	for _, dc := range rawDirectCounters {
		directCounters.AddRecord(&DirectCounter{
			preamble:    dc.GetPreamble(),
			directTable: tables.GetWithId(dc.GetDirectTableId()).(*Table),
		})
	}

	return directCounters, nil
}

func (dc *DirectCounter) GetName() string {
	return dc.preamble.GetName()
}

func (dc *DirectCounter) GetAlias() string {
	return dc.preamble.GetAlias()
}

func (dc *DirectCounter) GetId() uint32 {
	return dc.preamble.GetId()
}

func (dc *DirectCounter) GetAnnotations() []string {
	return dc.preamble.GetAnnotations()
}

func (dc *DirectCounter) GetDocBrief() string {
	return dc.preamble.Doc.GetBrief()
}

func (dc *DirectCounter) GetDocDescription() string {
	return dc.preamble.Doc.GetDescription()
}

func (dc *DirectCounter) GetCounterSpec() CounterSpec_Unit {
	return dc.counterSpec
}

func (dc *DirectCounter) GetDirectTable() *Table {
	return dc.directTable
}
