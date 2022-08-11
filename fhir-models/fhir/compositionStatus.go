// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/michaelsauter/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// CompositionStatus is documented here http://hl7.org/fhir/ValueSet/composition-status
type CompositionStatus int

const (
	CompositionStatusPreliminary CompositionStatus = iota
	CompositionStatusFinal
	CompositionStatusAmended
	CompositionStatusEnteredInError
)

func (code CompositionStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(false)
	err := enc.Encode(code.Code())
	return buffer.Bytes(), err
}
func (code *CompositionStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "preliminary":
		*code = CompositionStatusPreliminary
	case "final":
		*code = CompositionStatusFinal
	case "amended":
		*code = CompositionStatusAmended
	case "entered-in-error":
		*code = CompositionStatusEnteredInError
	default:
		return fmt.Errorf("unknown CompositionStatus code `%s`", s)
	}
	return nil
}
func (code CompositionStatus) String() string {
	return code.Code()
}
func (code CompositionStatus) Code() string {
	switch code {
	case CompositionStatusPreliminary:
		return "preliminary"
	case CompositionStatusFinal:
		return "final"
	case CompositionStatusAmended:
		return "amended"
	case CompositionStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code CompositionStatus) Display() string {
	switch code {
	case CompositionStatusPreliminary:
		return "Preliminary"
	case CompositionStatusFinal:
		return "Final"
	case CompositionStatusAmended:
		return "Amended"
	case CompositionStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code CompositionStatus) Definition() string {
	switch code {
	case CompositionStatusPreliminary:
		return "This is a preliminary composition or document (also known as initial or interim). The content may be incomplete or unverified."
	case CompositionStatusFinal:
		return "This version of the composition is complete and verified by an appropriate person and no further work is planned. Any subsequent updates would be on a new version of the composition."
	case CompositionStatusAmended:
		return "The composition content or the referenced resources have been modified (edited or added to) subsequent to being released as \"final\" and the composition is complete and verified by an authorized person."
	case CompositionStatusEnteredInError:
		return "The composition or document was originally created/issued in error, and this is an amendment that marks that the entire series should not be considered as valid."
	}
	return "<unknown>"
}
