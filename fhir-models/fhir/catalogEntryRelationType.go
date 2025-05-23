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

// CatalogEntryRelationType is documented here http://hl7.org/fhir/ValueSet/relation-type
type CatalogEntryRelationType int

const (
	CatalogEntryRelationTypeTriggers CatalogEntryRelationType = iota
	CatalogEntryRelationTypeIsReplacedBy
)

func (code CatalogEntryRelationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CatalogEntryRelationType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "triggers":
		*code = CatalogEntryRelationTypeTriggers
	case "is-replaced-by":
		*code = CatalogEntryRelationTypeIsReplacedBy
	default:
		return fmt.Errorf("unknown CatalogEntryRelationType code `%s`", s)
	}
	return nil
}
func (code CatalogEntryRelationType) String() string {
	return code.Code()
}
func (code CatalogEntryRelationType) Code() string {
	switch code {
	case CatalogEntryRelationTypeTriggers:
		return "triggers"
	case CatalogEntryRelationTypeIsReplacedBy:
		return "is-replaced-by"
	}
	return "<unknown>"
}
func (code CatalogEntryRelationType) Display() string {
	switch code {
	case CatalogEntryRelationTypeTriggers:
		return "Triggers"
	case CatalogEntryRelationTypeIsReplacedBy:
		return "Replaced By"
	}
	return "<unknown>"
}
func (code CatalogEntryRelationType) Definition() string {
	switch code {
	case CatalogEntryRelationTypeTriggers:
		return "the related entry represents an activity that may be triggered by the current item."
	case CatalogEntryRelationTypeIsReplacedBy:
		return "the related entry represents an item that replaces the current retired item."
	}
	return "<unknown>"
}
