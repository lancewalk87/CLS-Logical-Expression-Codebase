CC = gcc

#compiler flags:
CFLAGS = -g Wall

TARGET = myprog

all: $(TARGET)

$(TARGET): $(TARGET).c
  	$(CC) $(CFLAGS) -o $(TARGET) $(TARGET).c

clean:
  $(RM) $(TARGET)
