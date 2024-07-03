// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package module

import (
	"github.com/hashicorp/hcl/v2"
)

type Resource struct {
	Type        string
	Name        string
	Description string
	Provider    ProviderRef
	Mode        string
	RangePtr    *hcl.Range
}
