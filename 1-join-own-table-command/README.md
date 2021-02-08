# SQL Query Answer
* `SELECT `
* ` 	biasa."ID", biasa."Username", parent."Username" as "ParentUserName"`
* `FROM `
* `     "User" biasa`
* `LEFT JOIN`
* `     "User" parent ON parent."ID" = biasa."Parent";`

