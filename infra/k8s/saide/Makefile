SUBDIRS := $(dir $(wildcard */Makefile))

all: up

up:
	@for dir in $(SUBDIRS); do \
		echo "Deploying in $$dir"; \
		$(MAKE) -C $$dir up; \
	done

down:
	@for dir in $(SUBDIRS); do \
		echo "Deleting in $$dir"; \
		$(MAKE) -C $$dir down; \
	done

status:
	@for dir in $(SUBDIRS); do \
		echo "Checking status in $$dir"; \
		$(MAKE) -C $$dir status; \
	done

.PHONY: all up down status
