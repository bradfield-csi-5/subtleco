#include <stdio.h>
#include <string.h>

/*
Write a function expand(s1,s2) that expands shorthand notations like
a-z in the string s1 into the equivalent complete list abc...xyz in s2.
Allow for letters of either case and digits, and be prepared to handle
cases like a-b-c and a-z0-9 and -a-z. Arrange that a leading or trailing
- is taken literally.
*/

void expand(char[], char[]);

int main() {
  expand("a-z", "                          ");
  putchar('\n');
  expand("0-8", "         ");
  putchar('\n');
  expand("a-f-k", "           ");
  putchar('\n');
  expand("D-K", "        ");
  putchar('\n');
  expand("1-4-7", "       ");
  putchar('\n');
  expand("-a-z-", "                            ");
  putchar('\n');
  expand("a-z0-9", "                                    ");
  return 0;
}

// prev means a character has been seen
// start means a character and a - has been seen

void expand(char input[], char output[]) {
  int i, prev = 0, start = 0, pos = 0;
  char c, j, last;

  for (i = 0; i < strlen(input); i++) {
    c = input[i];

    if (c >= 'a' && c <= 'z') {
      prev = 1;
      if (start) {
        for (j = start; j <= c; j++) {
          if (j != last) {
            output[pos] = c;
            pos++;
            last = j;
          }
        }
        start = 0;
      }
    }

    if (c >= 'A' && c <= 'Z') {
      prev = 1;
      if (start) {
        for (j = start; j <= c; j++) {
          if (j != last) {
            output[pos] = c;
            pos++;
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
            output[pos] = c;
            pos++;
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
          output[pos] = c;
          pos++;
        }
      } else {
        output[pos] = c;
        pos++;
      }
    }
  }
  // printf("%s", output);
}
