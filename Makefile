compile:
	GOOS=linux GOARCH=amd64 go build -o build/civ6webhook.linux_amd64 civ6webhook.go 
	GOOS=linux GOARCH=arm go build -o build/civ6webhook.linux_arm civ6webhook.go 
	GOOS=linux GOARCH=arm64 go build -o build/civ6webhook.linux_arm64 civ6webhook.go 
