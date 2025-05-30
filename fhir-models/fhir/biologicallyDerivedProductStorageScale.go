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
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/lifebox-healthcare/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// BiologicallyDerivedProductStorageScale is documented here http://hl7.org/fhir/ValueSet/product-storage-scale
type BiologicallyDerivedProductStorageScale int

const (
	BiologicallyDerivedProductStorageScaleFarenheit BiologicallyDerivedProductStorageScale = iota
	BiologicallyDerivedProductStorageScaleCelsius
	BiologicallyDerivedProductStorageScaleKelvin
)

func (code BiologicallyDerivedProductStorageScale) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BiologicallyDerivedProductStorageScale) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "farenheit":
		*code = BiologicallyDerivedProductStorageScaleFarenheit
	case "celsius":
		*code = BiologicallyDerivedProductStorageScaleCelsius
	case "kelvin":
		*code = BiologicallyDerivedProductStorageScaleKelvin
	default:
		return fmt.Errorf("unknown BiologicallyDerivedProductStorageScale code `%s`", s)
	}
	return nil
}
func (code BiologicallyDerivedProductStorageScale) String() string {
	return code.Code()
}
func (code BiologicallyDerivedProductStorageScale) Code() string {
	switch code {
	case BiologicallyDerivedProductStorageScaleFarenheit:
		return "farenheit"
	case BiologicallyDerivedProductStorageScaleCelsius:
		return "celsius"
	case BiologicallyDerivedProductStorageScaleKelvin:
		return "kelvin"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStorageScale) Display() string {
	switch code {
	case BiologicallyDerivedProductStorageScaleFarenheit:
		return "Fahrenheit"
	case BiologicallyDerivedProductStorageScaleCelsius:
		return "Celsius"
	case BiologicallyDerivedProductStorageScaleKelvin:
		return "Kelvin"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStorageScale) Definition() string {
	switch code {
	case BiologicallyDerivedProductStorageScaleFarenheit:
		return "Fahrenheit temperature scale."
	case BiologicallyDerivedProductStorageScaleCelsius:
		return "Celsius or centigrade temperature scale."
	case BiologicallyDerivedProductStorageScaleKelvin:
		return "Kelvin absolute thermodynamic temperature scale."
	}
	return "<unknown>"
}
