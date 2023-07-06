#include "functions.h"
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct tnode {
  char *word;
  int count;
  struct tnode *left;
  struct tnode *right;
};

#define MAXWORD 100
struct tnode *addtree(struct tnode *, char *);
void treeprint(struct tnode *);
int getword(char *, int);


/* word frequency count */
int main() {
  struct tnode *root;
  char word[MAXWORD];

  root = NULL;
  while (getword(word, MAXWORD) != EOF)
    if (isalpha(word[0]))
      root = addtree(root, word);
  treeprint(root);
  return 0;
}

struct tnode *talloc(void);
char *my_strdup(char *);

struct tnode *addtree(struct tnode *p, char *w) {
  int cond;

  if (p == NULL) {
    p = talloc();
    p->word = my_strdup(w);
    p->count = 1;
    p->left = p->right = NULL;
  } else if ((cond = strcmp(w, p->word)) == 0)
    p->count++;
  else if (cond < 0) // less than into left subtree
    p->left = addtree(p->left, w);
  else
    p->right = addtree(p->right, w);
  return p;
}

/* treeprint: in-order print of tree p */
void treeprint(struct tnode *p) {
  if (p != NULL) {
    treeprint(p->left);
    printf("%4d %s\n", p->count, p->word);
    treeprint(p->right);
  }
}

/* talloc: make a tnode */
struct tnode *talloc(void) {
  return (struct tnode *) malloc(sizeof(struct tnode));
}

char *my_strdup(char *s) // make a duplicate of s 
{
  char *p;

  p = (char *) malloc(strlen(s) +1);
  if (p != NULL)
    strcpy(p, s);
  return p;
}
