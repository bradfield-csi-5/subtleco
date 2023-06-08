
#include <stdio.h>

#define IN 1  /* inside of a word */
#define OUT 0 /* outside of a word */

// count the lines, words, and chars in an input
int main() {
  int c, state;

  state = OUT;

  while ((c = getchar()) != EOF) {
    if (c == ' ')
      putchar('\n');
    else 
      putchar(c);
  }

  return 0;
}
