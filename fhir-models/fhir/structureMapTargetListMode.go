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

// StructureMapTargetListMode is documented here http://hl7.org/fhir/ValueSet/map-target-list-mode
type StructureMapTargetListMode int

const (
	StructureMapTargetListModeFirst StructureMapTargetListMode = iota
	StructureMapTargetListModeShare
	StructureMapTargetListModeLast
	StructureMapTargetListModeCollate
)

func (code StructureMapTargetListMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapTargetListMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "first":
		*code = StructureMapTargetListModeFirst
	case "share":
		*code = StructureMapTargetListModeShare
	case "last":
		*code = StructureMapTargetListModeLast
	case "collate":
		*code = StructureMapTargetListModeCollate
	default:
		return fmt.Errorf("unknown StructureMapTargetListMode code `%s`", s)
	}
	return nil
}
func (code StructureMapTargetListMode) String() string {
	return code.Code()
}
func (code StructureMapTargetListMode) Code() string {
	switch code {
	case StructureMapTargetListModeFirst:
		return "first"
	case StructureMapTargetListModeShare:
		return "share"
	case StructureMapTargetListModeLast:
		return "last"
	case StructureMapTargetListModeCollate:
		return "collate"
	}
	return "<unknown>"
}
func (code StructureMapTargetListMode) Display() string {
	switch code {
	case StructureMapTargetListModeFirst:
		return "First"
	case StructureMapTargetListModeShare:
		return "Share"
	case StructureMapTargetListModeLast:
		return "Last"
	case StructureMapTargetListModeCollate:
		return "Collate"
	}
	return "<unknown>"
}
func (code StructureMapTargetListMode) Definition() string {
	switch code {
	case StructureMapTargetListModeFirst:
		return "when the target list is being assembled, the items for this rule go first. If more than one rule defines a first item (for a given instance of mapping) then this is an error."
	case StructureMapTargetListModeShare:
		return "the target instance is shared with the target instances generated by another rule (up to the first common n items, then create new ones)."
	case StructureMapTargetListModeLast:
		return "when the target list is being assembled, the items for this rule go last. If more than one rule defines a last item (for a given instance of mapping) then this is an error."
	case StructureMapTargetListModeCollate:
		return "re-use the first item in the list, and keep adding content to it."
	}
	return "<unknown>"
}
