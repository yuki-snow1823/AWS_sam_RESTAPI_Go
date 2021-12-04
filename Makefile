.PHONY: build
    
build:
    sam build
local:
    @make build
    sam local start-api --profile go-academy --docker-network go-academy
deploy-g:
    @make build
    sam deploy --guided --profile go-academy
deploy:
    @make build
    sam deploy --profile go-academy
delete:
    sam delete --profile go-academy
s3-deploy:
    sam deploy --s3-bucket {バケット名} --profile go-academy