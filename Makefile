POSTGRES_CONTAINER=go-postgres-enum
POSTGRES_USER=postgres
POSTGRES_PASSWORD=passwd

all: build
.PHONY: all

help: ## Print this help dialog
help h:
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-20s %s\n" "target" "help" ; \
	printf "%-20s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
		IFS=$$':' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf '\033[36m'; \
		printf "%-20s %s" $$help_command ; \
		printf '\033[0m'; \
		printf "%s\n" $$help_info; \
	done
.PHONY: help

postgres-start: ## Run postgres
postgres-start: postgres-stop
	@echo "Start postgres"
	@docker run 										\
			--rm 										\
			-d 											\
			--name $(POSTGRES_CONTAINER) 				\
			-p 5432:5432 								\
			-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD)	\
			-e POSTGRES_USER=$(POSTGRES_USER)			\
			-e POSTGRES_DB=$(POSTGRES_USER) 			\
			-v $(PWD)/db:/docker-entrypoint-initdb.d 	\
			postgres:14.5
	@sleep 5
.PHONY: postgres-start

postgres-stop: ## Stops postgres if any
postgres-stop:
	@docker stop $(POSTGRES_CONTAINER) || true
.PHONY: postgres-stop

build: ## Builds example
build:
	@echo "Building example"
	@go build -mod=vendor -o bin/enums ./...
.PHONY: build

run: ## Runs example
run: postgres-start build
	@echo "Running example"
	@DB_USER=$(POSTGRES_USER) 			 \
		DB_PASSWORD=$(POSTGRES_PASSWORD) \
		DB_HOST=localhost 				 \
		DB_NAME=$(POSTGRES_USER) 		 \
		./bin/enums
.PHONY: run
