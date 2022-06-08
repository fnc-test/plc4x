/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetRecipientDevice is the data-structure of this message
type BACnetRecipientDevice struct {
	*BACnetRecipient
	DeviceValue *BACnetContextTagObjectIdentifier
}

// IBACnetRecipientDevice is the corresponding interface of BACnetRecipientDevice
type IBACnetRecipientDevice interface {
	IBACnetRecipient
	// GetDeviceValue returns DeviceValue (property field)
	GetDeviceValue() *BACnetContextTagObjectIdentifier
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetRecipientDevice) InitializeParent(parent *BACnetRecipient, peekedTagHeader *BACnetTagHeader) {
	m.BACnetRecipient.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetRecipientDevice) GetParent() *BACnetRecipient {
	return m.BACnetRecipient
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetRecipientDevice) GetDeviceValue() *BACnetContextTagObjectIdentifier {
	return m.DeviceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRecipientDevice factory function for BACnetRecipientDevice
func NewBACnetRecipientDevice(deviceValue *BACnetContextTagObjectIdentifier, peekedTagHeader *BACnetTagHeader) *BACnetRecipientDevice {
	_result := &BACnetRecipientDevice{
		DeviceValue:     deviceValue,
		BACnetRecipient: NewBACnetRecipient(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetRecipientDevice(structType interface{}) *BACnetRecipientDevice {
	if casted, ok := structType.(BACnetRecipientDevice); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetRecipientDevice); ok {
		return casted
	}
	if casted, ok := structType.(BACnetRecipient); ok {
		return CastBACnetRecipientDevice(casted.Child)
	}
	if casted, ok := structType.(*BACnetRecipient); ok {
		return CastBACnetRecipientDevice(casted.Child)
	}
	return nil
}

func (m *BACnetRecipientDevice) GetTypeName() string {
	return "BACnetRecipientDevice"
}

func (m *BACnetRecipientDevice) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetRecipientDevice) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deviceValue)
	lengthInBits += m.DeviceValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetRecipientDevice) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetRecipientDeviceParse(readBuffer utils.ReadBuffer) (*BACnetRecipientDevice, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRecipientDevice"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deviceValue)
	if pullErr := readBuffer.PullContext("deviceValue"); pullErr != nil {
		return nil, pullErr
	}
	_deviceValue, _deviceValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _deviceValueErr != nil {
		return nil, errors.Wrap(_deviceValueErr, "Error parsing 'deviceValue' field")
	}
	deviceValue := CastBACnetContextTagObjectIdentifier(_deviceValue)
	if closeErr := readBuffer.CloseContext("deviceValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientDevice"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetRecipientDevice{
		DeviceValue:     CastBACnetContextTagObjectIdentifier(deviceValue),
		BACnetRecipient: &BACnetRecipient{},
	}
	_child.BACnetRecipient.Child = _child
	return _child, nil
}

func (m *BACnetRecipientDevice) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetRecipientDevice"); pushErr != nil {
			return pushErr
		}

		// Simple Field (deviceValue)
		if pushErr := writeBuffer.PushContext("deviceValue"); pushErr != nil {
			return pushErr
		}
		_deviceValueErr := m.DeviceValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceValue"); popErr != nil {
			return popErr
		}
		if _deviceValueErr != nil {
			return errors.Wrap(_deviceValueErr, "Error serializing 'deviceValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetRecipientDevice"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetRecipientDevice) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
