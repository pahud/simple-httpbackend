build:
	docker build -t simple-httpbackend:latest .
tag:
	docker tag  simple-httpbackend:latest pahud/simple-httpbackend:latest
push:
	docker push pahud/simple-httpbackend:latest
