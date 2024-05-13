package cmd

import (
	"log"

	"github.com/camtrik/go-tools/internal/sql2struct"
	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql command tools",
	Long:  "sql command tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql to struct",
	Long:  "sql to struct",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}

		// query the columns info
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("Failed to get columns: %v", err)
		}

		// convert the columns to struct columns
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}

	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "database username")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "database password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "localhost:3306", "database host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "database charset")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "database type")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "database name")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "table name")

}
