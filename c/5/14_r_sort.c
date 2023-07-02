#include "functions.h"
#include <stdio.h>
#include <string.h>


/* Sort input lines */
int main(int argc, char *argv[]) {
  int nlines; /* number of input lines read */
  int numeric = 0;

  if ((nlines = readlines(lineptr, MAXLINES)) >= 0) {
    my_qsort((void **) lineptr, 0, nlines - 1, (int (*)(void*, void*))(numeric ? numcmp : strcmp));
    writelines(lineptr, nlines);
    return 0;
  } else {
    printf("TOOOOO many lines.");
    return 1;
  }
}

