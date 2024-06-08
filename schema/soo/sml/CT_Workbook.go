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
	"github.com/zhengweiye/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_Workbook struct {
	// Document Conformance Class
	ConformanceAttr sharedTypes.ST_ConformanceClass
	// File Version
	FileVersion *CT_FileVersion
	// File Sharing
	FileSharing *CT_FileSharing
	// Workbook Properties
	WorkbookPr *CT_WorkbookPr
	// Workbook Protection
	WorkbookProtection *CT_WorkbookProtection
	// Workbook Views
	BookViews *CT_BookViews
	// Sheets
	Sheets *CT_Sheets
	// Function Groups
	FunctionGroups *CT_FunctionGroups
	// External References
	ExternalReferences *CT_ExternalReferences
	// Defined Names
	DefinedNames *CT_DefinedNames
	// Calculation Properties
	CalcPr *CT_CalcPr
	// Embedded Object Size
	OleSize *CT_OleSize
	// Custom Workbook Views
	CustomWorkbookViews *CT_CustomWorkbookViews
	// PivotCaches
	PivotCaches *CT_PivotCaches
	// Smart Tag Properties
	SmartTagPr *CT_SmartTagPr
	// Smart Tag Types
	SmartTagTypes *CT_SmartTagTypes
	// Web Publishing Properties
	WebPublishing *CT_WebPublishing
	// File Recovery Properties
	FileRecoveryPr []*CT_FileRecoveryPr
	// Web Publish Objects
	WebPublishObjects *CT_WebPublishObjects
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Workbook() *CT_Workbook {
	ret := &CT_Workbook{}
	ret.Sheets = NewCT_Sheets()
	return ret
}

