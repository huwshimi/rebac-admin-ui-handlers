// Copyright 2024 Canonical Ltd.
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// ResourcesService defines an abstract backend to handle Resources related operations.
type ResourcesService interface {
	// ListResources returns a page of Resource objects of at least `size` elements if available.
	ListResources(ctx context.Context, params *resources.GetResourcesParams) (*resources.PaginatedResponse[resources.Resource], error)
}