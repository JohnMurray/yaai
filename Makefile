mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
root_dir := $(dir $(mkfile_path))

.PHONY: setup
setup:
	mkdir -p bin
	wget -O bin/antlr.jar https://www.antlr.org/download/antlr-4.9.3-complete.jar

.PHONY: antlr
antlr:
	java -jar bin/antlr.jar -Dlanguage=Go -o yaai/parser yaai/yaai.g4



.PHONY: antlr-debug
antlr-debug:
	@mkdir -p build/antlr-debug
	@java -jar bin/antlr.jar -o build/antlr-debug yaai/yaai.g4
	@cd build/antlr-debug/yaai && javac -cp ".:$(root_dir)/bin/antlr.jar" -g yaai*.java
	@echo "Run:"
	@echo "cd build/antlr-debug/yaai && java -cp  \".:$(root_dir)/bin/antlr.jar\" -Xmx500M org.antlr.v4.gui.TestRig yaai TOKEN -gui; popd"