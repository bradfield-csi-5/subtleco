#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAXOP 100
#define NUMBER '0'

int getop(char[]);
void push(double);
double pop(void);
char s[MAXOP];
char *input;

/* reverse Polish calculator */
int main(int argc, char *argv[]) {
  int type;
  double op2;
  char input_string[MAXOP];

  int i;
  for (i = 0; *++argv; i++) {
    strcat(input_string, *argv);
    if (i < argc - 1) {
      strcat(input_string, " ");
    }
  }
  input = input_string;
  printf("input = %s\n", input);

  while ((type = getop(s)) != EOF) {
    switch (type) {
    case NUMBER:
      push(atof(s));
      break;
    case '+':
      push(pop() + pop());
      break;
    case '*':
      push(pop() * pop());
      break;
    case '-':
      op2 = pop();
      push(pop() - op2);
      break;
    case '/':
      op2 = pop();
      if (op2 != 0.0)
        push(pop() / op2);
      else
        printf("Stop trying to divide by 0, it's never gonna happen\n");
      break;
    default:
      printf("What even is this: %s\n", s);
      break;
    }
  }
  printf("\t %.8g\n", pop());
  return 0;
}

#define MAXVAL 100

int sp = 0;
double val[MAXVAL];

// push: push f onto value stack
void push(double f) {
  if (sp < MAXVAL) {
    val[sp++] = f;
  } else
    printf("error: Stack is full, can't push %g\n", f);
}

// pop: pop and return top value from stack
double pop(void) {
  if (sp > 0) {
    return val[--sp];
  } else {
    printf("error: stack empty\n");
    return 0.0;
  }
}

#include <ctype.h>

int getch(void);
void ungetch(int);

// getop: get next operator or numeric operand
int getop(char s[]) {
  int i, c;

  while ((s[0] = c = getch()) == ' ' || c == '\t')
    ;

  s[1] = '\0';
  if (!isdigit(c) && c != '.') {
    return c; /* not a number */
  }
  i = 0;
  if (isdigit(c)) { /* collect integer part */
    while (isdigit(s[++i] = c = getch()))
      ;
  }
  if (c == '.') { /* collect fraction part */
    while (isdigit(s[++i] = c = getch()))
      ;
  }
  s[i] = '\0';
  if (c != EOF)
    ungetch(c);
  return NUMBER;
}

#define BUFSIZE 100

char buf[BUFSIZE];
int bufp = 0;

int getch(void) {
  return (bufp > 0) ? buf[--bufp] : *input != '\0' ? *input++ : EOF;
}

void ungetch(int c) {
  if (bufp >= BUFSIZE)
    printf("ungetch: too many chars\n");
  else
    buf[bufp++] = c;
}

