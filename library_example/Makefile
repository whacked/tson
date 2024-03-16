EXTERNAL_MAKEFILE := $(wildcard ~/setup/include/Makefile)

ifneq ($(EXTERNAL_MAKEFILE),)
include $(EXTERNAL_MAKEFILE)
else
.PHONY: default
default:
	@cat $(MAKEFILE_LIST) | grep '^.\+:$$'
endif

# run the example
example:
	go run .
