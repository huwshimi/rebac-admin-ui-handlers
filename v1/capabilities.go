// Copyright 2024 Canonical Ltd.
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"net/http"

	"github.com/canonical/rebac-admin-ui-handlers/v1/resources"
)

// GetCapabilities returns the list of endpoints implemented by this API.
// (GET /capabilities)
func (h handler) GetCapabilities(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	capabilities, err := h.Capabilities.ListCapabilities(ctx)
	if err != nil {
		writeServiceErrorResponse(w, h.CapabilitiesErrorMapper, err)
		return
	}

	response := resources.GetCapabilitiesResponse{
		Meta: resources.ResponseMeta{
			Size: len(capabilities),
		},
		Data:   capabilities,
		Status: http.StatusOK,
	}

	writeResponse(w, http.StatusOK, response)
}
