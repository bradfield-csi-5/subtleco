#include <stdio.h>

void itoa(int);

int main() {
  int res;
  res = 123;
  itoa(res);
}

void itoa(int n) {
  char c;
  if (n / 10) {
    itoa(n / 10);
  }
  c = (n % 10 + '0');
  putchar(c);
}
