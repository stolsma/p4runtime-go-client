package p4info

import v1 "github.com/p4lang/p4runtime/go/p4/config/v1"

type ActionParam struct {
	param *v1.Action_Param
}

func (ap *ActionParam) GetName() string {
	return ap.param.GetName()
}

func (ap *ActionParam) GetAlias() string {
	return ""
}

func (ap *ActionParam) GetId() uint32 {
	return ap.param.GetId()
}

func (ap *ActionParam) GetAnnotations() []string {
	return ap.param.GetAnnotations()
}

func (ap *ActionParam) GetBitWidth() int32 {
	return ap.param.GetBitwidth()
}

func (ap *ActionParam) GetDocBrief() string {
	return ap.param.Doc.GetBrief()
}

func (ap *ActionParam) GetDocDescription() string {
	return ap.param.Doc.GetDescription()
}

func (ap *ActionParam) GetType() string {
	return ap.param.GetTypeName().GetName()
}
