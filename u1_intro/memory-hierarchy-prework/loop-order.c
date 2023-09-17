#include <stdio.h>
#include <time.h>

  /*

Two different ways to loop over an array of arrays.

Spotted at:
http://stackoverflow.com/questions/9936132/why-does-the-order-of-the-loops-affect-performance-when-iterating-over-a-2d-arra

*/

void option_one() {
  int i, j;
  static int x[4000][4000];
  for (i = 0; i < 4000; i++) {
    for (j = 0; j < 4000; j++) {
      x[i][j] = i + j;
    }
  }
}

void option_two() {
  int i, j;
  static int x[4000][4000];
  for (i = 0; i < 4000; i++) {
    for (j = 0; j < 4000; j++) {
      x[j][i] = i + j;
    }
  }
}

int main() {
  clock_t start, stop;
  double slow, fast;

  start = clock();
  option_one();
  stop = clock();
  fast = (stop - start) / (double)CLOCKS_PER_SEC;

  start = clock();
  option_two();
  stop = clock();
  slow = (stop - start) / (double)CLOCKS_PER_SEC;

  printf("Fast took %f seconds\n", fast);
  printf("Slow took %f seconds\n", slow);
  return 0;
}
