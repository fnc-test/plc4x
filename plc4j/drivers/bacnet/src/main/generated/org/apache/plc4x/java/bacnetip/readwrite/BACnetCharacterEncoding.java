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
package org.apache.plc4x.java.bacnetip.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetCharacterEncoding {
  ISO_10646((byte) 0x0),
  IBM_Microsoft_DBCS((byte) 0x1),
  JIS_X_0208((byte) 0x2),
  ISO_10646_4((byte) 0x3),
  ISO_10646_2((byte) 0x4),
  ISO_8859_1((byte) 0x5);
  private static final Map<Byte, BACnetCharacterEncoding> map;

  static {
    map = new HashMap<>();
    for (BACnetCharacterEncoding value : BACnetCharacterEncoding.values()) {
      map.put((byte) value.getValue(), value);
    }
  }

  private final byte value;

  BACnetCharacterEncoding(byte value) {
    this.value = value;
  }

  public byte getValue() {
    return value;
  }

  public static BACnetCharacterEncoding enumForValue(byte value) {
    return map.get(value);
  }

  public static Boolean isDefined(byte value) {
    return map.containsKey(value);
  }
}
