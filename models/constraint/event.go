package constraint

type ConstraintEventResp struct {
	Pass bool `json:"pass"`

	ConstraintType ConstraintType `json:"constraintType,omitempty"`

	ConstraintEventResps []*ConstraintEventResp `json:"constraintEventResps,omitempty"`

	Matches []string `json:"matches"`
	Misses  []string `json:"misses"`
}
