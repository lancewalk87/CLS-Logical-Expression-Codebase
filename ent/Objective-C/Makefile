CC = gcc
SOURCES=Src/OSX/main.m
FRAMEWORKS:= -framework Foundation -framework Cocoa -framework AppKit
LIBRARIES:= -lobjc
CFLAGS=-Wall -Werror -arch i386 -g -v $(SOURCES)
LDFLAGS=$(LIBRARIES) $(FRAMEWORKS)
OUT=-o Build/main

all: $(SOURCES) $(OUT)

$(OUT): $(OBJECTS)
  $(CC) -o $(OBJECTS) $@ $(CFLAGS) $(LDFLAGS) $(OUT)

.m.o:
  $(CC) -c -Wall $< -o $@
