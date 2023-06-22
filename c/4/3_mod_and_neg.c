#include <math.h>
#include <stdio.h>
#include <stdlib.h>

#define MAXOP 100
#define NUMBER '0'

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
    case '%':
      op2 = pop();
      push(fmod(pop(), op2));
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
    // printf("PUSH %f\n", f);
    val[sp++] = f;
  } else
    printf("error: Stack is full, can't push %g\n", f);
}

// pop: pop and return top value from stack
double pop(void) {
  if (sp > 0) {
    // printf("POP %f\n", val[sp - 1]);
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
  int i, c, neg = 0;

  while ((s[0] = c = getch()) == ' ' || c == '\t')
    ;

  s[1] = '\0';
  if (c == '-' && (isdigit(c = getch()))) { // check for a negative number
    ungetch(c);
    neg = 1;
  }
  if (!isdigit(c) && c != '.' && !neg) {
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
  if (c != EOF) {
    printf("END OF OP UNGETCH -- ");
    ungetch(c);
  }
  printf("TYPE is NUMBER\n");
  return NUMBER;
}

#define BUFSIZE 100

char buf[BUFSIZE];
int bufp = 0;

int getch(void) { 
  char c;
  c = (bufp > 0) ? buf[--bufp] : getchar();
  printf("GETCH %c\n", c);
  return c;
}

void ungetch(int c) {
  if (bufp >= BUFSIZE)
    printf("ungetch: too many chars\n");
  else {
    printf("UNGETCH %c\n", c);
    buf[bufp++] = c;
  }
}
