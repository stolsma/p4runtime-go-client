package p4info

import v1 "github.com/p4lang/p4runtime/go/p4/config/v1"

type PkgInfo struct {
	pkgInfo *v1.PkgInfo
}

func getPkgInfo(rawPkgInfo *v1.PkgInfo) *PkgInfo {
	return &PkgInfo{
		pkgInfo: rawPkgInfo,
	}
}

func (pi *PkgInfo) GetName() string {
	return pi.pkgInfo.GetName()
}

func (pi *PkgInfo) GetVersion() string {
	return pi.pkgInfo.GetVersion()
}

func (pi *PkgInfo) GetDocBrief() string {
	return pi.pkgInfo.Doc.GetBrief()
}

func (pi *PkgInfo) GetDocDescription() string {
	return pi.pkgInfo.Doc.GetDescription()
}

func (pi *PkgInfo) GetAnnotations() []string {
	return pi.pkgInfo.GetAnnotations()
}

func (pi *PkgInfo) GetArch() string {
	return pi.pkgInfo.GetArch()
}

func (pi *PkgInfo) GetOrganization() string {
	return pi.pkgInfo.GetOrganization()
}

func (pi *PkgInfo) GetContact() string {
	return pi.pkgInfo.GetContact()
}

func (pi *PkgInfo) GetUrl() string {
	return pi.pkgInfo.GetUrl()
}
