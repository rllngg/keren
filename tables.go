package keren

type Pageable struct {
	PageCurrent int
	PageLimit   int
	PageTotal   int
}
type QueryResult struct {
	Total int
	Rows  [][]string
}
type Columns struct {
	Name     string
	Callback func(data []string) *Element
}
type DataTable struct {
	Root        *Root
	QueryResult QueryResult
	Columns     []Columns
	Page        *Pageable
	Filter      string
	OnQuery     func(page Pageable) QueryResult
}

func NewDataTable(root *Root) *DataTable {
	return &DataTable{
		Root:    root,
		Columns: []Columns{},
		Page: &Pageable{
			PageCurrent: 1,
			PageLimit:   10,
			PageTotal:   0,
		},
	}
}

func (table *DataTable) SetPage(page *Pageable) *DataTable {
	table.Page = page
	return table
}
func (table *DataTable) AddColumn(name string, callback func(data []string) *Element) *DataTable {
	table.Columns = append(table.Columns, Columns{
		Name:     name,
		Callback: callback,
	})
	return table
}
func (table *DataTable) Body(body *Element) *Element {
	// create table
	body.RemoveChildren()
	for _, data := range table.QueryResult.Rows {
		trElement := table.Root.Tr()
		for _, col := range table.Columns {
			trElement.Append(table.Root.Td(col.Callback(data)))
		}
		body.Append(trElement)
	}
	return body
}

func (table *DataTable) Element(triggerName string) *Element {
	// create table
	theadElement := table.Root.Thead()
	for _, data := range table.Columns {
		theadElement.Append(table.Root.Th(data.Name))
	}
	tbodyElement := table.Root.Tbody()
	table.QueryResult = table.OnQuery(*table.Page)
	table.Body(tbodyElement)
	tableElement := table.Root.Table(
		theadElement,
		tbodyElement,
	)
	ReloadTable := func() *Element {
		// filter
		table.QueryResult = table.OnQuery(*table.Page)
		// create table
		table.Body(tbodyElement)
		return tbodyElement
	}
	search_input := table.Root.Input("text", "search", "Search...", "Search").OnChange(func(e *Event) *Element {
		// search
		search := e.Element.Value
		if search == "" {
			return tableElement
		}
		table.Filter = search

		return ReloadTable()
	})
	div := table.Root.Div(
		search_input,
		tableElement,
	).OnEvent(triggerName, func(event *Event) *Element {
		return ReloadTable()
	})

	return div
}
