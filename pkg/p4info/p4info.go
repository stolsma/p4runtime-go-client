package p4info

import v1 "github.com/p4lang/p4runtime/go/p4/config/v1"

type P4Info struct {
	PkgInfo        *PkgInfo
	Tables         *Collection
	actions        *Collection
	ActionProfiles *Collection
	//	Counters                 []*Counter                  `json:"counters,omitempty"`
	DirectCounters *Collection
	//	Meters                   []*Meter                    `json:"meters,omitempty"`
	//	DirectMeters             []*DirectMeter              `json:"direct_meters,omitempty"`
	//	ControllerPacketMetadata []*ControllerPacketMetadata `json:"controller_packet_metadata,omitempty"`
	//	ValueSets                []*ValueSet                 `json:"value_sets,omitempty"`
	//	Registers                []*Register                 `json:"registers,omitempty"`
	//	Digests                  []*Digest                   `json:"digests,omitempty"`
	//	Externs                  []*Extern                   `json:"externs,omitempty"`
	//	TypeInfo                 *P4TypeInfo                 `json:"type_info,omitempty"`
}

func GetNewP4Info(rawP4Info *v1.P4Info) *P4Info {
	// Initialize empty collection
	var p4Info = &P4Info{}

	if rawP4Info == nil {
		return nil
	}

	p4Info.PkgInfo = getPkgInfo(rawP4Info.GetPkgInfo())
	p4Info.actions, _ = getActions(rawP4Info.GetActions())
	p4Info.Tables, _ = getTables(rawP4Info.GetTables(), p4Info.actions)
	p4Info.ActionProfiles, _ = getActionProfiles(rawP4Info.GetActionProfiles(), p4Info.Tables)
	p4Info.DirectCounters, _ = getDirectCounters(rawP4Info.GetDirectCounters(), p4Info.Tables)

	return p4Info
}
