package main

func main() {
	/** gin + database ==> Todo api
	 POST /todos
		body: { "status": "active", "title": "homework #2" }
		response: { id: 1, "status": "active", "title": "homework #2" }

	 GET /todos
		response: [{ id: 1, "status": "active", "title": "homework #2" }]

	 GET /todos/1
		response: { id: 1, "status": "active", "title": "homework #2" }

	PUT /todos/1
		body: { "status": "completed", "title": "homework #2" }
		response: { "status": "success" }

	DELETE /todos/1
		body: { "status": "completed", "title": "homework #2" }
		response: { "status": "deleted" }
	*/
}
