
.PHONY: create-stack
create-stack:
	./build_stack.sh

.PHONY: create-buildpack
create-buildpack:
	./create-buildpack.sh pip3
	./create-buildpack.sh python
	./create-buildpack.sh nodejs

.PHONY: create-builder
create-builder:
	cp ./builders/builder.toml ./output/builder.toml
	./create-builder.sh

.PHONY: all
all:
	make create-stack
	make create-buildpack
	make create-builder