#include <stdio.h>

#define IN 1;
#define OUT 0;

int main() {
  int c, state, count, i, j;
  int wcount[10];
  for (i = 0; i < 10; ++i) {
    wcount[i] = 0;
  }
  while ((c = getchar()) != EOF) {
    if (c == ' ' || c == '\t' || c == '\n') {
      state = OUT;
      wcount[count]++;
      count = 0;
    } else {
      ++count;
      state = IN;
    }
  }

  for (i = 10; i > 0; i--) {
    for (j = 1; j < 10; ++j) {
      if (wcount[j] >= i)
        printf("# ");
      else
        printf("  ");
    }
    putchar('\n');
  }

  for (i = 1; i <= 10; ++i) {
    printf("%d ", i);
  }
  return 0;
}
