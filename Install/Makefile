all: docker go imagemagick java jenkins puppet ruby vars.o xvfb yq

vars.o java.o:
	@cat $*
	
docker: vars.o
	@cat $@

go: vars.o
	@cat $@

imagemagick: vars.o
	@cat $@

java: vars.o
	@cat $@

jenkins: vars.o java.o
	@cat $@

puppet: vars.o
	@cat $@

ruby: vars.o
	@cat $@

xvfb: vars.o
	@cat $@

yq: vars.o
	@cat $@
