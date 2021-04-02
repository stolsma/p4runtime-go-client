package p4info

import v1 "github.com/p4lang/p4runtime/go/p4/config/v1"

type ActionProfile struct {
	preamble     *v1.Preamble
	Tables       *Collection
	withSelector bool
	size         int64
	maxGroupSize int32
}

func getActionProfiles(rawActionProfiles []*v1.ActionProfile, tables *Collection) (*Collection, error) {
	actionProfiles := GetNewCollection()

	for _, actionProfile := range rawActionProfiles {
		apTables := GetNewCollection()
		for _, tid := range actionProfile.TableIds {
			table := tables.GetWithId(tid)
			// TODO(stolsma) return error if table is not found
			apTables.AddRecord(table)
		}

		actionProfiles.AddRecord(&ActionProfile{
			preamble:     actionProfile.GetPreamble(),
			Tables:       apTables,
			withSelector: actionProfile.GetWithSelector(),
			size:         actionProfile.GetSize(),
			maxGroupSize: actionProfile.GetMaxGroupSize(),
		})
	}

	return actionProfiles, nil
}

func (a *ActionProfile) GetName() string {
	return a.preamble.GetName()
}

func (a *ActionProfile) GetAlias() string {
	return a.preamble.GetAlias()
}

func (a *ActionProfile) GetId() uint32 {
	return a.preamble.GetId()
}

func (a *ActionProfile) GetAnnotations() []string {
	return a.preamble.GetAnnotations()
}

func (a *ActionProfile) GetDocBrief() string {
	return a.preamble.Doc.GetBrief()
}

func (a *ActionProfile) GetDocDescription() string {
	return a.preamble.Doc.GetDescription()
}

func (a *ActionProfile) WithSelector() bool {
	return a.withSelector
}

func (a *ActionProfile) GetSize() int64 {
	return a.size
}

func (a *ActionProfile) GetMaxGroupSize() int32 {
	return a.maxGroupSize
}
