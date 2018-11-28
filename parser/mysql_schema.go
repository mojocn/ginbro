package parser

import "log"

//BareDatabaseSchema create a bare project
func (p *ParseEngine) BareDatabaseSchema() error {
	colID := newCol("id", "uint", "ID", "PRI", "int", "password", true)
	colUser := newCol("email", "char", "login username", "", "char(60)", "password", true)
	colPassword := newCol("password", "varchar", "login password", "", "varchar(255)", "password", true)
	colCT := newCol("created_at", "timestamp", "created_at", "", "timestamp", "password", true)
	colUT := newCol("updated_at", "timestamp", "updated_at", "", "timestamp", "password", true)

	table := newTable("users", p.OutPath, p.OutPackage, p.AuthTable, p.AuthPassword)
	table.Columns = append(table.Columns, colID, colUser, colPassword, colCT, colUT)

	p.Tables = append(p.Tables, table)
	return nil
}

//ParseDatabaseSchema get MySQL database schema
func (p *ParseEngine) ParseDatabaseSchema() error {

	sQ := "SELECT DISTINCT `TABLE_NAME` FROM `information_schema`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `COLUMN_NAME` = 'id'"
	rows, err := p.db.Query(sQ, p.MysqlDatabase)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if rows.Scan(&tableName) != nil {
			log.Println(err)
		} else {
			//create model and handler for every tableName
			s := p.parseDatabaseTableSchema(tableName)
			p.Tables = append(p.Tables, s)
		}
	}
	return nil
}

func (p *ParseEngine) parseDatabaseTableSchema(tableName string) *schema {
	//log.Println(tableName)
	rawSql := "SELECT `COLUMN_NAME`,`DATA_TYPE`,`COLUMN_COMMENT`,`COLUMN_KEY`,`COLUMN_TYPE` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND TABLE_NAME = ?"
	rows, err := p.db.Query(rawSql, p.MysqlDatabase, tableName)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	schema := newTable(tableName, p.OutPath, p.OutPackage, p.AuthTable, p.AuthPassword)
	for rows.Next() {
		var cName, dType, cComment, cKey, cType string
		if rows.Scan(&cName, &dType, &cComment, &cKey, &cType) == nil {
			c := newCol(cName, dType, cComment, cKey, cType, p.AuthPassword, schema.IsAuthTable)
			schema.Columns = append(schema.Columns, c)
		} else {
			//create model and handler for every tableName
			log.Printf("scan table %s's schema failed", tableName)
		}
	}
	return schema
}