func (m *CT_Workbook) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ConformanceAttr != sharedTypes.ST_ConformanceClassUnset {
		attr, err := m.ConformanceAttr.MarshalXMLAttr(xml.Name{Local: "conformance"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.FileVersion != nil {
		sefileVersion := xml.StartElement{Name: xml.Name{Local: "ma:fileVersion"}}
		e.EncodeElement(m.FileVersion, sefileVersion)
	}
	if m.FileSharing != nil {
		sefileSharing := xml.StartElement{Name: xml.Name{Local: "ma:fileSharing"}}
		e.EncodeElement(m.FileSharing, sefileSharing)
	}
	if m.WorkbookPr != nil {
		seworkbookPr := xml.StartElement{Name: xml.Name{Local: "ma:workbookPr"}}
		e.EncodeElement(m.WorkbookPr, seworkbookPr)
	}
	if m.WorkbookProtection != nil {
		seworkbookProtection := xml.StartElement{Name: xml.Name{Local: "ma:workbookProtection"}}
		e.EncodeElement(m.WorkbookProtection, seworkbookProtection)
	}
	if m.BookViews != nil {
		sebookViews := xml.StartElement{Name: xml.Name{Local: "ma:bookViews"}}
		e.EncodeElement(m.BookViews, sebookViews)
	}
	sesheets := xml.StartElement{Name: xml.Name{Local: "ma:sheets"}}
	e.EncodeElement(m.Sheets, sesheets)
	if m.FunctionGroups != nil {
		sefunctionGroups := xml.StartElement{Name: xml.Name{Local: "ma:functionGroups"}}
		e.EncodeElement(m.FunctionGroups, sefunctionGroups)
	}
	if m.ExternalReferences != nil {
		seexternalReferences := xml.StartElement{Name: xml.Name{Local: "ma:externalReferences"}}
		e.EncodeElement(m.ExternalReferences, seexternalReferences)
	}
	if m.DefinedNames != nil {
		sedefinedNames := xml.StartElement{Name: xml.Name{Local: "ma:definedNames"}}
		e.EncodeElement(m.DefinedNames, sedefinedNames)
	}
	if m.CalcPr != nil {
		secalcPr := xml.StartElement{Name: xml.Name{Local: "ma:calcPr"}}
		e.EncodeElement(m.CalcPr, secalcPr)
	}
	if m.OleSize != nil {
		seoleSize := xml.StartElement{Name: xml.Name{Local: "ma:oleSize"}}
		e.EncodeElement(m.OleSize, seoleSize)
	}
	if m.CustomWorkbookViews != nil {
		secustomWorkbookViews := xml.StartElement{Name: xml.Name{Local: "ma:customWorkbookViews"}}
		e.EncodeElement(m.CustomWorkbookViews, secustomWorkbookViews)
	}
	if m.PivotCaches != nil {
		sepivotCaches := xml.StartElement{Name: xml.Name{Local: "ma:pivotCaches"}}
		e.EncodeElement(m.PivotCaches, sepivotCaches)
	}
	if m.SmartTagPr != nil {
		sesmartTagPr := xml.StartElement{Name: xml.Name{Local: "ma:smartTagPr"}}
		e.EncodeElement(m.SmartTagPr, sesmartTagPr)
	}
	if m.SmartTagTypes != nil {
		sesmartTagTypes := xml.StartElement{Name: xml.Name{Local: "ma:smartTagTypes"}}
		e.EncodeElement(m.SmartTagTypes, sesmartTagTypes)
	}
	if m.WebPublishing != nil {
		sewebPublishing := xml.StartElement{Name: xml.Name{Local: "ma:webPublishing"}}
		e.EncodeElement(m.WebPublishing, sewebPublishing)
	}
	if m.FileRecoveryPr != nil {
		sefileRecoveryPr := xml.StartElement{Name: xml.Name{Local: "ma:fileRecoveryPr"}}
		for _, c := range m.FileRecoveryPr {
			e.EncodeElement(c, sefileRecoveryPr)
		}
	}
	if m.WebPublishObjects != nil {
		sewebPublishObjects := xml.StartElement{Name: xml.Name{Local: "ma:webPublishObjects"}}
		e.EncodeElement(m.WebPublishObjects, sewebPublishObjects)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Workbook) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Sheets = NewCT_Sheets()
	for _, attr := range start.Attr {
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Workbook:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fileVersion"}:
				m.FileVersion = NewCT_FileVersion()
				if err := d.DecodeElement(m.FileVersion, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fileSharing"}:
				m.FileSharing = NewCT_FileSharing()
				if err := d.DecodeElement(m.FileSharing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "workbookPr"}:
				m.WorkbookPr = NewCT_WorkbookPr()
				if err := d.DecodeElement(m.WorkbookPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "workbookProtection"}:
				m.WorkbookProtection = NewCT_WorkbookProtection()
				if err := d.DecodeElement(m.WorkbookProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "bookViews"}:
				m.BookViews = NewCT_BookViews()
				if err := d.DecodeElement(m.BookViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheets"}:
				if err := d.DecodeElement(m.Sheets, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "functionGroups"}:
				m.FunctionGroups = NewCT_FunctionGroups()
				if err := d.DecodeElement(m.FunctionGroups, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "externalReferences"}:
				m.ExternalReferences = NewCT_ExternalReferences()
				if err := d.DecodeElement(m.ExternalReferences, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "definedNames"}:
				m.DefinedNames = NewCT_DefinedNames()
				if err := d.DecodeElement(m.DefinedNames, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "calcPr"}:
				m.CalcPr = NewCT_CalcPr()
				if err := d.DecodeElement(m.CalcPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oleSize"}:
				m.OleSize = NewCT_OleSize()
				if err := d.DecodeElement(m.OleSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customWorkbookViews"}:
				m.CustomWorkbookViews = NewCT_CustomWorkbookViews()
				if err := d.DecodeElement(m.CustomWorkbookViews, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotCaches"}:
				m.PivotCaches = NewCT_PivotCaches()
				if err := d.DecodeElement(m.PivotCaches, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "smartTagPr"}:
				m.SmartTagPr = NewCT_SmartTagPr()
				if err := d.DecodeElement(m.SmartTagPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "smartTagTypes"}:
				m.SmartTagTypes = NewCT_SmartTagTypes()
				if err := d.DecodeElement(m.SmartTagTypes, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishing"}:
				m.WebPublishing = NewCT_WebPublishing()
				if err := d.DecodeElement(m.WebPublishing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fileRecoveryPr"}:
				tmp := NewCT_FileRecoveryPr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FileRecoveryPr = append(m.FileRecoveryPr, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "webPublishObjects"}:
				m.WebPublishObjects = NewCT_WebPublishObjects()
				if err := d.DecodeElement(m.WebPublishObjects, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Workbook %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Workbook
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Workbook and its children
func (m *CT_Workbook) Validate() error {
	return m.ValidateWithPath("CT_Workbook")
}

// ValidateWithPath validates the CT_Workbook and its children, prefixing error messages with path
func (m *CT_Workbook) ValidateWithPath(path string) error {
	if err := m.ConformanceAttr.ValidateWithPath(path + "/ConformanceAttr"); err != nil {
		return err
	}
	if m.FileVersion != nil {
		if err := m.FileVersion.ValidateWithPath(path + "/FileVersion"); err != nil {
			return err
		}
	}
	if m.FileSharing != nil {
		if err := m.FileSharing.ValidateWithPath(path + "/FileSharing"); err != nil {
			return err
		}
	}
	if m.WorkbookPr != nil {
		if err := m.WorkbookPr.ValidateWithPath(path + "/WorkbookPr"); err != nil {
			return err
		}
	}
	if m.WorkbookProtection != nil {
		if err := m.WorkbookProtection.ValidateWithPath(path + "/WorkbookProtection"); err != nil {
			return err
		}
	}
	if m.BookViews != nil {
		if err := m.BookViews.ValidateWithPath(path + "/BookViews"); err != nil {
			return err
		}
	}
	if err := m.Sheets.ValidateWithPath(path + "/Sheets"); err != nil {
		return err
	}
	if m.FunctionGroups != nil {
		if err := m.FunctionGroups.ValidateWithPath(path + "/FunctionGroups"); err != nil {
			return err
		}
	}
	if m.ExternalReferences != nil {
		if err := m.ExternalReferences.ValidateWithPath(path + "/ExternalReferences"); err != nil {
			return err
		}
	}
	if m.DefinedNames != nil {
		if err := m.DefinedNames.ValidateWithPath(path + "/DefinedNames"); err != nil {
			return err
		}
	}
	if m.CalcPr != nil {
		if err := m.CalcPr.ValidateWithPath(path + "/CalcPr"); err != nil {
			return err
		}
	}
	if m.OleSize != nil {
		if err := m.OleSize.ValidateWithPath(path + "/OleSize"); err != nil {
			return err
		}
	}
	if m.CustomWorkbookViews != nil {
		if err := m.CustomWorkbookViews.ValidateWithPath(path + "/CustomWorkbookViews"); err != nil {
			return err
		}
	}
	if m.PivotCaches != nil {
		if err := m.PivotCaches.ValidateWithPath(path + "/PivotCaches"); err != nil {
			return err
		}
	}
	if m.SmartTagPr != nil {
		if err := m.SmartTagPr.ValidateWithPath(path + "/SmartTagPr"); err != nil {
			return err
		}
	}
	if m.SmartTagTypes != nil {
		if err := m.SmartTagTypes.ValidateWithPath(path + "/SmartTagTypes"); err != nil {
			return err
		}
	}
	if m.WebPublishing != nil {
		if err := m.WebPublishing.ValidateWithPath(path + "/WebPublishing"); err != nil {
			return err
		}
	}
	for i, v := range m.FileRecoveryPr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FileRecoveryPr[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.WebPublishObjects != nil {
		if err := m.WebPublishObjects.ValidateWithPath(path + "/WebPublishObjects"); err != nil {
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
