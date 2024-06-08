// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/zhengweiye/gooxml"
)

type CT_ChartsheetViews struct {
	// Chart Sheet View
	SheetView []*CT_ChartsheetView
	ExtLst    *CT_ExtensionList
}

func NewCT_ChartsheetViews() *CT_ChartsheetViews {
	ret := &CT_ChartsheetViews{}
	return ret
}

func (m *CT_ChartsheetViews) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sesheetView := xml.StartElement{Name: xml.Name{Local: "ma:sheetView"}}
	for _, c := range m.SheetView {
		e.EncodeElement(c, sesheetView)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ChartsheetViews) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ChartsheetViews:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetView"}:
				tmp := NewCT_ChartsheetView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SheetView = append(m.SheetView, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ChartsheetViews %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ChartsheetViews
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ChartsheetViews and its children
func (m *CT_ChartsheetViews) Validate() error {
	return m.ValidateWithPath("CT_ChartsheetViews")
}

// ValidateWithPath validates the CT_ChartsheetViews and its children, prefixing error messages with path
func (m *CT_ChartsheetViews) ValidateWithPath(path string) error {
	for i, v := range m.SheetView {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SheetView[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
