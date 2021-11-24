package main

// type appContext struct {
// 	sess  *session.Session
// 	dbSvc *dynamodb.DynamoDB
// }
// func (a *App) start() {
// 	// log.Println("Server is Running on port 5000")
// 	// a.r.HandleFunc("/health", "", "")
// }
func main() {
	listTable()
	getSecret()
}
