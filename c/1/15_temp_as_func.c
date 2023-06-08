#include <stdio.h>

int temp(int fahr);

int main() {

  int t;

  t = 50;
  printf("%d fahrenheit became %d celcius.", t, temp(t));

  return 0;
}

int temp(int fahr) {
  return ((6.0/9.0)*(fahr-32));
}
