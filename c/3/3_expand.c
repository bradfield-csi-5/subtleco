#include <stdio.h>
#include <string.h>

void expand(char[]);

int main() {
  expand("a-z");
  putchar('\n');
  expand("0-8");
  putchar('\n');
  expand("a-f-k");
  putchar('\n');
  expand("1-4-7");
  putchar('\n');
  expand("-a-z-");
  return 0;
}

// prev means a character has been seen
// start means a character and a - has been seen

void expand(char input[]) {
  int i, prev = 0, start = 0;
  char c, j, last;

  for (i = 0; i < strlen(input); i++) {
    c = input[i];

    if (c >= 'a' && c <= 'z') {
      prev = 1;
      if (start) {
        for (j = start; j <= c; j++) {
          if (j != last) {
            putchar(j);
            last = j;
          }
        }
        start = 0;
      }
    }

    if (c >= '0' && c <= '9') {
      prev = 1;
      if (start) {
        for (j = start; j <= c; j++) {
          if (j != last) {
            putchar(j);
            last = j;
          }
        }
        start = 0;
      }
    }

    if (c == '-') {
      if (prev) {
        start = input[i - 1];
        if (!input[i + 1]) {
          putchar(c);
        }
      } else
        putchar(c);
    }
  }
}
