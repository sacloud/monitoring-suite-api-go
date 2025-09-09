// Copyright 2025- The sacloud/monitoring-suite-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitoringsuite_test

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	. "github.com/sacloud/monitoring-suite-api-go"
	v1 "github.com/sacloud/monitoring-suite-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

func TestLogRoutingOp_List(t *testing.T) {
	expected := v1.PaginatedLogRoutingList{
		IsOk:    v1.NewOptBool(true),
		Count:   1,
		From:    0,
		Results: []v1.LogRouting{TemplateLogRouting},
	}
	client := newTestClient(expected)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	routings, err := api.List(ctx, v1.LogsRoutingsListParams{})
	require.NoError(t, err)
	require.NotNil(t, routings)
	require.Equal(t, 1, len(routings))
}

func TestLogRoutingOp_Read(t *testing.T) {
	client := newTestClient(TemplateWrappedLogRouting)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	res, err := api.Read(ctx, "12345")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, TemplateWrappedLogRouting.GetID(), res.GetID())
	require.Equal(t, TemplateWrappedLogRouting.GetPublisher(), res.GetPublisher())
	require.Equal(t, TemplateWrappedLogRouting.GetPublisherCode(), res.GetPublisherCode())
	require.Equal(t, TemplateWrappedLogRouting.GetLogStorage(), res.GetLogStorage())
	require.Equal(t, TemplateWrappedLogRouting.GetLogStorageID(), res.GetLogStorageID())
	require.Equal(t, TemplateWrappedLogRouting.GetResourceID(), res.GetResourceID())
	require.Equal(t, TemplateWrappedLogRouting.GetVariant(), res.GetVariant())
}

func TestLogRoutingOp_Read_404(t *testing.T) {
	expected := newErrorResponse(404, "No LogRouting matches the given query.")
	client := newTestClient(expected, http.StatusNotFound)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	routing, err := api.Read(ctx, "99999")
	require.Nil(t, routing)
	require.Error(t, err)
	require.ErrorContains(t, err, "Not Found")
}

func TestLogRoutingOp_Create(t *testing.T) {
	client := newTestClient(TemplateWrappedLogRouting, http.StatusCreated)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	createReq := TemplateLogRouting
	res, err := api.Create(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, TemplateWrappedLogRouting.GetID(), res.GetID())
	require.Equal(t, TemplateWrappedLogRouting.GetPublisher(), res.GetPublisher())
	require.Equal(t, TemplateWrappedLogRouting.GetPublisherCode(), res.GetPublisherCode())
	require.Equal(t, TemplateWrappedLogRouting.GetLogStorage(), res.GetLogStorage())
	require.Equal(t, TemplateWrappedLogRouting.GetLogStorageID(), res.GetLogStorageID())
	require.Equal(t, TemplateWrappedLogRouting.GetResourceID(), res.GetResourceID())
	require.Equal(t, TemplateWrappedLogRouting.GetVariant(), res.GetVariant())
}

func TestLogRoutingOp_Create_400(t *testing.T) {
	expected := newErrorResponse(400, "Invalid request body.")
	client := newTestClient(expected, http.StatusBadRequest)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	createReq := v1.LogRouting{}
	routing, err := api.Create(ctx, createReq)
	require.Nil(t, routing)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid")
}

func TestLogRoutingOp_Update(t *testing.T) {
	client := newTestClient(TemplateWrappedLogRouting)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	updateReq := TemplateLogRouting
	res, err := api.Update(ctx, "12345", &updateReq)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, TemplateWrappedLogRouting.GetID(), res.GetID())
	require.Equal(t, TemplateWrappedLogRouting.GetPublisher(), res.GetPublisher())
	require.Equal(t, TemplateWrappedLogRouting.GetPublisherCode(), res.GetPublisherCode())
	require.Equal(t, TemplateWrappedLogRouting.GetLogStorage(), res.GetLogStorage())
	require.Equal(t, TemplateWrappedLogRouting.GetLogStorageID(), res.GetLogStorageID())
	require.Equal(t, TemplateWrappedLogRouting.GetResourceID(), res.GetResourceID())
	require.Equal(t, TemplateWrappedLogRouting.GetVariant(), res.GetVariant())
}

func TestLogRoutingOp_Update_400(t *testing.T) {
	expected := newErrorResponse(400, "Invalid update parameters.")
	client := newTestClient(expected, http.StatusBadRequest)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	updateReq := v1.LogRouting{}
	routing, err := api.Update(ctx, "0", &updateReq)
	require.Nil(t, routing)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid")
}

func TestLogRoutingOp_Delete(t *testing.T) {
	client := newTestClient(nil, http.StatusNoContent)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	err := api.Delete(ctx, "12345")
	require.NoError(t, err)
}

func TestLogRoutingOp_Delete_400(t *testing.T) {
	expected := newErrorResponse(400, "Invalid delete request.")
	client := newTestClient(expected, http.StatusBadRequest)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	err := api.Delete(ctx, "0")
	require.Error(t, err)
	require.ErrorContains(t, err, "Bad Request")
}

func TestLogRoutingOp_Integration(t *testing.T) {
	client, err := IntegratedClient(t)
	require.NoError(t, err)
	api := NewLogRoutingOp(client)
	ctx := context.Background()

	// TODO: Create or fetch a valid Publisher and LogTable (log storage) for integration
	// For now, these should be replaced with actual resource creation or fixture helpers
	publisher := v1.Publisher{
		Code:        "test-publisher-code",
		Description: v1.NewOptString("integration test publisher"),
		Variants:    nil, // add variants if required
	}
	logStorage := v1.LogStorage{
		ID:   1,
		Name: v1.NewOptString("integration-log-storage"),
	}

	createReq := v1.LogRouting{
		ResourceID:    v1.NewOptNilInt64(1), // replace with actual resource ID
		Publisher:     publisher,
		PublisherCode: publisher.Code,
		Variant:       "default", // replace with actual variant if required
		LogStorage:    logStorage,
		LogStorageID:  v1.NewNilInt64(1), // replace with actual log storage ID
	}
	created, err := api.Create(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	defer api.Delete(ctx, strconv.FormatInt(created.GetID(), 10))

	// Read
	read, err := api.Read(ctx, strconv.FormatInt(created.GetID(), 10))
	require.NoError(t, err)
	require.NotNil(t, read)
	require.Equal(t, created.GetID(), read.GetID())

	// Update (no updatable fields in LogRouting, so just call update with same data)
	updated, err := api.Update(ctx, strconv.FormatInt(created.GetID(), 10), created)
	require.NoError(t, err)
	require.NotNil(t, updated)

	// List
	routings, err := api.List(ctx, v1.LogsRoutingsListParams{Count: v1.NewOptInt(10), From: v1.NewOptInt(0)})
	require.NoError(t, err)
	found := false
	for _, r := range routings {
		if r.GetID() == created.GetID() {
			found = true
			break
		}
	}
	require.True(t, found, "created log routing not found in list")
}
