run: clean build docs
	./todolist.exe

build:
	go build -o ./todolist.exe

clean:
	rm -f ./todolist.exe 

docs: docs-clean
	swag init ./

docs-clean:
	rm -rf ./docs/

dr: docker-clean docker-build docker-run-db docker-run-pg-admin sleep 
	docker run -d -p 8080:8080 --name todolist todolist:dev

docker-run:
	docker run -p 8080:8080 --name todolist todolist:dev

docker-run-db:
	docker run -d -p 5432:5432 --name todolist-postgres-db -e POSTGRES_USER=spuser -e POSTGRES_PASSWORD=SPuser96 -e POSTGRES_DB=todoListDB -v /custom/mount:/var/lib/postgresql/data postgres

docker-build: 
	docker build . -t todolist:dev

docker-clean: docker-clean-db docker-stop 
	 docker container rm todolist || true 
	
docker-stop:
	docker container stop todolist || true

docker-clean-db: docker-stop-db
	docker container rm todolist-postgres-db || true 

docker-stop-db:
	docker container stop todolist-postgres-db  || true

docker-run-pg-admin:
	docker run -d -p 8888:80 --name pgadmin4_container -e PGADMIN_DEFAULT_EMAIL=pg@mail.com -e PGADMIN_DEFAULT_PASSWORD=password dpage/pgadmin4

sleep:
	sleep 7s