#include <stdio.h>

int main() {
  int c, t, b, nl;

  t = b = nl = 0;
  while ((c = getchar()) != EOF)
    if (c == '\n')
      ++nl;
    else if (c == '\t')
      ++t;
    else if (c == ' ')
      ++b;

  printf("there are %d lines up in here\n", nl);
  printf("there are %d tabs up in here\n", t);
  printf("there are %d blanks up in here\n", b);
  return 0;
}
