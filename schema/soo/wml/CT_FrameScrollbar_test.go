// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhengweiye/gooxml/schema/soo/wml"
)

func TestCT_FrameScrollbarConstructor(t *testing.T) {
	v := wml.NewCT_FrameScrollbar()
	if v == nil {
		t.Errorf("wml.NewCT_FrameScrollbar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FrameScrollbar should validate: %s", err)
	}
}

func TestCT_FrameScrollbarMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FrameScrollbar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FrameScrollbar()
	xml.Unmarshal(buf, v2)
}