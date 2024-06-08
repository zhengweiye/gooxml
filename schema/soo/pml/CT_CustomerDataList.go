// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/zhengweiye/gooxml"
)

type CT_CustomerDataList struct {
	// Customer Data
	CustData []*CT_CustomerData
	// Customer Data Tags
	Tags *CT_TagsData
}

func NewCT_CustomerDataList() *CT_CustomerDataList {
	ret := &CT_CustomerDataList{}
	return ret
}

func (m *CT_CustomerDataList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CustData != nil {
		secustData := xml.StartElement{Name: xml.Name{Local: "p:custData"}}
		for _, c := range m.CustData {
			e.EncodeElement(c, secustData)
		}
	}
	if m.Tags != nil {
		setags := xml.StartElement{Name: xml.Name{Local: "p:tags"}}
		e.EncodeElement(m.Tags, setags)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomerDataList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomerDataList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "custData"}:
				tmp := NewCT_CustomerData()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustData = append(m.CustData, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tags"}:
				m.Tags = NewCT_TagsData()
				if err := d.DecodeElement(m.Tags, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_CustomerDataList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomerDataList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomerDataList and its children
func (m *CT_CustomerDataList) Validate() error {
	return m.ValidateWithPath("CT_CustomerDataList")
}

// ValidateWithPath validates the CT_CustomerDataList and its children, prefixing error messages with path
func (m *CT_CustomerDataList) ValidateWithPath(path string) error {
	for i, v := range m.CustData {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustData[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Tags != nil {
		if err := m.Tags.ValidateWithPath(path + "/Tags"); err != nil {
			return err
		}
	}
	return nil
}
