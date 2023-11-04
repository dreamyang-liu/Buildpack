
.PHONY: create-stack
create-stack:
	./build_stack.sh

.PHONY: create-buildpack
create-buildpack:
	./create-buildpack.sh pip3
	./create-buildpack.sh python

.PHONY: create-builder
create-builder:
	./create-builder.sh

.PHONY: all
all:
	make create-stack
	make create-buildpack
	make create-builder