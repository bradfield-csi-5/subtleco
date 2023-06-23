#include <stdio.h>
#include <stdlib.h>

#define MAXOP 100
#define NUMBER '0'

// 4-11: Modify getop so that it doesn't need to use ungetch.

int getop(char[]);
void push(double);
double pop(void);

/* reverse Polish calculator */
int main() {
  int type;
  double op2;
  char s[MAXOP];

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
    case '\n':
      printf("\t %.8g\n", pop());
      break;
    default:
      printf("What even is this: %s\n", s);
      break;
    }
  }
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
  static char b = 0;

  while ((s[0] = c = b ? b : getch()) == ' ' || c == '\t')
    if (b)
      b = 0;

  s[1] = '\0';
  if (!isdigit(c) && c != '.') {
    return c; /* not a number */
  }
  i = 0;
  if (isdigit(c)) { /* collect integer part */
    while (isdigit(s[++i] = c = b ? b : getch()))
      if (b)
        b = 0;
  }
  if (c == '.') { /* collect fraction part */
    while (isdigit(s[++i] = c = b ? b : getch()))
      if (b)
        b = 0;
  }
  s[i] = '\0';
  if (c != EOF)
    b = c;
  return NUMBER;
}

#define BUFSIZE 100

char buf[BUFSIZE];
int bufp = 0;

int getch(void) { return (bufp > 0) ? buf[--bufp] : getchar(); }

void ungetch(int c) {
  if (bufp >= BUFSIZE)
    printf("ungetch: too many chars\n");
  else
    buf[bufp++] = c;
}
