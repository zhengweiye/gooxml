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

	"github.com/zhengweiye/gooxml"
)

type CT_TupleCache struct {
	// Entries
	Entries *CT_PCDSDTCEntries
	// Sets
	Sets *CT_Sets
	// OLAP Query Cache
	QueryCache *CT_QueryCache
	// Server Formats
	ServerFormats *CT_ServerFormats
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_TupleCache() *CT_TupleCache {
	ret := &CT_TupleCache{}
	return ret
}

func (m *CT_TupleCache) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Entries != nil {
		seentries := xml.StartElement{Name: xml.Name{Local: "ma:entries"}}
		e.EncodeElement(m.Entries, seentries)
	}
	if m.Sets != nil {
		sesets := xml.StartElement{Name: xml.Name{Local: "ma:sets"}}
		e.EncodeElement(m.Sets, sesets)
	}
	if m.QueryCache != nil {
		sequeryCache := xml.StartElement{Name: xml.Name{Local: "ma:queryCache"}}
		e.EncodeElement(m.QueryCache, sequeryCache)
	}
	if m.ServerFormats != nil {
		seserverFormats := xml.StartElement{Name: xml.Name{Local: "ma:serverFormats"}}
		e.EncodeElement(m.ServerFormats, seserverFormats)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TupleCache) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TupleCache:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "entries"}:
				m.Entries = NewCT_PCDSDTCEntries()
				if err := d.DecodeElement(m.Entries, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sets"}:
				m.Sets = NewCT_Sets()
				if err := d.DecodeElement(m.Sets, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "queryCache"}:
				m.QueryCache = NewCT_QueryCache()
				if err := d.DecodeElement(m.QueryCache, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "serverFormats"}:
				m.ServerFormats = NewCT_ServerFormats()
				if err := d.DecodeElement(m.ServerFormats, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TupleCache %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TupleCache
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TupleCache and its children
func (m *CT_TupleCache) Validate() error {
	return m.ValidateWithPath("CT_TupleCache")
}

// ValidateWithPath validates the CT_TupleCache and its children, prefixing error messages with path
func (m *CT_TupleCache) ValidateWithPath(path string) error {
	if m.Entries != nil {
		if err := m.Entries.ValidateWithPath(path + "/Entries"); err != nil {
			return err
		}
	}
	if m.Sets != nil {
		if err := m.Sets.ValidateWithPath(path + "/Sets"); err != nil {
			return err
		}
	}
	if m.QueryCache != nil {
		if err := m.QueryCache.ValidateWithPath(path + "/QueryCache"); err != nil {
			return err
		}
	}
	if m.ServerFormats != nil {
		if err := m.ServerFormats.ValidateWithPath(path + "/ServerFormats"); err != nil {
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
