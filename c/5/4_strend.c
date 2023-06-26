#include <stdio.h>

// Write the function strend(s,t) which returns 1 if the string t
// occurs at the end of the string s, and zero otherwise

int strend(char *s, char *t) {
  while (*s != *t)
    s++;
  while (*s++ == *t++)
    if (*s == '\0') {
      return 1;
    }
  return 0;
}

int main() {
  char one[] = "coffee time", two[] = "time";
  printf("should be 1 :: %d\n", strend(one,two));

  char three[] = "not so much", four[] = "butter";
  printf("should be 0 :: %d\n", strend(three,four));
  return 0;
}
