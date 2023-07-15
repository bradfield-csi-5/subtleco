#include <stdarg.h>
#include <stdio.h>

/* minprintf: minimal printf with variable argument list */
void minprintf(char *fmt, ...) {
  va_list ap; // points to each unnamed arg in turn
  char *p, *sval;
  int ival;
  double dval;

  va_start(ap, fmt); // make ap point to the first unnamed arg
  for (p = fmt; *p; p++) {
    if (*p != '%') {
      putchar(*p);
      continue;
    }
    switch (*++p) {
    case 'd':
      ival = va_arg(ap, int);
      printf("%d", ival);
      break;
    case 'f':
      dval = va_arg(ap, double);
      printf("%f", dval);
      break;
    case 's':
      for (sval = va_arg(ap, char *); *sval; sval++)
        putchar(*sval);
      break;
    default:
      putchar(*p);
      break;
    }
  }
  va_end(ap);
}

int main() {
  int num = 23;
  char *string = "Hello, world!";
  double d = 3.141592653589793238462643383279502884197169399375105820974944592307816406286;
  minprintf("%d\n", num);
  minprintf("%s\n", string);
  minprintf("%f\n", d);
  minprintf("%d %s %f\n", num, string, d, num);
  return 0;
}
