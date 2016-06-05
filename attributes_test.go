package undot

import "testing"

func TestParseAttributesEmpty(t *testing.T) {
	n := NewNode()
	attr := ""
	dot := ParseAttributes(attr,n)
	if len(n.Attributes) != 0 {t.Error("Attributes shoudl be empty")}
	if dot != "" {t.Error("Dot string should be empty")}
}

func TestParseAttributesSingle(t *testing.T) {
	n := NewNode()
	attr := "label=1234"
	dot := ParseAttributes(attr,n)
	if len(n.Attributes) != 1 {t.Error("Attributes should have single entry")}
	if n.Attributes["label"] == "" {t.Error("Attribute 'label' missing")}
	if n.Attributes["label"] != "1234" {t.Errorf("Attribute value does not match, found '%s'", n.Attributes["label"])}
	if dot != "" {t.Error("Dot string should be empty")}
}

func TestParseAttributesSingleQuoted(t *testing.T) {
	n := NewNode()
	attr := "label=\"1234\""
	dot := ParseAttributes(attr,n)
	if len(n.Attributes) != 1 {t.Error("Attributes should have single entry")}
	if n.Attributes["label"] == "" {t.Error("Attribute 'label' missing")}
	if n.Attributes["label"] != "1234" {t.Errorf("Attribute value does not match, found '%s'", n.Attributes["label"])}
	if dot != "" {t.Error("Dot string should be empty")}
}

func TestParseAttributesSingleQuotedMultivar(t *testing.T) {
	n := NewNode()
	attr := "label=\"1234 | abcdef | hello\""
	dot := ParseAttributes(attr,n)
	if len(n.Attributes) != 1 {t.Error("Attributes should have single entry")}
	if n.Attributes["label"] == "" {t.Error("Attribute 'label' missing")}
	if n.Attributes["label"] != "1234 | abcdef | hello" {t.Errorf("Attribute value does not match, found '%s'", n.Attributes["label"])}
	if dot != "" {t.Error("Dot string should be empty")}
}

func TestParseAttributesMultiple(t *testing.T) {
	n := NewNode()
	attr := "label=1234,type=circle,value=\"1234 | bob\""
	dot := ParseAttributes(attr,n)
	if len(n.Attributes) != 3 {t.Error("Attributes should have three entries")}
	if n.Attributes["label"] == "" {t.Error("Attribute 'label' missing")}
	if n.Attributes["label"] != "1234" {t.Errorf("Attribute value does not match, found '%s'", n.Attributes["label"])}
	if n.Attributes["type"] == "" {t.Error("Attribute 'label' missing")}
	if n.Attributes["type"] != "circle" {t.Errorf("Attribute value does not match, found '%s'", n.Attributes["circle"])}
	if n.Attributes["value"] == "" {t.Error("Attribute 'label' missing")}
	if n.Attributes["value"] != "1234 | bob" {t.Errorf("Attribute value does not match, found '%s'", n.Attributes["value"])}
	if dot != "" {t.Error("Dot string should be empty")}
}
