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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// DiagnosticInfo is the corresponding interface of DiagnosticInfo
type DiagnosticInfo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetInnerDiagnosticInfoSpecified returns InnerDiagnosticInfoSpecified (property field)
	GetInnerDiagnosticInfoSpecified() bool
	// GetInnerStatusCodeSpecified returns InnerStatusCodeSpecified (property field)
	GetInnerStatusCodeSpecified() bool
	// GetAdditionalInfoSpecified returns AdditionalInfoSpecified (property field)
	GetAdditionalInfoSpecified() bool
	// GetLocaleSpecified returns LocaleSpecified (property field)
	GetLocaleSpecified() bool
	// GetLocalizedTextSpecified returns LocalizedTextSpecified (property field)
	GetLocalizedTextSpecified() bool
	// GetNamespaceURISpecified returns NamespaceURISpecified (property field)
	GetNamespaceURISpecified() bool
	// GetSymbolicIdSpecified returns SymbolicIdSpecified (property field)
	GetSymbolicIdSpecified() bool
	// GetSymbolicId returns SymbolicId (property field)
	GetSymbolicId() *int32
	// GetNamespaceURI returns NamespaceURI (property field)
	GetNamespaceURI() *int32
	// GetLocale returns Locale (property field)
	GetLocale() *int32
	// GetLocalizedText returns LocalizedText (property field)
	GetLocalizedText() *int32
	// GetAdditionalInfo returns AdditionalInfo (property field)
	GetAdditionalInfo() PascalString
	// GetInnerStatusCode returns InnerStatusCode (property field)
	GetInnerStatusCode() StatusCode
	// GetInnerDiagnosticInfo returns InnerDiagnosticInfo (property field)
	GetInnerDiagnosticInfo() DiagnosticInfo
}

// DiagnosticInfoExactly can be used when we want exactly this type and not a type which fulfills DiagnosticInfo.
// This is useful for switch cases.
type DiagnosticInfoExactly interface {
	DiagnosticInfo
	isDiagnosticInfo() bool
}

// _DiagnosticInfo is the data-structure of this message
type _DiagnosticInfo struct {
	InnerDiagnosticInfoSpecified bool
	InnerStatusCodeSpecified     bool
	AdditionalInfoSpecified      bool
	LocaleSpecified              bool
	LocalizedTextSpecified       bool
	NamespaceURISpecified        bool
	SymbolicIdSpecified          bool
	SymbolicId                   *int32
	NamespaceURI                 *int32
	Locale                       *int32
	LocalizedText                *int32
	AdditionalInfo               PascalString
	InnerStatusCode              StatusCode
	InnerDiagnosticInfo          DiagnosticInfo
	// Reserved Fields
	reservedField0 *bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DiagnosticInfo) GetInnerDiagnosticInfoSpecified() bool {
	return m.InnerDiagnosticInfoSpecified
}

func (m *_DiagnosticInfo) GetInnerStatusCodeSpecified() bool {
	return m.InnerStatusCodeSpecified
}

func (m *_DiagnosticInfo) GetAdditionalInfoSpecified() bool {
	return m.AdditionalInfoSpecified
}

func (m *_DiagnosticInfo) GetLocaleSpecified() bool {
	return m.LocaleSpecified
}

func (m *_DiagnosticInfo) GetLocalizedTextSpecified() bool {
	return m.LocalizedTextSpecified
}

func (m *_DiagnosticInfo) GetNamespaceURISpecified() bool {
	return m.NamespaceURISpecified
}

func (m *_DiagnosticInfo) GetSymbolicIdSpecified() bool {
	return m.SymbolicIdSpecified
}

func (m *_DiagnosticInfo) GetSymbolicId() *int32 {
	return m.SymbolicId
}

func (m *_DiagnosticInfo) GetNamespaceURI() *int32 {
	return m.NamespaceURI
}

func (m *_DiagnosticInfo) GetLocale() *int32 {
	return m.Locale
}

func (m *_DiagnosticInfo) GetLocalizedText() *int32 {
	return m.LocalizedText
}

func (m *_DiagnosticInfo) GetAdditionalInfo() PascalString {
	return m.AdditionalInfo
}

func (m *_DiagnosticInfo) GetInnerStatusCode() StatusCode {
	return m.InnerStatusCode
}

