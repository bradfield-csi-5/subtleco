#include <stdio.h>
#define TAB_LENGTH 8
int main() {
  int c;
  int i;
  int pos;
  int rem;

  pos = 0;

  while ((c = getchar()) != EOF)
    if (c != '\t') {
      putchar(c);
      ++pos;
    } else {
      rem = TAB_LENGTH - pos % TAB_LENGTH;
      for (i = 0; i < rem; ++i) {
        putchar(' ');
      }
      pos = pos + rem;
    }

  return 0;
}
