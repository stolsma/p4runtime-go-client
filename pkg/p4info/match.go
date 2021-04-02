package p4info

type Match struct {
	name  string
	id    uint32
	mType string
}

func (m *Match) GetName() string {
	return m.name
}

func (m *Match) GetAlias() string {
	return ""
}

func (m *Match) GetId() uint32 {
	return m.id
}

func (m *Match) GetType() string {
	return m.mType
}
