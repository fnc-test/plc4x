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

// OpcuaNodeIdServicesVariableMulti is an enum
type OpcuaNodeIdServicesVariableMulti int32

type IOpcuaNodeIdServicesVariableMulti interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_EnumStrings                                 OpcuaNodeIdServicesVariableMulti = 2377
	OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_Definition                                  OpcuaNodeIdServicesVariableMulti = 3780
	OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_ValuePrecision                              OpcuaNodeIdServicesVariableMulti = 3781
	OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_Definition                             OpcuaNodeIdServicesVariableMulti = 11239
	OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValuePrecision                         OpcuaNodeIdServicesVariableMulti = 11240
	OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_EnumValues                             OpcuaNodeIdServicesVariableMulti = 11241
	OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValueAsText                            OpcuaNodeIdServicesVariableMulti = 11461
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_Definition               OpcuaNodeIdServicesVariableMulti = 19078
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValuePrecision           OpcuaNodeIdServicesVariableMulti = 19079
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumValues               OpcuaNodeIdServicesVariableMulti = 19080
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsText              OpcuaNodeIdServicesVariableMulti = 19081
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumDictionaryEntries    OpcuaNodeIdServicesVariableMulti = 19082
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsDictionaryEntries OpcuaNodeIdServicesVariableMulti = 19083
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_Definition                   OpcuaNodeIdServicesVariableMulti = 19085
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValuePrecision               OpcuaNodeIdServicesVariableMulti = 19086
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumValues                   OpcuaNodeIdServicesVariableMulti = 19087
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsText                  OpcuaNodeIdServicesVariableMulti = 19088
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumDictionaryEntries        OpcuaNodeIdServicesVariableMulti = 19089
	OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsDictionaryEntries     OpcuaNodeIdServicesVariableMulti = 19090
)

var OpcuaNodeIdServicesVariableMultiValues []OpcuaNodeIdServicesVariableMulti

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableMultiValues = []OpcuaNodeIdServicesVariableMulti{
		OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_EnumStrings,
		OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_Definition,
		OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_ValuePrecision,
		OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_Definition,
		OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValuePrecision,
		OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_EnumValues,
		OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValueAsText,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_Definition,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValuePrecision,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumValues,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsText,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumDictionaryEntries,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsDictionaryEntries,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_Definition,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValuePrecision,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumValues,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsText,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumDictionaryEntries,
		OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsDictionaryEntries,
	}
}

