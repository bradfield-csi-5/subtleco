#include <stdio.h>
#include <string.h>
#define NUMBER 0
#define LETTER 1

void expand(char[]);

int main() {
  expand("a-z");
  return 0;
}

void expand(char input[]) {
  int i, prev, start, type;
  char c, j;
  for (i = 0; i < strlen(input); i++) {
    c = input[i];
    if (c >= 'a' && c <= 'z') {
      type = LETTER;
      prev = 1;
      if (start) {
        for (j = start; j <= c; j++) {
          putchar(j);
        }
      }
    }

    if (c >= '0' && c <= '9') {
      type = NUMBER;
    }

    if (c == '-' && prev) {
      start = input[i - 1];
    }
  }
}
