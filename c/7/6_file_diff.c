#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char **argv) {
  FILE *file_one = fopen(argv[1], "r");
  FILE *file_two = fopen(argv[2], "r");
  char fline_one[100];
  char fline_two[100];

  while (fgets(fline_one, 100, file_one)) {
    fgets(fline_two, 100, file_two);
    if (strcmp(fline_one, fline_two)) {
      printf("there's a difference!\n %s is not\n %s\n", fline_one, fline_two);
    }
  }
}
