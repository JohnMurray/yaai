.PHONY: setup
setup:
	mkdir -p bin
	wget -O bin/antlr.jar https://www.antlr.org/download/antlr-4.9.3-complete.jar

.PHONY: antlr
antlr:
	java -jar bin/antlr.jar -Dlanguage=Go -o yaai/parser yaai/yaai.g4