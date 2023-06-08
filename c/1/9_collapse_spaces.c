#include <stdio.h>

#define SPACE 1
#define NS 0

int c, prev;

int main() {

  while ((c = getchar()) != EOF) {
    if (c == ' ') {
      if (prev == NS) {
        putchar(c);
      }
      prev = SPACE;
    } else {
      prev = NS;
      putchar(c);
    }
  }

  return 0;
}
