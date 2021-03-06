package generator

var GetDataType = map[string]string{
	"char":      "string",
	"varchar":   "string",
	"longtext":  "string",
	"text":      "string",
	"integer":   "int",
	"tinyint":   "int",
	"int":       "int",
	"bigint":    "int64",
	"boole":     "bool",
	"decimal":   "float64",
	"datetime":  "time.Time",
	"date":      "time.Time",
	"timestamp": "time.Time",
}

const (
	Underline       string = "_"
	TAB             string = "	"
	ENTER           string = "\r\n"
	COLUMN          string = "${columnName}"
	MODEL_NAME      string = "${modelName}"
	TABLE_NAME      string = "${tableName}"
	TABLE_NAME_FUNC string = `func (${modelName}) TableName() string {
							 return "${tableName}"
						 }`
)
