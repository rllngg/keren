package keren

type Pageable struct {
	Limit   int
	Current int
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
	Root          *Root
	QueryResult   QueryResult
	Columns       []Columns
	Limit         int
	Page          int
	Filter        string
	RenderElement *Element
	OnQuery       func(page Pageable) QueryResult
}

func NewDataTable(root *Root) *DataTable {
	return &DataTable{
		Root:    root,
		Columns: []Columns{},
		Limit:   10,
		Page:    0,
	}
}

func (table *DataTable) SetPage(page int) *DataTable {
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
func (table *DataTable) GetPagination() *Element {
	return table.Root.Div(

		table.Root.Button("Prev", "primary btn-sm").Disabled(table.Page <= 0).OnClick(func(event *Event) *Element {

			table.Page = table.Page - 1
			return table.GetTable()
		}),

		table.Root.Button("Next", "primary btn-sm").OnClick(func(event *Event) *Element {
			table.Page = table.Page + 1
			return table.GetTable()
		}),
	).Class("d-flex gap-2")
}
func (table *DataTable) GetTable() *Element {
	// create table
	theadElement := table.Root.Thead()
	for _, data := range table.Columns {
		theadElement.Append(table.Root.Th(data.Name))
	}
	tbodyElement := table.Root.Tbody()
	table.QueryResult = table.OnQuery(Pageable{
		Limit:   table.Limit,
		Current: table.Page,
	})
	table.Body(tbodyElement)
	tableElement := table.Root.Table(
		theadElement,
		tbodyElement,
	)
	if table.RenderElement != nil {
		table.RenderElement.RemoveChildren().Body(tableElement,
			table.GetPagination())
	} else {
		table.RenderElement = table.Root.Div(
			tableElement,
			table.GetPagination(),
		)
	}

	return table.RenderElement
}
func (table *DataTable) Element(triggerName string) *Element {
	// create table
	div := table.Root.Div(
		table.Root.Form(
			table.Root.Row(
				table.Root.Col(
					table.Root.TextInput("search", "search", "Search").Bind(&table.Filter),
				).AddClass("col-md-3"),
				table.Root.Col(
					table.Root.Select("limit", "Limit", [][]string{
						{"1", "1"},
						{"5", "5"},
						{"10", "10"},
						{"50", "50"},
						{"100", "100"},
					}).Bind(&table.Limit).OnChange(func(event *Event) *Element {

						return table.GetTable()
					}).AddClass("col-md-2"),
				),
			).AddClass("align-items-center"),
		).OnSubmit(func(event *Event) *Element {
			table.Page = 0
			return table.GetTable()
		}),
		table.GetTable(),
	).OnEvent(triggerName, func(event *Event) *Element {
		return table.GetTable()
	})

	return div
}
