.DEFAULT_GOAL := test

BRANCH?=master

vendor/$(BRANCH):
	mkdir -p vendor/$(BRANCH)/src/github.com/go-sql-driver/mysql
	cd vendor/$(BRANCH)/src/github.com/go-sql-driver/mysql && git clone https://github.com/go-sql-driver/mysql.git . && git reset --hard $(BRANCH)

test: vendor/$(BRANCH)
	GOPATH=$(shell pwd)/vendor/$(BRANCH) go test . 

