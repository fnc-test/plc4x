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
package org.apache.plc4x.protocol.iec608705104;

import org.apache.plc4x.plugins.codegenerator.language.mspec.parser.MessageFormatParser;
import org.apache.plc4x.plugins.codegenerator.language.mspec.protocol.ProtocolHelpers;
import org.apache.plc4x.plugins.codegenerator.language.mspec.protocol.ValidatableTypeContext;
import org.apache.plc4x.plugins.codegenerator.protocol.Protocol;
import org.apache.plc4x.plugins.codegenerator.protocol.TypeContext;
import org.apache.plc4x.plugins.codegenerator.types.exceptions.GenerationException;

public class IEC608705104Protocol implements Protocol, ProtocolHelpers {

    @Override
    public String getName() {
        return "iec-60870-5-104";
    }

    @Override
    public TypeContext getTypeContext() throws GenerationException {
        ValidatableTypeContext typeContext = new MessageFormatParser().parse(getMspecStream("iec-60870-5-104"));
        typeContext.validate();
        return typeContext;
    }

}
