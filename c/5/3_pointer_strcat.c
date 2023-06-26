#include <stdio.h>

// Write a pointer version of this:
// strcat: concatenate t to end of s; s must be big enough;
void old_strcat(char s[], char t[]) {
  int i, j;

  i = j = 0;
  while (s[i] != '\0')
    i++;
  while ((s[i++] = t[j++]) != '\0') // Copy t
    ;
}

void p_strcat(char *s, char *t) {
  while (*s)            // find the end of the first string
    s++;                // This must be in the loop to not pass '\0'
  while ((*s++ = *t++)) // copy t until t is \0
    ;
}

int main() {
  char og[100] = "hello", new[] = " there";
  old_strcat(og, new);
  printf("old style :: %s\n", og);

  char new_og[100] = "friends", new_new[] = " forever";
  p_strcat(new_og, new_new);
  printf("new style :: %s\n", new_og);

  return 0;
}
