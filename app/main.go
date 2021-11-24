package main

import "log"

// type appContext struct {
// 	sess  *session.Session
// 	dbSvc *dynamodb.DynamoDB
// }

func (a *App) start() {
	log.Println("Server is Running on port 5000")
	var it Item
	a.r.HandleFunc("/health", GetItemJson(it)).Method("GET")
}

func main() {

}
