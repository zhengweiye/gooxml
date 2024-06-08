// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/zhengweiye/gooxml"
)

type EG_TextBulletColor struct {
	BuClrTx *CT_TextBulletColorFollowText
	BuClr   *CT_Color
}

func NewEG_TextBulletColor() *EG_TextBulletColor {
	ret := &EG_TextBulletColor{}
	return ret
}

func (m *EG_TextBulletColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BuClrTx != nil {
		sebuClrTx := xml.StartElement{Name: xml.Name{Local: "a:buClrTx"}}
		e.EncodeElement(m.BuClrTx, sebuClrTx)
	}
	if m.BuClr != nil {
		sebuClr := xml.StartElement{Name: xml.Name{Local: "a:buClr"}}
		e.EncodeElement(m.BuClr, sebuClr)
	}
	return nil
}

func (m *EG_TextBulletColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextBulletColor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buClrTx"}:
				m.BuClrTx = NewCT_TextBulletColorFollowText()
				if err := d.DecodeElement(m.BuClrTx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "buClr"}:
				m.BuClr = NewCT_Color()
				if err := d.DecodeElement(m.BuClr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_TextBulletColor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextBulletColor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TextBulletColor and its children
func (m *EG_TextBulletColor) Validate() error {
	return m.ValidateWithPath("EG_TextBulletColor")
}

// ValidateWithPath validates the EG_TextBulletColor and its children, prefixing error messages with path
func (m *EG_TextBulletColor) ValidateWithPath(path string) error {
	if m.BuClrTx != nil {
		if err := m.BuClrTx.ValidateWithPath(path + "/BuClrTx"); err != nil {
			return err
		}
	}
	if m.BuClr != nil {
		if err := m.BuClr.ValidateWithPath(path + "/BuClr"); err != nil {
			return err
		}
	}
	return nil
}