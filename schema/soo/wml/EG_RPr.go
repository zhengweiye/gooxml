// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/zhengweiye/gooxml"
)

type EG_RPr struct {
	// Run Properties
	RPr *CT_RPr
}

func NewEG_RPr() *EG_RPr {
	ret := &EG_RPr{}
	return ret
}

func (m *EG_RPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	return nil
}

func (m *EG_RPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_RPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"}:
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_RPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_RPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_RPr and its children
func (m *EG_RPr) Validate() error {
	return m.ValidateWithPath("EG_RPr")
}

// ValidateWithPath validates the EG_RPr and its children, prefixing error messages with path
func (m *EG_RPr) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	return nil
}
