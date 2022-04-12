SHELL=/bin/bash -O extglob -c
#.SHELLFLAGS=""

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
root_dir := $(dir $(mkfile_path))
parser_dir := "$(root_dir)yaai/parser"

.PHONY: setup
setup:
	mkdir -p bin
	wget -O bin/antlr.jar https://www.antlr.org/download/antlr-4.9.3-complete.jar

.PHONY: antlr
antlr:
# Delete all of the generated code. This ensures that if any files
# dissapear from the generation process, we're not left with dangling
# files.
	@echo "-------------------------------------------------"
	@echo "Cleaning generated parser files"
	@rm -f $(parser_dir)/yaai*

	@echo "-------------------------------------------------"
	@echo "Running ANTLR"
	java -jar bin/antlr.jar -Dlanguage=Go -o yaai/parser yaai/yaai.g4

# The generate files' package doesn't match the directory layout, just
# move the generated content up one directory.
	@echo "-------------------------------------------------"
	@echo "Organizing generated files"
	@mv $(parser_dir)/yaai/* $(parser_dir)/
	@rm -r $(parser_dir)/yaai



.PHONY: antlr-debug
antlr-debug:
	@mkdir -p build/antlr-debug
	@java -jar bin/antlr.jar -o build/antlr-debug yaai/yaai.g4
	@cd build/antlr-debug/yaai && javac -cp ".:$(root_dir)/bin/antlr.jar" -g yaai*.java
	@echo "Run:"
	@echo "cd build/antlr-debug/yaai && java -cp  \".:$(root_dir)/bin/antlr.jar\" -Xmx500M org.antlr.v4.gui.TestRig yaai TOKEN -gui; popd"
