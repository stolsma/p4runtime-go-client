package p4info

import v1 "github.com/p4lang/p4runtime/go/p4/config/v1"

type Action struct {
	preamble *v1.Preamble
	params   []*ActionParam
}

func getActions(rawActions []*v1.Action) (*Collection, error) {
	actions := GetNewCollection()

	for _, action := range rawActions {
		var params []*ActionParam
		for _, p := range action.Params {
			params = append(params, &ActionParam{
				param: p,
			})
		}

		actions.AddRecord(&Action{
			preamble: action.GetPreamble(),
			params:   params,
		})
	}

	return actions, nil
}

func (a *Action) GetName() string {
	return a.preamble.GetName()
}

func (a *Action) GetAlias() string {
	return a.preamble.GetAlias()
}

func (a *Action) GetId() uint32 {
	return a.preamble.GetId()
}

func (a *Action) GetAnnotations() []string {
	return a.preamble.GetAnnotations()
}

func (a *Action) GetDocBrief() string {
	return a.preamble.Doc.GetBrief()
}

func (a *Action) GetDocDescription() string {
	return a.preamble.Doc.GetDescription()
}
