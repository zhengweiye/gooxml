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

type EG_Text3D struct {
	Sp3d   *CT_Shape3D
	FlatTx *CT_FlatText
}

func NewEG_Text3D() *EG_Text3D {
	ret := &EG_Text3D{}
	return ret
}

func (m *EG_Text3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Sp3d != nil {
		sesp3d := xml.StartElement{Name: xml.Name{Local: "a:sp3d"}}
		e.EncodeElement(m.Sp3d, sesp3d)
	}
	if m.FlatTx != nil {
		seflatTx := xml.StartElement{Name: xml.Name{Local: "a:flatTx"}}
		e.EncodeElement(m.FlatTx, seflatTx)
	}
	return nil
}

func (m *EG_Text3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Text3D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sp3d"}:
				m.Sp3d = NewCT_Shape3D()
				if err := d.DecodeElement(m.Sp3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "flatTx"}:
				m.FlatTx = NewCT_FlatText()
				if err := d.DecodeElement(m.FlatTx, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_Text3D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Text3D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_Text3D and its children
func (m *EG_Text3D) Validate() error {
	return m.ValidateWithPath("EG_Text3D")
}

// ValidateWithPath validates the EG_Text3D and its children, prefixing error messages with path
func (m *EG_Text3D) ValidateWithPath(path string) error {
	if m.Sp3d != nil {
		if err := m.Sp3d.ValidateWithPath(path + "/Sp3d"); err != nil {
			return err
		}
	}
	if m.FlatTx != nil {
		if err := m.FlatTx.ValidateWithPath(path + "/FlatTx"); err != nil {
			return err
		}
	}
	return nil
}