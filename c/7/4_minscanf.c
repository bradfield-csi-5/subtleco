#include <stdarg.h>
#include <stdio.h>

int minscanf(char *fmt, ...) {
  va_list ap;
  char *p, *sval;
  int *ival;
  double *dval;

  int scanned = 0;

  va_start(ap, fmt); // make ap point to the first unnamed arg
  for (p = fmt; *p; p++) {
    if (*p != '%') {
      continue;
    }
    switch (*++p) {
    case 'd':
      ival = va_arg(ap, int *);
      scanned += scanf("%d", ival);
      break;
    case 'f':
      dval = va_arg(ap, double *);
      scanned += scanf("%lf", dval);
      break;
    case 's':
      sval = va_arg(ap, char *);
      scanned += scanf("%s", sval);
      break;
    }
  }
  va_end(ap);
  return scanned;
}

int main() {
  int i;
  double j;
  char string[100];
  while (minscanf("%d", &i) == 1) {
    printf("%d\n", i);
  }
  while (minscanf("%f", &j) == 1) {
    printf("%f\n", j);
  }
  while (minscanf("%s", string) == 1)
    printf("%s", string);
  return 0;
}
