package p4info

import v1 "github.com/p4lang/p4runtime/go/p4/config/v1"

type Table struct {
	preamble *v1.Preamble
	match    *Collection
	action   *Collection
}

func getTables(rawTables []*v1.Table, actions *Collection) (*Collection, error) {
	tables := GetNewCollection()

	for _, rawTable := range rawTables {
		t := &Table{
			preamble: rawTable.GetPreamble(),
			match:    GetNewCollection(),
			action:   GetNewCollection(),
		}

		for _, m := range rawTable.MatchFields {
			name := m.GetName()
			t.match.AddRecord(&Match{
				name:  name,
				id:    m.GetId(),
				mType: m.GetMatchType().String(),
			})
		}

		for _, a := range rawTable.ActionRefs {
			t.action.AddRecord(actions.GetWithId(a.GetId()))
		}

		tables.AddRecord(t)
	}

	return tables, nil
}

func (t *Table) GetName() string {
	return t.preamble.GetName()
}

func (t *Table) GetAlias() string {
	return t.preamble.GetAlias()
}

func (t *Table) GetId() uint32 {
	return t.preamble.GetId()
}

func (t *Table) GetAnnotations() []string {
	return t.preamble.GetAnnotations()
}

func (t *Table) GetDocBrief() string {
	return t.preamble.Doc.GetBrief()
}

func (t *Table) GetDocDescription() string {
	return t.preamble.Doc.GetDescription()
}

/*
func (t *Table) NewTableEntry(
	mfs map[string]MatchInterface,
	action *p4_v1.TableAction,
	options *TableEntryOptions,
) *p4_v1.TableEntry {

	entry := &p4_v1.TableEntry{
		TableId:         t.id,
		IsDefaultAction: (mfs == nil),
		Action:          action,
	}

	for id, mf := range mfs {
		mid := t.matchId(id)
		entry.Match = append(entry.Match, mf.get(uint32(mid)))
	}

	if options != nil {
		entry.IdleTimeoutNs = options.IdleTimeout.Nanoseconds()
	}

	return entry
}
*/

func (t *Table) GetMatchIdWithName(name string) uint32 {
	return t.match.GetIdWithName(name)
}

func (t *Table) GetMatchNameWithId(id uint32) string {
	return t.match.GetNameWithId(id)
}

func (t *Table) GetMatchWithName(name string) *Match {
	match := t.match.GetWithName(name)
	return match.(*Match)
}

func (t *Table) GetMatchWithId(id uint32) *Match {
	match := t.match.GetWithId(id)
	return match.(*Match)
}
