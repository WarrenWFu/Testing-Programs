CC=g++
CPPFLAGS=-pg -g -Wall
CPPFLAGS+=-I/mnt/winshared/ylink_optimize/library/libaps/include/
CPPFLAGS+=-I/mnt/winshared/ylink_optimize/library/libMQAgent/inc
LDFLAGS=-L/home/spring/project/qhjk/product/Linux64/lib
LIBS=-laps.base -laps.iconv -lMQAgent
OBJECTS=myMQ.o
TARGETS=myMQ

all:$(TARGETS)
.PHONY:all

$(TARGETS):$(OBJECTS)
	$(CC) $(LDFLAGS) $(LIBS) -o $@ $^
	@rm -f $(OBJECTS)

$(OBJECTS):%.o:%.cpp

clean:
	@rm -f $(OBJECTS) $(TARGETS)
.PHONY:clean



