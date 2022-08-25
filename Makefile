SHELL:=/bin/bash

ifdef test_run
	TEST_ARGS := -run $(test_run)
endif

run: check-modd-exists
	@modd -f ./.modd/server.modd.conf

check-modd-exists:
	@modd --version > /dev/null

migrate_up=go run main.go migrate --direction=up --step=0
migrate:
	@if [ "$(DIRECTION)" = "" ] || [ "$(STEP)" = "" ]; then\
    	$(migrate_up);\
	else\
		go run main.go migrate --direction=$(DIRECTION) --step=$(STEP);\
    fi