package flowcontext

import (
	"github.com/kasv2/kasv2d/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}
