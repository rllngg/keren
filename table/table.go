package table

import "github.com/erlanggatampan/keren"

type Pageable struct {
	PageCurrent int
	PageLimit   int
	PageTotal   int
}
type QueryResult struct {
	Total int
	Rows  [][]string
}
type DataTable struct {
	Root        *keren.Root
	QueryResult QueryResult
	Columns     []string
	Page        *Pageable
	Filter      string
	OnQuery     func(page Pageable) QueryResult
}

func NewDataTable(root *keren.Root) *DataTable {
	return &DataTable{
		Root:    root,
		Columns: []string{},
		Page: &Pageable{
			PageCurrent: 1,
			PageLimit:   10,
			PageTotal:   0,
		},
	}
}

func (table *DataTable) SetColumns(columns ...string) *DataTable {
	table.Columns = columns
	return table
}
func (table *DataTable) SetPage(page *Pageable) *DataTable {
	table.Page = page
	return table
}
func (table *DataTable) Body(body *keren.Element) *keren.Element {
	// create table
	body.RemoveChildren()
	for _, data := range table.QueryResult.Rows {
		trElement := table.Root.Tr()
		for i := range table.Columns {
			trElement.Append(table.Root.Td().SetInnerHTML(data[i]))
		}
		body.Append(trElement)
	}
	return body
}
func (table *DataTable) Element() *keren.Element {
	// create table
	theadElement := table.Root.Thead()
	for _, column := range table.Columns {
		theadElement.Append(table.Root.Th(column))
	}
	tbodyElement := table.Root.Tbody()
	table.QueryResult = table.OnQuery(*table.Page)
	table.Body(tbodyElement)
	tableElement := table.Root.Table(
		theadElement,
		tbodyElement,
	)
	search_input := table.Root.Input("text", "search", ".", "Search").OnChange(func(e *keren.Event) *keren.Element {
		// search
		search := e.Element.Value
		if search == "" {
			return tableElement
		}
		table.Filter = search
		// filter
		table.QueryResult = table.OnQuery(*table.Page)
		// create table
		table.Body(tbodyElement)
		return tbodyElement
	})
	div := table.Root.Div(
		search_input,
		tableElement,
	)

	return div
}