func (m *_DiagnosticInfo) GetInnerDiagnosticInfo() DiagnosticInfo {
	return m.InnerDiagnosticInfo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDiagnosticInfo factory function for _DiagnosticInfo
func NewDiagnosticInfo(innerDiagnosticInfoSpecified bool, innerStatusCodeSpecified bool, additionalInfoSpecified bool, localeSpecified bool, localizedTextSpecified bool, namespaceURISpecified bool, symbolicIdSpecified bool, symbolicId *int32, namespaceURI *int32, locale *int32, localizedText *int32, additionalInfo PascalString, innerStatusCode StatusCode, innerDiagnosticInfo DiagnosticInfo) *_DiagnosticInfo {
	return &_DiagnosticInfo{InnerDiagnosticInfoSpecified: innerDiagnosticInfoSpecified, InnerStatusCodeSpecified: innerStatusCodeSpecified, AdditionalInfoSpecified: additionalInfoSpecified, LocaleSpecified: localeSpecified, LocalizedTextSpecified: localizedTextSpecified, NamespaceURISpecified: namespaceURISpecified, SymbolicIdSpecified: symbolicIdSpecified, SymbolicId: symbolicId, NamespaceURI: namespaceURI, Locale: locale, LocalizedText: localizedText, AdditionalInfo: additionalInfo, InnerStatusCode: innerStatusCode, InnerDiagnosticInfo: innerDiagnosticInfo}
}

// Deprecated: use the interface for direct cast
func CastDiagnosticInfo(structType any) DiagnosticInfo {
	if casted, ok := structType.(DiagnosticInfo); ok {
		return casted
	}
	if casted, ok := structType.(*DiagnosticInfo); ok {
		return *casted
	}
	return nil
}

func (m *_DiagnosticInfo) GetTypeName() string {
	return "DiagnosticInfo"
}

func (m *_DiagnosticInfo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (innerDiagnosticInfoSpecified)
	lengthInBits += 1

	// Simple field (innerStatusCodeSpecified)
	lengthInBits += 1

	// Simple field (additionalInfoSpecified)
	lengthInBits += 1

	// Simple field (localeSpecified)
	lengthInBits += 1

	// Simple field (localizedTextSpecified)
	lengthInBits += 1

	// Simple field (namespaceURISpecified)
	lengthInBits += 1

	// Simple field (symbolicIdSpecified)
	lengthInBits += 1

	// Optional Field (symbolicId)
	if m.SymbolicId != nil {
		lengthInBits += 32
	}

	// Optional Field (namespaceURI)
	if m.NamespaceURI != nil {
		lengthInBits += 32
	}

	// Optional Field (locale)
	if m.Locale != nil {
		lengthInBits += 32
	}

	// Optional Field (localizedText)
	if m.LocalizedText != nil {
		lengthInBits += 32
	}

	// Optional Field (additionalInfo)
	if m.AdditionalInfo != nil {
		lengthInBits += m.AdditionalInfo.GetLengthInBits(ctx)
	}

	// Optional Field (innerStatusCode)
	if m.InnerStatusCode != nil {
		lengthInBits += m.InnerStatusCode.GetLengthInBits(ctx)
	}

	// Optional Field (innerDiagnosticInfo)
	if m.InnerDiagnosticInfo != nil {
		lengthInBits += m.InnerDiagnosticInfo.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_DiagnosticInfo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DiagnosticInfoParse(ctx context.Context, theBytes []byte) (DiagnosticInfo, error) {
	return DiagnosticInfoParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DiagnosticInfoParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DiagnosticInfo, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DiagnosticInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DiagnosticInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DiagnosticInfo")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (innerDiagnosticInfoSpecified)
	_innerDiagnosticInfoSpecified, _innerDiagnosticInfoSpecifiedErr := readBuffer.ReadBit("innerDiagnosticInfoSpecified")
	if _innerDiagnosticInfoSpecifiedErr != nil {
		return nil, errors.Wrap(_innerDiagnosticInfoSpecifiedErr, "Error parsing 'innerDiagnosticInfoSpecified' field of DiagnosticInfo")
	}
	innerDiagnosticInfoSpecified := _innerDiagnosticInfoSpecified

	// Simple Field (innerStatusCodeSpecified)
	_innerStatusCodeSpecified, _innerStatusCodeSpecifiedErr := readBuffer.ReadBit("innerStatusCodeSpecified")
	if _innerStatusCodeSpecifiedErr != nil {
		return nil, errors.Wrap(_innerStatusCodeSpecifiedErr, "Error parsing 'innerStatusCodeSpecified' field of DiagnosticInfo")
	}
	innerStatusCodeSpecified := _innerStatusCodeSpecified

	// Simple Field (additionalInfoSpecified)
	_additionalInfoSpecified, _additionalInfoSpecifiedErr := readBuffer.ReadBit("additionalInfoSpecified")
	if _additionalInfoSpecifiedErr != nil {
		return nil, errors.Wrap(_additionalInfoSpecifiedErr, "Error parsing 'additionalInfoSpecified' field of DiagnosticInfo")
	}
	additionalInfoSpecified := _additionalInfoSpecified

	// Simple Field (localeSpecified)
	_localeSpecified, _localeSpecifiedErr := readBuffer.ReadBit("localeSpecified")
	if _localeSpecifiedErr != nil {
		return nil, errors.Wrap(_localeSpecifiedErr, "Error parsing 'localeSpecified' field of DiagnosticInfo")
	}
	localeSpecified := _localeSpecified

	// Simple Field (localizedTextSpecified)
	_localizedTextSpecified, _localizedTextSpecifiedErr := readBuffer.ReadBit("localizedTextSpecified")
	if _localizedTextSpecifiedErr != nil {
		return nil, errors.Wrap(_localizedTextSpecifiedErr, "Error parsing 'localizedTextSpecified' field of DiagnosticInfo")
	}
	localizedTextSpecified := _localizedTextSpecified

	// Simple Field (namespaceURISpecified)
	_namespaceURISpecified, _namespaceURISpecifiedErr := readBuffer.ReadBit("namespaceURISpecified")
	if _namespaceURISpecifiedErr != nil {
		return nil, errors.Wrap(_namespaceURISpecifiedErr, "Error parsing 'namespaceURISpecified' field of DiagnosticInfo")
	}
	namespaceURISpecified := _namespaceURISpecified

	// Simple Field (symbolicIdSpecified)
	_symbolicIdSpecified, _symbolicIdSpecifiedErr := readBuffer.ReadBit("symbolicIdSpecified")
	if _symbolicIdSpecifiedErr != nil {
		return nil, errors.Wrap(_symbolicIdSpecifiedErr, "Error parsing 'symbolicIdSpecified' field of DiagnosticInfo")
	}
	symbolicIdSpecified := _symbolicIdSpecified

	// Optional Field (symbolicId) (Can be skipped, if a given expression evaluates to false)
	var symbolicId *int32 = nil
	if symbolicIdSpecified {
		_val, _err := readBuffer.ReadInt32("symbolicId", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'symbolicId' field of DiagnosticInfo")
		}
		symbolicId = &_val
	}

	// Optional Field (namespaceURI) (Can be skipped, if a given expression evaluates to false)
	var namespaceURI *int32 = nil
	if namespaceURISpecified {
		_val, _err := readBuffer.ReadInt32("namespaceURI", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'namespaceURI' field of DiagnosticInfo")
		}
		namespaceURI = &_val
	}

	// Optional Field (locale) (Can be skipped, if a given expression evaluates to false)
	var locale *int32 = nil
	if localeSpecified {
		_val, _err := readBuffer.ReadInt32("locale", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'locale' field of DiagnosticInfo")
		}
		locale = &_val
	}

	// Optional Field (localizedText) (Can be skipped, if a given expression evaluates to false)
	var localizedText *int32 = nil
	if localizedTextSpecified {
		_val, _err := readBuffer.ReadInt32("localizedText", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'localizedText' field of DiagnosticInfo")
		}
		localizedText = &_val
	}

	// Optional Field (additionalInfo) (Can be skipped, if a given expression evaluates to false)
	var additionalInfo PascalString = nil
	if additionalInfoSpecified {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("additionalInfo"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for additionalInfo")
		}
		_val, _err := PascalStringParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'additionalInfo' field of DiagnosticInfo")
		default:
			additionalInfo = _val.(PascalString)
			if closeErr := readBuffer.CloseContext("additionalInfo"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for additionalInfo")
			}
		}
	}

	// Optional Field (innerStatusCode) (Can be skipped, if a given expression evaluates to false)
	var innerStatusCode StatusCode = nil
	if innerStatusCodeSpecified {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("innerStatusCode"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for innerStatusCode")
		}
		_val, _err := StatusCodeParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'innerStatusCode' field of DiagnosticInfo")
		default:
			innerStatusCode = _val.(StatusCode)
			if closeErr := readBuffer.CloseContext("innerStatusCode"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for innerStatusCode")
			}
		}
	}

	// Optional Field (innerDiagnosticInfo) (Can be skipped, if a given expression evaluates to false)
	var innerDiagnosticInfo DiagnosticInfo = nil
	if innerDiagnosticInfoSpecified {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("innerDiagnosticInfo"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for innerDiagnosticInfo")
		}
		_val, _err := DiagnosticInfoParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'innerDiagnosticInfo' field of DiagnosticInfo")
		default:
			innerDiagnosticInfo = _val.(DiagnosticInfo)
			if closeErr := readBuffer.CloseContext("innerDiagnosticInfo"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for innerDiagnosticInfo")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("DiagnosticInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DiagnosticInfo")
	}

	// Create the instance
	return &_DiagnosticInfo{
		InnerDiagnosticInfoSpecified: innerDiagnosticInfoSpecified,
		InnerStatusCodeSpecified:     innerStatusCodeSpecified,
		AdditionalInfoSpecified:      additionalInfoSpecified,
		LocaleSpecified:              localeSpecified,
		LocalizedTextSpecified:       localizedTextSpecified,
		NamespaceURISpecified:        namespaceURISpecified,
		SymbolicIdSpecified:          symbolicIdSpecified,
		SymbolicId:                   symbolicId,
		NamespaceURI:                 namespaceURI,
		Locale:                       locale,
		LocalizedText:                localizedText,
		AdditionalInfo:               additionalInfo,
		InnerStatusCode:              innerStatusCode,
		InnerDiagnosticInfo:          innerDiagnosticInfo,
		reservedField0:               reservedField0,
	}, nil
}

func (m *_DiagnosticInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DiagnosticInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DiagnosticInfo"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DiagnosticInfo")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (innerDiagnosticInfoSpecified)
	innerDiagnosticInfoSpecified := bool(m.GetInnerDiagnosticInfoSpecified())
	_innerDiagnosticInfoSpecifiedErr := writeBuffer.WriteBit("innerDiagnosticInfoSpecified", (innerDiagnosticInfoSpecified))
	if _innerDiagnosticInfoSpecifiedErr != nil {
		return errors.Wrap(_innerDiagnosticInfoSpecifiedErr, "Error serializing 'innerDiagnosticInfoSpecified' field")
	}

	// Simple Field (innerStatusCodeSpecified)
	innerStatusCodeSpecified := bool(m.GetInnerStatusCodeSpecified())
	_innerStatusCodeSpecifiedErr := writeBuffer.WriteBit("innerStatusCodeSpecified", (innerStatusCodeSpecified))
	if _innerStatusCodeSpecifiedErr != nil {
		return errors.Wrap(_innerStatusCodeSpecifiedErr, "Error serializing 'innerStatusCodeSpecified' field")
	}

	// Simple Field (additionalInfoSpecified)
	additionalInfoSpecified := bool(m.GetAdditionalInfoSpecified())
	_additionalInfoSpecifiedErr := writeBuffer.WriteBit("additionalInfoSpecified", (additionalInfoSpecified))
	if _additionalInfoSpecifiedErr != nil {
		return errors.Wrap(_additionalInfoSpecifiedErr, "Error serializing 'additionalInfoSpecified' field")
	}

	// Simple Field (localeSpecified)
	localeSpecified := bool(m.GetLocaleSpecified())
	_localeSpecifiedErr := writeBuffer.WriteBit("localeSpecified", (localeSpecified))
	if _localeSpecifiedErr != nil {
		return errors.Wrap(_localeSpecifiedErr, "Error serializing 'localeSpecified' field")
	}

	// Simple Field (localizedTextSpecified)
	localizedTextSpecified := bool(m.GetLocalizedTextSpecified())
	_localizedTextSpecifiedErr := writeBuffer.WriteBit("localizedTextSpecified", (localizedTextSpecified))
	if _localizedTextSpecifiedErr != nil {
		return errors.Wrap(_localizedTextSpecifiedErr, "Error serializing 'localizedTextSpecified' field")
	}

	// Simple Field (namespaceURISpecified)
	namespaceURISpecified := bool(m.GetNamespaceURISpecified())
	_namespaceURISpecifiedErr := writeBuffer.WriteBit("namespaceURISpecified", (namespaceURISpecified))
	if _namespaceURISpecifiedErr != nil {
		return errors.Wrap(_namespaceURISpecifiedErr, "Error serializing 'namespaceURISpecified' field")
	}

	// Simple Field (symbolicIdSpecified)
	symbolicIdSpecified := bool(m.GetSymbolicIdSpecified())
	_symbolicIdSpecifiedErr := writeBuffer.WriteBit("symbolicIdSpecified", (symbolicIdSpecified))
	if _symbolicIdSpecifiedErr != nil {
		return errors.Wrap(_symbolicIdSpecifiedErr, "Error serializing 'symbolicIdSpecified' field")
	}

	// Optional Field (symbolicId) (Can be skipped, if the value is null)
	var symbolicId *int32 = nil
	if m.GetSymbolicId() != nil {
		symbolicId = m.GetSymbolicId()
		_symbolicIdErr := writeBuffer.WriteInt32("symbolicId", 32, *(symbolicId))
		if _symbolicIdErr != nil {
			return errors.Wrap(_symbolicIdErr, "Error serializing 'symbolicId' field")
		}
	}

	// Optional Field (namespaceURI) (Can be skipped, if the value is null)
	var namespaceURI *int32 = nil
	if m.GetNamespaceURI() != nil {
		namespaceURI = m.GetNamespaceURI()
		_namespaceURIErr := writeBuffer.WriteInt32("namespaceURI", 32, *(namespaceURI))
		if _namespaceURIErr != nil {
			return errors.Wrap(_namespaceURIErr, "Error serializing 'namespaceURI' field")
		}
	}

	// Optional Field (locale) (Can be skipped, if the value is null)
	var locale *int32 = nil
	if m.GetLocale() != nil {
		locale = m.GetLocale()
		_localeErr := writeBuffer.WriteInt32("locale", 32, *(locale))
		if _localeErr != nil {
			return errors.Wrap(_localeErr, "Error serializing 'locale' field")
		}
	}

	// Optional Field (localizedText) (Can be skipped, if the value is null)
	var localizedText *int32 = nil
	if m.GetLocalizedText() != nil {
		localizedText = m.GetLocalizedText()
		_localizedTextErr := writeBuffer.WriteInt32("localizedText", 32, *(localizedText))
		if _localizedTextErr != nil {
			return errors.Wrap(_localizedTextErr, "Error serializing 'localizedText' field")
		}
	}

	// Optional Field (additionalInfo) (Can be skipped, if the value is null)
	var additionalInfo PascalString = nil
	if m.GetAdditionalInfo() != nil {
		if pushErr := writeBuffer.PushContext("additionalInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for additionalInfo")
		}
		additionalInfo = m.GetAdditionalInfo()
		_additionalInfoErr := writeBuffer.WriteSerializable(ctx, additionalInfo)
		if popErr := writeBuffer.PopContext("additionalInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for additionalInfo")
		}
		if _additionalInfoErr != nil {
			return errors.Wrap(_additionalInfoErr, "Error serializing 'additionalInfo' field")
		}
	}

	// Optional Field (innerStatusCode) (Can be skipped, if the value is null)
	var innerStatusCode StatusCode = nil
	if m.GetInnerStatusCode() != nil {
		if pushErr := writeBuffer.PushContext("innerStatusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerStatusCode")
		}
		innerStatusCode = m.GetInnerStatusCode()
		_innerStatusCodeErr := writeBuffer.WriteSerializable(ctx, innerStatusCode)
		if popErr := writeBuffer.PopContext("innerStatusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerStatusCode")
		}
		if _innerStatusCodeErr != nil {
			return errors.Wrap(_innerStatusCodeErr, "Error serializing 'innerStatusCode' field")
		}
	}

	// Optional Field (innerDiagnosticInfo) (Can be skipped, if the value is null)
	var innerDiagnosticInfo DiagnosticInfo = nil
	if m.GetInnerDiagnosticInfo() != nil {
		if pushErr := writeBuffer.PushContext("innerDiagnosticInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerDiagnosticInfo")
		}
		innerDiagnosticInfo = m.GetInnerDiagnosticInfo()
		_innerDiagnosticInfoErr := writeBuffer.WriteSerializable(ctx, innerDiagnosticInfo)
		if popErr := writeBuffer.PopContext("innerDiagnosticInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerDiagnosticInfo")
		}
		if _innerDiagnosticInfoErr != nil {
			return errors.Wrap(_innerDiagnosticInfoErr, "Error serializing 'innerDiagnosticInfo' field")
		}
	}

	if popErr := writeBuffer.PopContext("DiagnosticInfo"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DiagnosticInfo")
	}
	return nil
}

func (m *_DiagnosticInfo) isDiagnosticInfo() bool {
	return true
}

func (m *_DiagnosticInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
