docker_build:
	docker build . -t template

docker_build_test:
	docker build . -t template_test --target=test

docker_run: docker_build_test
	docker-compose up

graphiql:
	docker run -it --rm -p 4040:4040 friendsofgo/graphiql