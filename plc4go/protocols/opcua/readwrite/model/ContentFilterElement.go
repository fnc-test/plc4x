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

// ContentFilterElement is the corresponding interface of ContentFilterElement
type ContentFilterElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetFilterOperator returns FilterOperator (property field)
	GetFilterOperator() FilterOperator
	// GetNoOfFilterOperands returns NoOfFilterOperands (property field)
	GetNoOfFilterOperands() int32
	// GetFilterOperands returns FilterOperands (property field)
	GetFilterOperands() []ExtensionObject
}

// ContentFilterElementExactly can be used when we want exactly this type and not a type which fulfills ContentFilterElement.
// This is useful for switch cases.
type ContentFilterElementExactly interface {
	ContentFilterElement
	isContentFilterElement() bool
}

// _ContentFilterElement is the data-structure of this message
type _ContentFilterElement struct {
	*_ExtensionObjectDefinition
	FilterOperator     FilterOperator
	NoOfFilterOperands int32
	FilterOperands     []ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ContentFilterElement) GetIdentifier() string {
	return "585"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ContentFilterElement) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ContentFilterElement) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ContentFilterElement) GetFilterOperator() FilterOperator {
	return m.FilterOperator
}

func (m *_ContentFilterElement) GetNoOfFilterOperands() int32 {
	return m.NoOfFilterOperands
}

func (m *_ContentFilterElement) GetFilterOperands() []ExtensionObject {
	return m.FilterOperands
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewContentFilterElement factory function for _ContentFilterElement
func NewContentFilterElement(filterOperator FilterOperator, noOfFilterOperands int32, filterOperands []ExtensionObject) *_ContentFilterElement {
	_result := &_ContentFilterElement{
		FilterOperator:             filterOperator,
		NoOfFilterOperands:         noOfFilterOperands,
		FilterOperands:             filterOperands,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastContentFilterElement(structType any) ContentFilterElement {
	if casted, ok := structType.(ContentFilterElement); ok {
		return casted
	}
	if casted, ok := structType.(*ContentFilterElement); ok {
		return *casted
	}
	return nil
}

func (m *_ContentFilterElement) GetTypeName() string {
	return "ContentFilterElement"
}

func (m *_ContentFilterElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (filterOperator)
	lengthInBits += 32

	// Simple field (noOfFilterOperands)
	lengthInBits += 32

	// Array field
	if len(m.FilterOperands) > 0 {
		for _curItem, element := range m.FilterOperands {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.FilterOperands), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ContentFilterElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ContentFilterElementParse(ctx context.Context, theBytes []byte, identifier string) (ContentFilterElement, error) {
	return ContentFilterElementParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ContentFilterElementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ContentFilterElement, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ContentFilterElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContentFilterElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (filterOperator)
	if pullErr := readBuffer.PullContext("filterOperator"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for filterOperator")
	}
	_filterOperator, _filterOperatorErr := FilterOperatorParseWithBuffer(ctx, readBuffer)
	if _filterOperatorErr != nil {
		return nil, errors.Wrap(_filterOperatorErr, "Error parsing 'filterOperator' field of ContentFilterElement")
	}
	filterOperator := _filterOperator
	if closeErr := readBuffer.CloseContext("filterOperator"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for filterOperator")
	}

	// Simple Field (noOfFilterOperands)
	_noOfFilterOperands, _noOfFilterOperandsErr := readBuffer.ReadInt32("noOfFilterOperands", 32)
	if _noOfFilterOperandsErr != nil {
		return nil, errors.Wrap(_noOfFilterOperandsErr, "Error parsing 'noOfFilterOperands' field of ContentFilterElement")
	}
	noOfFilterOperands := _noOfFilterOperands

	// Array field (filterOperands)
	if pullErr := readBuffer.PullContext("filterOperands", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for filterOperands")
	}
	// Count array
	filterOperands := make([]ExtensionObject, utils.Max(noOfFilterOperands, 0))
	// This happens when the size is set conditional to 0
	if len(filterOperands) == 0 {
		filterOperands = nil
	}
	{
		_numItems := uint16(utils.Max(noOfFilterOperands, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectParseWithBuffer(arrayCtx, readBuffer, bool(true))
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'filterOperands' field of ContentFilterElement")
			}
			filterOperands[_curItem] = _item.(ExtensionObject)
		}
	}
	if closeErr := readBuffer.CloseContext("filterOperands", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for filterOperands")
	}

	if closeErr := readBuffer.CloseContext("ContentFilterElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContentFilterElement")
	}

	// Create a partially initialized instance
	_child := &_ContentFilterElement{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		FilterOperator:             filterOperator,
		NoOfFilterOperands:         noOfFilterOperands,
		FilterOperands:             filterOperands,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ContentFilterElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContentFilterElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ContentFilterElement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ContentFilterElement")
		}

		// Simple Field (filterOperator)
		if pushErr := writeBuffer.PushContext("filterOperator"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for filterOperator")
		}
		_filterOperatorErr := writeBuffer.WriteSerializable(ctx, m.GetFilterOperator())
		if popErr := writeBuffer.PopContext("filterOperator"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for filterOperator")
		}
		if _filterOperatorErr != nil {
			return errors.Wrap(_filterOperatorErr, "Error serializing 'filterOperator' field")
		}

		// Simple Field (noOfFilterOperands)
		noOfFilterOperands := int32(m.GetNoOfFilterOperands())
		_noOfFilterOperandsErr := writeBuffer.WriteInt32("noOfFilterOperands", 32, (noOfFilterOperands))
		if _noOfFilterOperandsErr != nil {
			return errors.Wrap(_noOfFilterOperandsErr, "Error serializing 'noOfFilterOperands' field")
		}

		// Array Field (filterOperands)
		if pushErr := writeBuffer.PushContext("filterOperands", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for filterOperands")
		}
		for _curItem, _element := range m.GetFilterOperands() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetFilterOperands()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'filterOperands' field")
			}
		}
		if popErr := writeBuffer.PopContext("filterOperands", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for filterOperands")
		}

		if popErr := writeBuffer.PopContext("ContentFilterElement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ContentFilterElement")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ContentFilterElement) isContentFilterElement() bool {
	return true
}

func (m *_ContentFilterElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
