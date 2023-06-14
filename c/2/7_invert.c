#include <math.h>
#include <stdio.h>

int invert(int, int, int);
int setbits(int, int, int, int);

int main() {
  int result;
  result = invert(156, 5, 3);
  printf("%d", result);
}
int invert(int x, int p, int n) {
  int result;
  result = setbits(x, p, n, ~x ); // assuming 8 bits
  return result;
}

int setbits(int x, int p, int n, int y) {
  int template, x_adj, y_adj, result;
  printf("y = %d\n", y);
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
