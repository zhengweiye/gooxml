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
	"strconv"

	"github.com/zhengweiye/gooxml"
)

type CT_FutureMetadata struct {
	// Metadata Type Name
	NameAttr string
	// Future Metadata Block Count
	CountAttr *uint32
	// Future Metadata Block
	Bk []*CT_FutureMetadataBlock
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_FutureMetadata() *CT_FutureMetadata {
	ret := &CT_FutureMetadata{}
	return ret
}

func (m *CT_FutureMetadata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.Bk != nil {
		sebk := xml.StartElement{Name: xml.Name{Local: "ma:bk"}}
		for _, c := range m.Bk {
			e.EncodeElement(c, sebk)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FutureMetadata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_FutureMetadata:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "bk"}:
				tmp := NewCT_FutureMetadataBlock()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Bk = append(m.Bk, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_FutureMetadata %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FutureMetadata
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FutureMetadata and its children
func (m *CT_FutureMetadata) Validate() error {
	return m.ValidateWithPath("CT_FutureMetadata")
}

// ValidateWithPath validates the CT_FutureMetadata and its children, prefixing error messages with path
func (m *CT_FutureMetadata) ValidateWithPath(path string) error {
	for i, v := range m.Bk {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Bk[%d]", path, i)); err != nil {
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
