#include <math.h>
#include <stdio.h>
#define MAX_SIZE sizeof(int) * 8

int rightrot(int, int);

int main() {
  int result;
  result = rightrot(150, 3); // 10010110
  printf("result = %d", result);
}

int rightrot(int x, int n) {
  int result, cache, start;
  cache = x & (((int)pow(2, n)) - 1);
  x = x >> n;
  printf("cache = %d\n", cache);
  start = cache << (MAX_SIZE - n);
  printf("start = %d\n", start);
  result = x | start;
  return result;
}
