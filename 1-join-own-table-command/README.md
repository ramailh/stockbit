# SQL Query Answer
- `SELECT biasa."ID", biasa."Username", parent."Username" as "ParentUserName"`
- `FROM "User" as biasa`
- `LEFT JOIN "User" as parent ON parent."ID" = biasa."Parent";`
