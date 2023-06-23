#include <stdio.h>
#include <string.h>

void reverse(char[], int, int);

int main() {
  char test[] = "hello";
  reverse(test, 0, (strlen(test) - 1));
  return 0;
}

void reverse(char input[], int start, int end){
  char temp;
  if (start < end) {
    temp = input[end];
    input[end--] = input[start];
    input[start++] = temp;
    reverse(input, start, end);
  } else {
    printf("%s", input);
  }
}
