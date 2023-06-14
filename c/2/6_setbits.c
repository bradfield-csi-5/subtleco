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
  int template, x_adj, y_adj, result;
  while (n > 0) {
    template += pow(2,p);
    p--;
    n--;
  }
  x_adj = x & ~template;
  printf("x_adj = %d\n", x_adj);
  y_adj = y & template;
  printf("y_adj = %d\n", y_adj);
  result = x_adj | y_adj;
  return result;
}
