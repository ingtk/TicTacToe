run:
	TZ=Asia/Tokyo go run main.go

mock:
	mockgen -source=api/api.go -destination=db/mock_db.go -package=db

tools:
	go install go.uber.org/mock/mockgen@latest
