// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/jcerato/device-sdk-go/internal/common"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/uuid"
)

var (
	initOnce sync.Once
)

// Init basic state for cache
func InitCache() {
	initOnce.Do(func() {
		ctx := context.WithValue(context.Background(), common.CorrelationHeader, uuid.New().String())

		vds, err := common.ValueDescriptorClient.ValueDescriptors(ctx)
		if err != nil {
			common.LoggingClient.Error(fmt.Sprintf("Value Descriptor cache initialization failed: %v", err))
			vds = make([]contract.ValueDescriptor, 0)
		}
		newValueDescriptorCache(vds)

		ds, err := common.DeviceClient.DevicesForServiceByName(common.ServiceName, ctx)
		if err != nil {
			common.LoggingClient.Error(fmt.Sprintf("Device cache initialization failed: %v", err))
			ds = make([]contract.Device, 0)
		}
		newDeviceCache(ds)

		dps := make([]contract.DeviceProfile, len(ds))
		for i, d := range ds {
			dps[i] = d.Profile
		}
		newProfileCache(dps)
	})
}
