# MIT License
#
# Copyright (c) 2024 Alessandro Salerno
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.


CC=gcc
CFLAGS=-O2 -std=c2x -Wno-unused-parameter -Iinclude/ -flto -DHTMC_CGI_INTF -DEXT_HTMC_BUILD="\"$(shell date +%y.%m.%d)\""
BIN=bin/
EXEC=$(BIN)/htmc
LIB=$(BIN)/libhtmc.a

rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard ,$d, $2) $(filter $(subst *, %, $2),$d))
SRC=$(call rwildcard, src, *.c)
OBJ=$(patsubst src/%.c,obj/%.o, $(SRC))
RELOC_OBJ=$(patsubst src/%.c,lib/%.o, $(SRC))

.PHONEY: all
all: obj $(OBJ) $(RELOC_OBJ) $(EXEC) $(LIB)

.PHONEY: htmc
htmc: $(EXEC)

.PHONEY: libhtmc
libhtmc: $(LIB)

$(EXEC): obj $(OBJ)
	$(CC) -flto $(OBJ) -o $(EXEC)

$(LIB): obj $(RELOC_OBJ)
	ar rcs $(LIB) $(RELOC_OBJ)

obj/%.o: src/%.c
	@mkdir -p $(@D)
	$(CC) $(CFLAGS) -c $^ -o $@

lib/%.o: src/%.c
	@mkdir -p $(@D)
	$(CC) $(CFLAGS) -fPIC -g -w -c $^ -o $@

obj:
	@mkdir -p obj/
	@mkdir -p tmp/
	@mkdir -p $(BIN)
	@mkdir -p lib/

clean:
	rm -rf obj/; \
	rm -rf tmp/; \
	rm -rf $(BIN); \
	rm -rf lib/
