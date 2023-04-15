/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package cbus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	spiModel "github.com/apache/plc4x/plc4go/spi/model"
)

func TestNewSubscriptionHandle(t *testing.T) {
	type args struct {
		subscriber       *Subscriber
		tagName          string
		tag              any
		subscriptionType spiModel.SubscriptionType
		interval         time.Duration
	}
	tests := []struct {
		name string
		args args
		want *SubscriptionHandle
	}{
		{
			name: "empty",
			want: func() *SubscriptionHandle {
				subscriptionHandle := SubscriptionHandle{}
				subscriptionHandle.DefaultPlcSubscriptionHandle = spiModel.NewDefaultPlcSubscriptionHandle(nil)
				return &subscriptionHandle
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewSubscriptionHandle(tt.args.subscriber, tt.args.tagName, tt.args.tag, tt.args.subscriptionType, tt.args.interval), "NewSubscriptionHandle(%v, %v, %v, %v, %v)", tt.args.subscriber, tt.args.tagName, tt.args.tag, tt.args.subscriptionType, tt.args.interval)
		})
	}
}