func OpcuaNodeIdServicesVariableMultiByValue(value int32) (enum OpcuaNodeIdServicesVariableMulti, ok bool) {
	switch value {
	case 11239:
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_Definition, true
	case 11240:
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValuePrecision, true
	case 11241:
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_EnumValues, true
	case 11461:
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValueAsText, true
	case 19078:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_Definition, true
	case 19079:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValuePrecision, true
	case 19080:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumValues, true
	case 19081:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsText, true
	case 19082:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumDictionaryEntries, true
	case 19083:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsDictionaryEntries, true
	case 19085:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_Definition, true
	case 19086:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValuePrecision, true
	case 19087:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumValues, true
	case 19088:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsText, true
	case 19089:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumDictionaryEntries, true
	case 19090:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsDictionaryEntries, true
	case 2377:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_EnumStrings, true
	case 3780:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_Definition, true
	case 3781:
		return OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_ValuePrecision, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableMultiByName(value string) (enum OpcuaNodeIdServicesVariableMulti, ok bool) {
	switch value {
	case "MultiStateValueDiscreteType_Definition":
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_Definition, true
	case "MultiStateValueDiscreteType_ValuePrecision":
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValuePrecision, true
	case "MultiStateValueDiscreteType_EnumValues":
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_EnumValues, true
	case "MultiStateValueDiscreteType_ValueAsText":
		return OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValueAsText, true
	case "MultiStateDictionaryEntryDiscreteBaseType_Definition":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_Definition, true
	case "MultiStateDictionaryEntryDiscreteBaseType_ValuePrecision":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValuePrecision, true
	case "MultiStateDictionaryEntryDiscreteBaseType_EnumValues":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumValues, true
	case "MultiStateDictionaryEntryDiscreteBaseType_ValueAsText":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsText, true
	case "MultiStateDictionaryEntryDiscreteBaseType_EnumDictionaryEntries":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumDictionaryEntries, true
	case "MultiStateDictionaryEntryDiscreteBaseType_ValueAsDictionaryEntries":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsDictionaryEntries, true
	case "MultiStateDictionaryEntryDiscreteType_Definition":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_Definition, true
	case "MultiStateDictionaryEntryDiscreteType_ValuePrecision":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValuePrecision, true
	case "MultiStateDictionaryEntryDiscreteType_EnumValues":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumValues, true
	case "MultiStateDictionaryEntryDiscreteType_ValueAsText":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsText, true
	case "MultiStateDictionaryEntryDiscreteType_EnumDictionaryEntries":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumDictionaryEntries, true
	case "MultiStateDictionaryEntryDiscreteType_ValueAsDictionaryEntries":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsDictionaryEntries, true
	case "MultiStateDiscreteType_EnumStrings":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_EnumStrings, true
	case "MultiStateDiscreteType_Definition":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_Definition, true
	case "MultiStateDiscreteType_ValuePrecision":
		return OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_ValuePrecision, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableMultiKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableMultiValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableMulti(structType any) OpcuaNodeIdServicesVariableMulti {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableMulti {
		if sOpcuaNodeIdServicesVariableMulti, ok := typ.(OpcuaNodeIdServicesVariableMulti); ok {
			return sOpcuaNodeIdServicesVariableMulti
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableMulti) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableMulti) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableMultiParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableMulti, error) {
	return OpcuaNodeIdServicesVariableMultiParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableMultiParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableMulti, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableMulti", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableMulti")
	}
	if enum, ok := OpcuaNodeIdServicesVariableMultiByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableMulti")
		return OpcuaNodeIdServicesVariableMulti(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableMulti) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableMulti) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableMulti", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableMulti) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_Definition:
		return "MultiStateValueDiscreteType_Definition"
	case OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValuePrecision:
		return "MultiStateValueDiscreteType_ValuePrecision"
	case OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_EnumValues:
		return "MultiStateValueDiscreteType_EnumValues"
	case OpcuaNodeIdServicesVariableMulti_MultiStateValueDiscreteType_ValueAsText:
		return "MultiStateValueDiscreteType_ValueAsText"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_Definition:
		return "MultiStateDictionaryEntryDiscreteBaseType_Definition"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValuePrecision:
		return "MultiStateDictionaryEntryDiscreteBaseType_ValuePrecision"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumValues:
		return "MultiStateDictionaryEntryDiscreteBaseType_EnumValues"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsText:
		return "MultiStateDictionaryEntryDiscreteBaseType_ValueAsText"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_EnumDictionaryEntries:
		return "MultiStateDictionaryEntryDiscreteBaseType_EnumDictionaryEntries"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteBaseType_ValueAsDictionaryEntries:
		return "MultiStateDictionaryEntryDiscreteBaseType_ValueAsDictionaryEntries"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_Definition:
		return "MultiStateDictionaryEntryDiscreteType_Definition"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValuePrecision:
		return "MultiStateDictionaryEntryDiscreteType_ValuePrecision"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumValues:
		return "MultiStateDictionaryEntryDiscreteType_EnumValues"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsText:
		return "MultiStateDictionaryEntryDiscreteType_ValueAsText"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_EnumDictionaryEntries:
		return "MultiStateDictionaryEntryDiscreteType_EnumDictionaryEntries"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDictionaryEntryDiscreteType_ValueAsDictionaryEntries:
		return "MultiStateDictionaryEntryDiscreteType_ValueAsDictionaryEntries"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_EnumStrings:
		return "MultiStateDiscreteType_EnumStrings"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_Definition:
		return "MultiStateDiscreteType_Definition"
	case OpcuaNodeIdServicesVariableMulti_MultiStateDiscreteType_ValuePrecision:
		return "MultiStateDiscreteType_ValuePrecision"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableMulti) String() string {
	return e.PLC4XEnumName()
}
