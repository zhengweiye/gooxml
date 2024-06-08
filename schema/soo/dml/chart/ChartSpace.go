// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"

	"github.com/zhengweiye/gooxml"
	"github.com/zhengweiye/gooxml/schema/soo/dml"
)

type ChartSpace struct {
	CT_ChartSpace
}

func NewChartSpace() *ChartSpace {
	ret := &ChartSpace{}
	ret.CT_ChartSpace = *NewCT_ChartSpace()
	return ret
}

func (m *ChartSpace) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:c"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "c:chartSpace"
	return m.CT_ChartSpace.MarshalXML(e, start)
}

func (m *ChartSpace) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ChartSpace = *NewCT_ChartSpace()
lChartSpace:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "date1904"}:
				m.Date1904 = NewCT_Boolean()
				if err := d.DecodeElement(m.Date1904, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "lang"}:
				m.Lang = NewCT_TextLanguageID()
				if err := d.DecodeElement(m.Lang, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "roundedCorners"}:
				m.RoundedCorners = NewCT_Boolean()
				if err := d.DecodeElement(m.RoundedCorners, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "style"}:
				m.Style = NewCT_Style()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "clrMapOvr"}:
				m.ClrMapOvr = dml.NewCT_ColorMapping()
				if err := d.DecodeElement(m.ClrMapOvr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pivotSource"}:
				m.PivotSource = NewCT_PivotSource()
				if err := d.DecodeElement(m.PivotSource, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "protection"}:
				m.Protection = NewCT_Protection()
				if err := d.DecodeElement(m.Protection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "chart"}:
				if err := d.DecodeElement(m.Chart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"}:
				m.TxPr = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "externalData"}:
				m.ExternalData = NewCT_ExternalData()
				if err := d.DecodeElement(m.ExternalData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "printSettings"}:
				m.PrintSettings = NewCT_PrintSettings()
				if err := d.DecodeElement(m.PrintSettings, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "userShapes"}:
				m.UserShapes = NewCT_RelId()
				if err := d.DecodeElement(m.UserShapes, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on ChartSpace %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lChartSpace
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ChartSpace and its children
func (m *ChartSpace) Validate() error {
	return m.ValidateWithPath("ChartSpace")
}

// ValidateWithPath validates the ChartSpace and its children, prefixing error messages with path
func (m *ChartSpace) ValidateWithPath(path string) error {
	if err := m.CT_ChartSpace.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
