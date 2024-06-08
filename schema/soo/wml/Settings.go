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

	"github.com/zhengweiye/gooxml"
	"github.com/zhengweiye/gooxml/schema/soo/ofc/math"
	"github.com/zhengweiye/gooxml/schema/soo/schemaLibrary"
)

type Settings struct {
	CT_Settings
}

func NewSettings() *Settings {
	ret := &Settings{}
	ret.CT_Settings = *NewCT_Settings()
	return ret
}

func (m *Settings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:settings"
	return m.CT_Settings.MarshalXML(e, start)
}

func (m *Settings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Settings = *NewCT_Settings()
lSettings:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "writeProtection"}:
				m.WriteProtection = NewCT_WriteProtection()
				if err := d.DecodeElement(m.WriteProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "view"}:
				m.View = NewCT_View()
				if err := d.DecodeElement(m.View, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "zoom"}:
				m.Zoom = NewCT_Zoom()
				if err := d.DecodeElement(m.Zoom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "removePersonalInformation"}:
				m.RemovePersonalInformation = NewCT_OnOff()
				if err := d.DecodeElement(m.RemovePersonalInformation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "removeDateAndTime"}:
				m.RemoveDateAndTime = NewCT_OnOff()
				if err := d.DecodeElement(m.RemoveDateAndTime, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotDisplayPageBoundaries"}:
				m.DoNotDisplayPageBoundaries = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDisplayPageBoundaries, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "displayBackgroundShape"}:
				m.DisplayBackgroundShape = NewCT_OnOff()
				if err := d.DecodeElement(m.DisplayBackgroundShape, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printPostScriptOverText"}:
				m.PrintPostScriptOverText = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintPostScriptOverText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printFractionalCharacterWidth"}:
				m.PrintFractionalCharacterWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFractionalCharacterWidth, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printFormsData"}:
				m.PrintFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFormsData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "embedTrueTypeFonts"}:
				m.EmbedTrueTypeFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedTrueTypeFonts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "embedSystemFonts"}:
				m.EmbedSystemFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedSystemFonts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveSubsetFonts"}:
				m.SaveSubsetFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveSubsetFonts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveFormsData"}:
				m.SaveFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveFormsData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "mirrorMargins"}:
				m.MirrorMargins = NewCT_OnOff()
				if err := d.DecodeElement(m.MirrorMargins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "alignBordersAndEdges"}:
				m.AlignBordersAndEdges = NewCT_OnOff()
				if err := d.DecodeElement(m.AlignBordersAndEdges, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bordersDoNotSurroundHeader"}:
				m.BordersDoNotSurroundHeader = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundHeader, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bordersDoNotSurroundFooter"}:
				m.BordersDoNotSurroundFooter = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "gutterAtTop"}:
				m.GutterAtTop = NewCT_OnOff()
				if err := d.DecodeElement(m.GutterAtTop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hideSpellingErrors"}:
				m.HideSpellingErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideSpellingErrors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hideGrammaticalErrors"}:
				m.HideGrammaticalErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideGrammaticalErrors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "activeWritingStyle"}:
				tmp := NewCT_WritingStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ActiveWritingStyle = append(m.ActiveWritingStyle, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofState"}:
				m.ProofState = NewCT_Proof()
				if err := d.DecodeElement(m.ProofState, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "formsDesign"}:
				m.FormsDesign = NewCT_OnOff()
				if err := d.DecodeElement(m.FormsDesign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "attachedTemplate"}:
				m.AttachedTemplate = NewCT_Rel()
				if err := d.DecodeElement(m.AttachedTemplate, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "linkStyles"}:
				m.LinkStyles = NewCT_OnOff()
				if err := d.DecodeElement(m.LinkStyles, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "stylePaneFormatFilter"}:
				m.StylePaneFormatFilter = NewCT_StylePaneFilter()
				if err := d.DecodeElement(m.StylePaneFormatFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "stylePaneSortMethod"}:
				m.StylePaneSortMethod = NewCT_StyleSort()
				if err := d.DecodeElement(m.StylePaneSortMethod, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "documentType"}:
				m.DocumentType = NewCT_DocType()
				if err := d.DecodeElement(m.DocumentType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "mailMerge"}:
				m.MailMerge = NewCT_MailMerge()
				if err := d.DecodeElement(m.MailMerge, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "revisionView"}:
				m.RevisionView = NewCT_TrackChangesView()
				if err := d.DecodeElement(m.RevisionView, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "trackRevisions"}:
				m.TrackRevisions = NewCT_OnOff()
				if err := d.DecodeElement(m.TrackRevisions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotTrackMoves"}:
				m.DoNotTrackMoves = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackMoves, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotTrackFormatting"}:
				m.DoNotTrackFormatting = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackFormatting, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "documentProtection"}:
				m.DocumentProtection = NewCT_DocProtect()
				if err := d.DecodeElement(m.DocumentProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "autoFormatOverride"}:
				m.AutoFormatOverride = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoFormatOverride, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "styleLockTheme"}:
				m.StyleLockTheme = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockTheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "styleLockQFSet"}:
				m.StyleLockQFSet = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockQFSet, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "defaultTabStop"}:
				m.DefaultTabStop = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DefaultTabStop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "autoHyphenation"}:
				m.AutoHyphenation = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoHyphenation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "consecutiveHyphenLimit"}:
				m.ConsecutiveHyphenLimit = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.ConsecutiveHyphenLimit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hyphenationZone"}:
				m.HyphenationZone = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.HyphenationZone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotHyphenateCaps"}:
				m.DoNotHyphenateCaps = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotHyphenateCaps, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "showEnvelope"}:
				m.ShowEnvelope = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowEnvelope, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "summaryLength"}:
				m.SummaryLength = NewCT_DecimalNumberOrPrecent()
				if err := d.DecodeElement(m.SummaryLength, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "clickAndTypeStyle"}:
				m.ClickAndTypeStyle = NewCT_String()
				if err := d.DecodeElement(m.ClickAndTypeStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "defaultTableStyle"}:
				m.DefaultTableStyle = NewCT_String()
				if err := d.DecodeElement(m.DefaultTableStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "evenAndOddHeaders"}:
				m.EvenAndOddHeaders = NewCT_OnOff()
				if err := d.DecodeElement(m.EvenAndOddHeaders, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookFoldRevPrinting"}:
				m.BookFoldRevPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldRevPrinting, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookFoldPrinting"}:
				m.BookFoldPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldPrinting, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookFoldPrintingSheets"}:
				m.BookFoldPrintingSheets = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.BookFoldPrintingSheets, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridHorizontalSpacing"}:
				m.DrawingGridHorizontalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalSpacing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridVerticalSpacing"}:
				m.DrawingGridVerticalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalSpacing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "displayHorizontalDrawingGridEvery"}:
				m.DisplayHorizontalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayHorizontalDrawingGridEvery, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "displayVerticalDrawingGridEvery"}:
				m.DisplayVerticalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayVerticalDrawingGridEvery, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotUseMarginsForDrawingGridOrigin"}:
				m.DoNotUseMarginsForDrawingGridOrigin = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseMarginsForDrawingGridOrigin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridHorizontalOrigin"}:
				m.DrawingGridHorizontalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalOrigin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridVerticalOrigin"}:
				m.DrawingGridVerticalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalOrigin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotShadeFormData"}:
				m.DoNotShadeFormData = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotShadeFormData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noPunctuationKerning"}:
				m.NoPunctuationKerning = NewCT_OnOff()
				if err := d.DecodeElement(m.NoPunctuationKerning, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "characterSpacingControl"}:
				m.CharacterSpacingControl = NewCT_CharacterSpacing()
				if err := d.DecodeElement(m.CharacterSpacingControl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printTwoOnOne"}:
				m.PrintTwoOnOne = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintTwoOnOne, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "strictFirstAndLastChars"}:
				m.StrictFirstAndLastChars = NewCT_OnOff()
				if err := d.DecodeElement(m.StrictFirstAndLastChars, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noLineBreaksAfter"}:
				m.NoLineBreaksAfter = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksAfter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noLineBreaksBefore"}:
				m.NoLineBreaksBefore = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksBefore, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "savePreviewPicture"}:
				m.SavePreviewPicture = NewCT_OnOff()
				if err := d.DecodeElement(m.SavePreviewPicture, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotValidateAgainstSchema"}:
				m.DoNotValidateAgainstSchema = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotValidateAgainstSchema, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveInvalidXml"}:
				m.SaveInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveInvalidXml, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ignoreMixedContent"}:
				m.IgnoreMixedContent = NewCT_OnOff()
				if err := d.DecodeElement(m.IgnoreMixedContent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "alwaysShowPlaceholderText"}:
				m.AlwaysShowPlaceholderText = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysShowPlaceholderText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotDemarcateInvalidXml"}:
				m.DoNotDemarcateInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDemarcateInvalidXml, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveXmlDataOnly"}:
				m.SaveXmlDataOnly = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveXmlDataOnly, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "useXSLTWhenSaving"}:
				m.UseXSLTWhenSaving = NewCT_OnOff()
				if err := d.DecodeElement(m.UseXSLTWhenSaving, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveThroughXslt"}:
				m.SaveThroughXslt = NewCT_SaveThroughXslt()
				if err := d.DecodeElement(m.SaveThroughXslt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "showXMLTags"}:
				m.ShowXMLTags = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowXMLTags, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "alwaysMergeEmptyNamespace"}:
				m.AlwaysMergeEmptyNamespace = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysMergeEmptyNamespace, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "updateFields"}:
				m.UpdateFields = NewCT_OnOff()
				if err := d.DecodeElement(m.UpdateFields, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hdrShapeDefaults"}:
				m.HdrShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.HdrShapeDefaults, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "footnotePr"}:
				m.FootnotePr = NewCT_FtnDocProps()
				if err := d.DecodeElement(m.FootnotePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnotePr"}:
				m.EndnotePr = NewCT_EdnDocProps()
				if err := d.DecodeElement(m.EndnotePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "compat"}:
				m.Compat = NewCT_Compat()
				if err := d.DecodeElement(m.Compat, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docVars"}:
				m.DocVars = NewCT_DocVars()
				if err := d.DecodeElement(m.DocVars, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rsids"}:
				m.Rsids = NewCT_DocRsids()
				if err := d.DecodeElement(m.Rsids, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "mathPr"}:
				m.MathPr = math.NewMathPr()
				if err := d.DecodeElement(m.MathPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "attachedSchema"}:
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AttachedSchema = append(m.AttachedSchema, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "themeFontLang"}:
				m.ThemeFontLang = NewCT_Language()
				if err := d.DecodeElement(m.ThemeFontLang, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "clrSchemeMapping"}:
				m.ClrSchemeMapping = NewCT_ColorSchemeMapping()
				if err := d.DecodeElement(m.ClrSchemeMapping, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotIncludeSubdocsInStats"}:
				m.DoNotIncludeSubdocsInStats = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotIncludeSubdocsInStats, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotAutoCompressPictures"}:
				m.DoNotAutoCompressPictures = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotAutoCompressPictures, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "forceUpgrade"}:
				m.ForceUpgrade = NewCT_Empty()
				if err := d.DecodeElement(m.ForceUpgrade, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "captions"}:
				m.Captions = NewCT_Captions()
				if err := d.DecodeElement(m.Captions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "readModeInkLockDown"}:
				m.ReadModeInkLockDown = NewCT_ReadingModeInkLockDown()
				if err := d.DecodeElement(m.ReadModeInkLockDown, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "smartTagType"}:
				tmp := NewCT_SmartTagType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SmartTagType = append(m.SmartTagType, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/schemaLibrary/2006/main", Local: "schemaLibrary"}:
				m.SchemaLibrary = schemaLibrary.NewSchemaLibrary()
				if err := d.DecodeElement(m.SchemaLibrary, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "shapeDefaults"}:
				m.ShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.ShapeDefaults, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotEmbedSmartTags"}:
				m.DoNotEmbedSmartTags = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotEmbedSmartTags, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "decimalSymbol"}:
				m.DecimalSymbol = NewCT_String()
				if err := d.DecodeElement(m.DecimalSymbol, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "listSeparator"}:
				m.ListSeparator = NewCT_String()
				if err := d.DecodeElement(m.ListSeparator, &el); err != nil {
					return err
				}
			default:
				any := &gooxml.XSDAny{}
				if err := d.DecodeElement(any, &el); err != nil {
					return err
				}
				m.Extra = append(m.Extra, any)
			}
		case xml.EndElement:
			break lSettings
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Settings and its children
func (m *Settings) Validate() error {
	return m.ValidateWithPath("Settings")
}

// ValidateWithPath validates the Settings and its children, prefixing error messages with path
func (m *Settings) ValidateWithPath(path string) error {
	if err := m.CT_Settings.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
