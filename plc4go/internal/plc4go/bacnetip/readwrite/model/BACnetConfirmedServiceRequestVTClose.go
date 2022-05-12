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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestVTClose is the data-structure of this message
type BACnetConfirmedServiceRequestVTClose struct {
	*BACnetConfirmedServiceRequest

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestVTClose is the corresponding interface of BACnetConfirmedServiceRequestVTClose
type IBACnetConfirmedServiceRequestVTClose interface {
	IBACnetConfirmedServiceRequest
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

func (m *BACnetConfirmedServiceRequestVTClose) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestVTClose) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestVTClose) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

// NewBACnetConfirmedServiceRequestVTClose factory function for BACnetConfirmedServiceRequestVTClose
func NewBACnetConfirmedServiceRequestVTClose(serviceRequestLength uint16) *BACnetConfirmedServiceRequestVTClose {
	_result := &BACnetConfirmedServiceRequestVTClose{
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestVTClose(structType interface{}) *BACnetConfirmedServiceRequestVTClose {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTClose); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTClose); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestVTClose(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestVTClose(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestVTClose) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTClose"
}

func (m *BACnetConfirmedServiceRequestVTClose) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestVTClose) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestVTClose) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestVTCloseParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestVTClose, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTClose"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseAssertError{"TODO: implement me"}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTClose"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestVTClose{
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestVTClose) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTClose"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTClose"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestVTClose) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
