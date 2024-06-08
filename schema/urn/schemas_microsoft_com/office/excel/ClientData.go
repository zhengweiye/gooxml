// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package excel

import (
	"encoding/xml"

	"github.com/zhengweiye/gooxml"
	"github.com/zhengweiye/gooxml/schema/soo/ofc/sharedTypes"
)

type ClientData struct {
	CT_ClientData
}

func NewClientData() *ClientData {
	ret := &ClientData{}
	ret.CT_ClientData = *NewCT_ClientData()
	return ret
}

func (m *ClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:excel"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "urn:schemas-microsoft-com:office:excel"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:ClientData"
	return m.CT_ClientData.MarshalXML(e, start)
}

func (m *ClientData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ClientData = *NewCT_ClientData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "ObjectType" {
			m.ObjectTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lClientData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MoveWithCells"}:
				m.MoveWithCells = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.MoveWithCells, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "SizeWithCells"}:
				m.SizeWithCells = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.SizeWithCells, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Anchor"}:
				m.Anchor = new(string)
				if err := d.DecodeElement(m.Anchor, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Locked"}:
				m.Locked = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Locked, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DefaultSize"}:
				m.DefaultSize = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.DefaultSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "PrintObject"}:
				m.PrintObject = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.PrintObject, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Disabled"}:
				m.Disabled = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Disabled, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoFill"}:
				m.AutoFill = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoLine"}:
				m.AutoLine = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoLine, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoPict"}:
				m.AutoPict = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoPict, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaMacro"}:
				m.FmlaMacro = new(string)
				if err := d.DecodeElement(m.FmlaMacro, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "TextHAlign"}:
				m.TextHAlign = new(string)
				if err := d.DecodeElement(m.TextHAlign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "TextVAlign"}:
				m.TextVAlign = new(string)
				if err := d.DecodeElement(m.TextVAlign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "LockText"}:
				m.LockText = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.LockText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "JustLastX"}:
				m.JustLastX = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.JustLastX, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "SecretEdit"}:
				m.SecretEdit = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.SecretEdit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Default"}:
				m.Default = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Default, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Help"}:
				m.Help = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Help, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Cancel"}:
				m.Cancel = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Cancel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Dismiss"}:
				m.Dismiss = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Dismiss, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Accel"}:
				m.Accel = new(int64)
				if err := d.DecodeElement(m.Accel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Accel2"}:
				m.Accel2 = new(int64)
				if err := d.DecodeElement(m.Accel2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Row"}:
				m.Row = new(int64)
				if err := d.DecodeElement(m.Row, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Column"}:
				m.Column = new(int64)
				if err := d.DecodeElement(m.Column, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Visible"}:
				m.Visible = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Visible, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "RowHidden"}:
				m.RowHidden = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.RowHidden, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ColHidden"}:
				m.ColHidden = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.ColHidden, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "VTEdit"}:
				m.VTEdit = new(int64)
				if err := d.DecodeElement(m.VTEdit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MultiLine"}:
				m.MultiLine = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.MultiLine, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "VScroll"}:
				m.VScroll = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.VScroll, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ValidIds"}:
				m.ValidIds = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.ValidIds, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaRange"}:
				m.FmlaRange = new(string)
				if err := d.DecodeElement(m.FmlaRange, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "WidthMin"}:
				m.WidthMin = new(int64)
				if err := d.DecodeElement(m.WidthMin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Sel"}:
				m.Sel = new(int64)
				if err := d.DecodeElement(m.Sel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "NoThreeD2"}:
				m.NoThreeD2 = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.NoThreeD2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "SelType"}:
				m.SelType = new(string)
				if err := d.DecodeElement(m.SelType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MultiSel"}:
				m.MultiSel = new(string)
				if err := d.DecodeElement(m.MultiSel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "LCT"}:
				m.LCT = new(string)
				if err := d.DecodeElement(m.LCT, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ListItem"}:
				m.ListItem = new(string)
				if err := d.DecodeElement(m.ListItem, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DropStyle"}:
				m.DropStyle = new(string)
				if err := d.DecodeElement(m.DropStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Colored"}:
				m.Colored = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Colored, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DropLines"}:
				m.DropLines = new(int64)
				if err := d.DecodeElement(m.DropLines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Checked"}:
				m.Checked = new(int64)
				if err := d.DecodeElement(m.Checked, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaLink"}:
				m.FmlaLink = new(string)
				if err := d.DecodeElement(m.FmlaLink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaPict"}:
				m.FmlaPict = new(string)
				if err := d.DecodeElement(m.FmlaPict, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "NoThreeD"}:
				m.NoThreeD = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.NoThreeD, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FirstButton"}:
				m.FirstButton = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.FirstButton, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaGroup"}:
				m.FmlaGroup = new(string)
				if err := d.DecodeElement(m.FmlaGroup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Val"}:
				m.Val = new(int64)
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Min"}:
				m.Min = new(int64)
				if err := d.DecodeElement(m.Min, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Max"}:
				m.Max = new(int64)
				if err := d.DecodeElement(m.Max, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Inc"}:
				m.Inc = new(int64)
				if err := d.DecodeElement(m.Inc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Page"}:
				m.Page = new(int64)
				if err := d.DecodeElement(m.Page, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Horiz"}:
				m.Horiz = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Horiz, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Dx"}:
				m.Dx = new(int64)
				if err := d.DecodeElement(m.Dx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MapOCX"}:
				m.MapOCX = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.MapOCX, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "CF"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.CF = append(m.CF, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Camera"}:
				m.Camera = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Camera, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "RecalcAlways"}:
				m.RecalcAlways = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.RecalcAlways, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoScale"}:
				m.AutoScale = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoScale, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DDE"}:
				m.DDE = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.DDE, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "UIObj"}:
				m.UIObj = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.UIObj, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptText"}:
				m.ScriptText = new(string)
				if err := d.DecodeElement(m.ScriptText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptExtended"}:
				m.ScriptExtended = new(string)
				if err := d.DecodeElement(m.ScriptExtended, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptLanguage"}:
				m.ScriptLanguage = new(uint32)
				if err := d.DecodeElement(m.ScriptLanguage, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptLocation"}:
				m.ScriptLocation = new(uint32)
				if err := d.DecodeElement(m.ScriptLocation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaTxbx"}:
				m.FmlaTxbx = new(string)
				if err := d.DecodeElement(m.FmlaTxbx, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on ClientData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lClientData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ClientData and its children
func (m *ClientData) Validate() error {
	return m.ValidateWithPath("ClientData")
}

// ValidateWithPath validates the ClientData and its children, prefixing error messages with path
func (m *ClientData) ValidateWithPath(path string) error {
	if err := m.CT_ClientData.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
