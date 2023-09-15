dcu: dcd
	docker compose up

dcd: 
	docker compose down

cicd:
	dagger run go run Infrastructure/cicd/dagger/dagger.go