package domain

type Diff interface{ }

func (s *Schema) Diff(originalSchema *Schema) []Diff {
	diffsQueue := []Diff{}

	for tableName, table := range s.Tables {
		originalTable, originalHasTable := originalSchema.Tables[tableName]
		if !originalHasTable {
			diffsQueue = append(diffsQueue, NewCreateTableDiff(table))
		} else {
			diffsQueue = append(diffsQueue, table.Diff(originalTable)...)
		}
	}

	for tableName := range originalSchema.Tables {
		if _, ok := s.Tables[tableName]; !ok {
			diffsQueue = append(diffsQueue, NewDropTableDiff(tableName))
		}
	}

	return diffsQueue
}

func (t *Table) Diff(originalTable *Table) []Diff {
	diffsQueue := []Diff{}

	for columnName, column := range t.Columns {
		originalColumn, originalHasColumn := originalTable.Columns[columnName]
		if !originalHasColumn {
			diffsQueue = append(diffsQueue, NewCreateColumnDiff(t.Name, column))
		} else {
			diffsQueue = append(diffsQueue, column.Diff(originalColumn)...)
		}
	}

	for columnName := range originalTable.Columns {
		if _, ok := t.Columns[columnName]; !ok {
			diffsQueue = append(diffsQueue, NewDropColumnDiff(t.Name, columnName))
		}
	}

	return diffsQueue
}

func (t *Column) Diff(originalColumn *Column) []Diff {
	return nil
}
