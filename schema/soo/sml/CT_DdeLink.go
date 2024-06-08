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

type CT_DdeLink struct {
	// Service name
	DdeServiceAttr string
	// Topic for DDE server
	DdeTopicAttr string
	// DDE Items Collection
	DdeItems *CT_DdeItems
}

func NewCT_DdeLink() *CT_DdeLink {
	ret := &CT_DdeLink{}
	return ret
}

func (m *CT_DdeLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ddeService"},
		Value: fmt.Sprintf("%v", m.DdeServiceAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ddeTopic"},
		Value: fmt.Sprintf("%v", m.DdeTopicAttr)})
	e.EncodeToken(start)
	if m.DdeItems != nil {
		seddeItems := xml.StartElement{Name: xml.Name{Local: "ma:ddeItems"}}
		e.EncodeElement(m.DdeItems, seddeItems)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DdeLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ddeService" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DdeServiceAttr = parsed
			continue
		}
		if attr.Name.Local == "ddeTopic" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DdeTopicAttr = parsed
			continue
		}
	}
lCT_DdeLink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ddeItems"}:
				m.DdeItems = NewCT_DdeItems()
				if err := d.DecodeElement(m.DdeItems, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_DdeLink %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DdeLink
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DdeLink and its children
func (m *CT_DdeLink) Validate() error {
	return m.ValidateWithPath("CT_DdeLink")
}

// ValidateWithPath validates the CT_DdeLink and its children, prefixing error messages with path
func (m *CT_DdeLink) ValidateWithPath(path string) error {
	if m.DdeItems != nil {
		if err := m.DdeItems.ValidateWithPath(path + "/DdeItems"); err != nil {
			return err
		}
	}
	return nil
}
