run local:
	nodemon --exec go run main.go --signal SIGTERM

docker prod:
	docker build -t opportunity . -f Dockerfile.production --build-arg ATLAS_URI=$ATLAS_URI

docker:
	docker build -t opportunity . -f Dockerfile --build-arg ATLAS_URI=$ATLAS_URI

serve:
	kubectl port-forward service/authv2 8080:8080

redis:
	kubectl port-forward service/redis 6379:6379