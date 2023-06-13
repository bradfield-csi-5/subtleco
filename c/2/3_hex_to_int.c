#include <stdio.h>
#include <string.h>
#include <math.h>

char thing[] = "0xa3faca";

int main() {
  int c, i, result, len, place, butts, power;
  len = strlen(thing);

  result = 0;
  for (i = len; i > 1; i--) {
    c = thing[i];
    power = pow(16, (len-i-1));
    if (c >= '0' && c <= '9') {
      result += (c - '0') * power;
    } else {
      switch (c) {
        case 'a':
        case 'A':
          result += (10) * power;
          break;
        case 'b':
        case 'B':
          result += (11) * power;
          break;
        case 'c':
        case 'C':
          result += (12) * power;
          break;
        case 'd':
        case 'D':
          result += (13) * power;
          break;
        case 'e':
        case 'E':
          result += (14) * power;
          break;
        case 'f':
        case 'F':
          result += (15) * power;
          break;
      }
    }
  }
  printf("%d", result);
  return 0;
}
