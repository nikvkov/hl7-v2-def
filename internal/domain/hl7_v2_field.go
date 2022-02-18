package domain

type Field struct {
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	Position     string      `json:"position"`
	Length       int         `json:"length"`
	DataType     string      `json:"dataType"`
	DataTypeName string      `json:"dataTypeName"`
	Usage        string      `json:"usage"`
	Rpt          string      `json:"rpt"`
	TableId      *string     `json:"tableId"`
	TableName    *string     `json:"tableName"`
	Name         string      `json:"name"`
	Description  interface{} `json:"description"`
}
