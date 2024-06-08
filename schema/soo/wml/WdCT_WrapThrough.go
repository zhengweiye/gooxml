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
	"fmt"
	"strconv"

	"github.com/zhengweiye/gooxml"
)

type WdCT_WrapThrough struct {
	WrapTextAttr WdST_WrapText
	DistLAttr    *uint32
	DistRAttr    *uint32
	WrapPolygon  *WdCT_WrapPath
}

func NewWdCT_WrapThrough() *WdCT_WrapThrough {
	ret := &WdCT_WrapThrough{}
	ret.WrapTextAttr = WdST_WrapText(1)
	ret.WrapPolygon = NewWdCT_WrapPath()
	return ret
}

func (m *WdCT_WrapThrough) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.WrapTextAttr.MarshalXMLAttr(xml.Name{Local: "wrapText"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"},
			Value: fmt.Sprintf("%v", *m.DistLAttr)})
	}
	if m.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"},
			Value: fmt.Sprintf("%v", *m.DistRAttr)})
	}
	e.EncodeToken(start)
	sewrapPolygon := xml.StartElement{Name: xml.Name{Local: "wp:wrapPolygon"}}
	e.EncodeElement(m.WrapPolygon, sewrapPolygon)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WrapThrough) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.WrapTextAttr = WdST_WrapText(1)
	m.WrapPolygon = NewWdCT_WrapPath()
	for _, attr := range start.Attr {
		if attr.Name.Local == "wrapText" {
			m.WrapTextAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
			continue
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
			continue
		}
	}
lWdCT_WrapThrough:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapPolygon"}:
				if err := d.DecodeElement(m.WrapPolygon, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_WrapThrough %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WrapThrough
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WrapThrough and its children
func (m *WdCT_WrapThrough) Validate() error {
	return m.ValidateWithPath("WdCT_WrapThrough")
}

// ValidateWithPath validates the WdCT_WrapThrough and its children, prefixing error messages with path
func (m *WdCT_WrapThrough) ValidateWithPath(path string) error {
	if m.WrapTextAttr == WdST_WrapTextUnset {
		return fmt.Errorf("%s/WrapTextAttr is a mandatory field", path)
	}
	if err := m.WrapTextAttr.ValidateWithPath(path + "/WrapTextAttr"); err != nil {
		return err
	}
	if err := m.WrapPolygon.ValidateWithPath(path + "/WrapPolygon"); err != nil {
		return err
	}
	return nil
}
