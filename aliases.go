// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"github.com/knocklabs/knock-mgmt-go/internal/apierror"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// The information about a paginated result.
//
// This is an alias to an internal type.
type PageInfo = shared.PageInfo
