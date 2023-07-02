#include "functions.h"
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]) {
  printf("argc = %d\n", argc);
  printf("argv[0] = %s\n", argv[0]);
  printf("argv[1] = %s\n", argv[1]);
  int nlines, n, i;

  n = 10;

  if ( argc > 1 && argv[1][0] == '-') {
    n = atof(argv[1]) * -1; // why don't i need * here 
  }

  printf("n = %d\n", n);

  if ((nlines = readlines(lineptr, MAXLINES)) >= 0) {
    i = 0;
    while (n--) {
      printf("%s\n", lineptr[i++]);
    }
  }
  return 0;
}

