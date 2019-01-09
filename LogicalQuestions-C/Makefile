CC      = gcc
CFLAGS  = -g
RM      = rm -f

default: all

all: Hello

Hello: Hello.c
    $(CC) $(CFLAGS) -o Hello Hello.c

clean veryclean:
    $(RM) Hello
