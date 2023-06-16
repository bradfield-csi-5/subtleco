#include <stdio.h>

int main() {
  char in[] = "hello", out[] = "     ";
  int i;
  for (i = 0; i < 6; i++){
    out[i] = in[i];
  }
  printf("%s", out);

  return 0;
}
