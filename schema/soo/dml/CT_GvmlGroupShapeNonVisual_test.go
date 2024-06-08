// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhengweiye/gooxml/schema/soo/dml"
)

func TestCT_GvmlGroupShapeNonVisualConstructor(t *testing.T) {
	v := dml.NewCT_GvmlGroupShapeNonVisual()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlGroupShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlGroupShapeNonVisual should validate: %s", err)
	}
}

func TestCT_GvmlGroupShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlGroupShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlGroupShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
