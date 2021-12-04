package main

func SelectAll(db *sql.DB, request events.APIGatewayProxyRequest) {

	rows, err := db.Query("select id, content from tasks")
	if err !=nil {
		log.Fatal(err)
	}
	var resultTask []Task
	for rows.Next() {
		task := Task{}
		if err := rows.Scan(&task.Id, &task.Content); err != nil {
			log.Fatal(err)
		}
		resultTask = append(resultTask, task)
	}
	for _, v := range resultTask {
		log.Printf("id: %d content:%s", v.Id, v.Content)
	} 
}