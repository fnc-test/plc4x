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

// CloseSessionResponse is the corresponding interface of CloseSessionResponse
type CloseSessionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
}

// CloseSessionResponseExactly can be used when we want exactly this type and not a type which fulfills CloseSessionResponse.
// This is useful for switch cases.
type CloseSessionResponseExactly interface {
	CloseSessionResponse
	isCloseSessionResponse() bool
}

// _CloseSessionResponse is the data-structure of this message
type _CloseSessionResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSessionResponse) GetIdentifier() string {
	return "476"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSessionResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CloseSessionResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSessionResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCloseSessionResponse factory function for _CloseSessionResponse
func NewCloseSessionResponse(responseHeader ExtensionObjectDefinition) *_CloseSessionResponse {
	_result := &_CloseSessionResponse{
		ResponseHeader:             responseHeader,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCloseSessionResponse(structType any) CloseSessionResponse {
	if casted, ok := structType.(CloseSessionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSessionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSessionResponse) GetTypeName() string {
	return "CloseSessionResponse"
}

func (m *_CloseSessionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CloseSessionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CloseSessionResponseParse(ctx context.Context, theBytes []byte, identifier string) (CloseSessionResponse, error) {
	return CloseSessionResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CloseSessionResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CloseSessionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CloseSessionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSessionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of CloseSessionResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	if closeErr := readBuffer.CloseContext("CloseSessionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSessionResponse")
	}

	// Create a partially initialized instance
	_child := &_CloseSessionResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CloseSessionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSessionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSessionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSessionResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		if popErr := writeBuffer.PopContext("CloseSessionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSessionResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSessionResponse) isCloseSessionResponse() bool {
	return true
}

func (m *_CloseSessionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
