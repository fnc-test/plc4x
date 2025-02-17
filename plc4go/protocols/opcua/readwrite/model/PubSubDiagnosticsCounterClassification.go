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

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// PubSubDiagnosticsCounterClassification is an enum
type PubSubDiagnosticsCounterClassification uint32

type IPubSubDiagnosticsCounterClassification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationInformation PubSubDiagnosticsCounterClassification = 0
	PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationError       PubSubDiagnosticsCounterClassification = 1
)

var PubSubDiagnosticsCounterClassificationValues []PubSubDiagnosticsCounterClassification

func init() {
	_ = errors.New
	PubSubDiagnosticsCounterClassificationValues = []PubSubDiagnosticsCounterClassification{
		PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationInformation,
		PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationError,
	}
}

func PubSubDiagnosticsCounterClassificationByValue(value uint32) (enum PubSubDiagnosticsCounterClassification, ok bool) {
	switch value {
	case 0:
		return PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationInformation, true
	case 1:
		return PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationError, true
	}
	return 0, false
}

func PubSubDiagnosticsCounterClassificationByName(value string) (enum PubSubDiagnosticsCounterClassification, ok bool) {
	switch value {
	case "pubSubDiagnosticsCounterClassificationInformation":
		return PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationInformation, true
	case "pubSubDiagnosticsCounterClassificationError":
		return PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationError, true
	}
	return 0, false
}

func PubSubDiagnosticsCounterClassificationKnows(value uint32) bool {
	for _, typeValue := range PubSubDiagnosticsCounterClassificationValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPubSubDiagnosticsCounterClassification(structType any) PubSubDiagnosticsCounterClassification {
	castFunc := func(typ any) PubSubDiagnosticsCounterClassification {
		if sPubSubDiagnosticsCounterClassification, ok := typ.(PubSubDiagnosticsCounterClassification); ok {
			return sPubSubDiagnosticsCounterClassification
		}
		return 0
	}
	return castFunc(structType)
}

func (m PubSubDiagnosticsCounterClassification) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m PubSubDiagnosticsCounterClassification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PubSubDiagnosticsCounterClassificationParse(ctx context.Context, theBytes []byte) (PubSubDiagnosticsCounterClassification, error) {
	return PubSubDiagnosticsCounterClassificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PubSubDiagnosticsCounterClassificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PubSubDiagnosticsCounterClassification, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("PubSubDiagnosticsCounterClassification", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PubSubDiagnosticsCounterClassification")
	}
	if enum, ok := PubSubDiagnosticsCounterClassificationByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PubSubDiagnosticsCounterClassification")
		return PubSubDiagnosticsCounterClassification(val), nil
	} else {
		return enum, nil
	}
}

func (e PubSubDiagnosticsCounterClassification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PubSubDiagnosticsCounterClassification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("PubSubDiagnosticsCounterClassification", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PubSubDiagnosticsCounterClassification) PLC4XEnumName() string {
	switch e {
	case PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationInformation:
		return "pubSubDiagnosticsCounterClassificationInformation"
	case PubSubDiagnosticsCounterClassification_pubSubDiagnosticsCounterClassificationError:
		return "pubSubDiagnosticsCounterClassificationError"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e PubSubDiagnosticsCounterClassification) String() string {
	return e.PLC4XEnumName()
}
