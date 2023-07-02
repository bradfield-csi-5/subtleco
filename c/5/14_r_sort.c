#include "functions.h"
#include <stdio.h>
#include <string.h>

/* Sort input lines */
int main(int argc, char *argv[]) {
  int nlines; /* number of input lines read */
  int c, numeric = 0, reverse = 0;

  while (--argc > 0 && (*++argv)[0] == '-')
    while ((c = *++argv[0]))
      switch (c) {
      case 'n':
        numeric = 1;
        break;
      case 'r':
        reverse = 1;
        break;
      default:
        printf("BAD");
        return 1;
      }

  if ((nlines = readlines(lineptr, MAXLINES)) >= 0) {
    my_qsort((void **)lineptr, 0, nlines - 1,
             (int (*)(void *, void *))(numeric ? numcmp : strcmp));
    if (reverse) {
      int left = 0, right = nlines - 1;
      while (left < right) {
        char *temp;
        temp = lineptr[left];
        lineptr[left++] = lineptr[right];
        lineptr[right--] = temp;
      }
    }

    writelines(lineptr, nlines);
    return 0;
  } else {
    printf("TOOOOO many lines.");
    return 1;
  }
}
