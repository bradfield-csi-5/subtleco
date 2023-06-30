#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAXLINES 1000
char *lineptr[MAXLINES];

char *lines[100];

int mygetline(char *, int);
int readlines(char *lineptr[], int nlines);
char *alloc(int);

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

#define MAXLEN 1000 /* max length of any input line */

/* readlines: read input lines */
int readlines(char *lineptr[], int maxlines) {
  int len, nlines;
  char *p, line[MAXLEN];

  nlines = 0;
  while ((len = mygetline(line, MAXLEN)) > 0)
    if (nlines >= maxlines || (p = alloc(len)) == NULL)
      return -1;
    else {
      line[len - 1] = '\0'; /* delete newline */
      strcpy(p, line);
      lineptr[nlines++] = p;
    }
  return nlines;
}

/* parse a single input line to \n or EOF */
int mygetline(char s[], int lim) {
  int c, i;

  for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; ++i)
    s[i] = c;

  if (c == '\n') {
    s[i] = c;
    ++i;
  }
  s[i] = '\0';
  return i;
}

#define ALLOCSIZE 10000 // size of available space

static char allocbuf[ALLOCSIZE]; // storage for alloc
static char *allocp = allocbuf;  // next free position
char *alloc(int n)               // return pointer to n characters
{
  if (allocbuf + ALLOCSIZE - allocp >= n) {
    allocp += n;
    return allocp - n;
  } else
    return 0;
}
