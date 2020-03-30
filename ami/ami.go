package ami

import specv1 "github.com/baetyl/baetyl-go/spec/v1"

//go:generate mockgen -destination=../mock/ami.go -package=mock github.com/baetyl/baetyl-core/ami AMI

// AMI app model interfaces
type AMI interface {
	Collect() (specv1.Report, error)
	Apply([]specv1.AppInfo) error
}
