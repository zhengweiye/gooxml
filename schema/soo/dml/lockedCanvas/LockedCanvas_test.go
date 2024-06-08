// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package lockedCanvas_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhengweiye/gooxml/schema/soo/dml/lockedCanvas"
)

func TestLockedCanvasConstructor(t *testing.T) {
	v := lockedCanvas.NewLockedCanvas()
	if v == nil {
		t.Errorf("lockedCanvas.NewLockedCanvas must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed lockedCanvas.LockedCanvas should validate: %s", err)
	}
}

func TestLockedCanvasMarshalUnmarshal(t *testing.T) {
	v := lockedCanvas.NewLockedCanvas()
	buf, _ := xml.Marshal(v)
	v2 := lockedCanvas.NewLockedCanvas()
	xml.Unmarshal(buf, v2)
}
