package pkg

import (
	"fmt"

	"github.com/kaddiya/dbi/internal"
)

func (dbinpsector *DBInspectorImpl) GetDatabaseMetadata(config *DBConfig) ([]*Table, error) {
	var result []*Table
	dbCfgBldr := &internal.DBInspectorBuilderImpl{}
	internalInspector, err := dbCfgBldr.GetDbInspectorInstance(config.Protocol, config.UserName, config.Password, config.Host, config.DBName, config.SSLMode)
	if err != nil {
		return result, err
	}

	tbls, tblErr := internalInspector.GetTables()
	if tblErr != nil {
		return result, err
	}

	for _, val := range tbls {
		t := &Table{}
		t.TableName = val.TableName
		var reflectionErr error
		var colList []*internal.DbiColumns
		var constraintsList []*internal.DbiConstraints
		var keyUsages []*internal.DbiKeyUsages

		colList, reflectionErr = internalInspector.GetColumnsForTable(val.TableName)
		constraintsList, reflectionErr = internalInspector.GetConstraintsForTable(val.TableName)
		keyUsages, reflectionErr = internalInspector.GetKeyUsageForTable(val.TableName)
		if reflectionErr != nil {
			fmt.Println("Could not get the column data for " + val.TableName + " due to " + reflectionErr.Error())
		} else {

			for _, cols := range colList {
				c := &Column{}

				for _, keyUsages := range keyUsages {
					if cols.ColumnName == keyUsages.ColumnName {
						c.ColumnName = cols.ColumnName
						for _, constraints := range constraintsList {
							if constraints.ConstraintName == keyUsages.ConstraintName {
								c.ConstraintName = constraints.ConstraintName
								c.ConstraintType = constraints.ConstraintType
							}
						}
					}

				}
				t.Columns = append(t.Columns, c)
			}
		}
		result = append(result, t)
	}
	return result, nil
}
