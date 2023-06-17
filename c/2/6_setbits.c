#include <math.h>
#include <stdio.h>

int setbits(int, int, int, int);

int main() {
  int result;
  result = setbits(122, 5, 3, 147);
  printf("%d", result);
}
// Write a function that returns x with the n bits that begin at p set to the
// right most n bits of y
int setbits(int x, int p, int n, int y) {
  int x_template = 0, y_template = 0, x_adj = 0, y_adj = 0, result = 0;

  // Create cookie cutter template to remove bits from x
  while (n > 0) {
    x_template += pow(2, p);
    y_template += pow(2, n - 1);
    p--;
    n--;
  }
  printf("x_template = %d\n", x_template);
  printf("y_template = %d\n", y_template);
  x_adj = x & ~x_template;
  y_adj = y & y_template;
  y_adj <<= (p + 1);
  result = x_adj | y_adj;
  return result;
}
