package main

type Type string

const (
	TypeString  Type = "string"
	TypeNumeric Type = "numeric"
)

type Header struct {
	Title       string            `json:"title"`
	SecondTitle string            `json:"secondTitle,omitempty"`
	SheetName   string            `json:"sheetName,omitempty"`
	IndexName   string            `json:"indexName,omitempty"`
	FreezeCol   int               `json:"freezeCol,omitempty"`
	Columns     map[string]Column `json:"columns"`
}

type Column struct {
	Name    string            `json:"name"`
	Type    Type              `json:"type"`
	Merge   bool              `json:"merge,omitempty"`
	Columns map[string]Column `json:"columns,omitempty"`
}

type Excel struct {
	Header Header                   `json:"header"`
	Data   []map[string]interface{} `json:"data"`
}
