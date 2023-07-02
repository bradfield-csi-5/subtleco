#include "functions.h"
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *lineptr[MAXLINES];

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

/* numcmp: compare s1 and s2 numerically */
int numcmp(char *s1, char *s2)
{
  double v1, v2;

  v1 = atof(s1);
  v2 = atof(s2);
  if (v1 < v2)
    return -1;
  else
    return 0;
}

/* swap, with a void pointer to handle any type of pointer */
void swap(void *v[], int i, int j)
{
  void *temp;
  temp = v[i];
  v[i] = v[j];
  v[j] = temp;
}

/* writelines: write output lines */
void writelines(char *lineptr[], int nlines)
{
  while(nlines-- > 0)
    printf("%s\n", *lineptr++);
}

/* my_qsort: sort v[left]...v[right] into increasing order */
void my_qsort(void *v[], int left, int right, int (*comp)(void *, void *))
{
  int i, last;
  void swap(void *v[], int i, int j);

  if (left >= right) // do nothing if the array is less than 2 items
    return;
  swap(v, left, (left + right)/2);
  last = left;
  for (i = left+1; i <= right; i++)
    if ((*comp)(v[i], v[left]) < 0)
      swap(v, ++last, i);
  swap(v, left, last);
  my_qsort(v, left, last-1, comp);
  my_qsort(v, last +1, right, comp);
}
