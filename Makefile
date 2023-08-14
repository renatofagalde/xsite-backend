.PHONY: project_code code

# Deployment config
STAGE=dev
PROJECT_NAME=xsite-backend
PROFILE=${PROJECT_NAME}_${STAGE}
BUCKET_NAME=cfn.${PROJECT_NAME}.${STAGE}
PARAMETERS=`cat env.${STAGE}`

# Stack config
STACK_NAME=${PROJECT_NAME}-${STAGE}
TEMPLATE_NAME=template
handlers := $(shell find . -name '*main.go')
main := "main"

deps:
	@echo "\nInstalling dependencies"
	cd hello-world; echo "Inside of hello-world directory" \
		go get ./...

clean:
	@echo "\nRemoving old builds"
	rm -rf bin
	rm -rf .aws-sam
	rm -rf samconfig.toml

test:
	@echo "\nRunning unit tests"
	cd hello-world; \
		go test -short  ./...

local:
	@echo "\nServing locally"
	env ${PARAMETERS} sam local start-api

build:
	@echo "\nBuilding handlers"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o user/handler/list/ user/handler/list/main.go || exit 1; \

deploy:
	@echo "\nPackaging AWS SAM Application"
	sam build
	aws s3 mb s3://${STACK_NAME}
	aws cloudformation package --template-file ${TEMPLATE_NAME}.yaml --s3-bucket xsitebackend --output-template-file ${TEMPLATE_NAME}-output.yaml
	aws cloudformation deploy --template-file ${TEMPLATE_NAME}-output.yaml --stack-name ${STACK_NAME} --capabilities CAPABILITY_NAMED_IAM  --parameter-overrides LogLevel=infoP

describe:
	@echo "\nDescribe stack"
	aws cloudformation describe-stacks \
		--stack-name ${STACK_NAME} \
		--profile ${PROFILE} \
		--query 'Stacks[].Outputs[]'

mocks:
	@echo "\nGenerating mocks"
	mockgen -source=domain/user.go -destination=domain/mock/user_mock.go # -mock_names Repository=MockRepository

publish: clean build deploy

$(V).SILENT: