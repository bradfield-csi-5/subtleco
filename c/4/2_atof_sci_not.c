#include <ctype.h>
#include <stdio.h>

double atof(char s[]) {
  double val, power, e;
  int i, sign, op, n;

  for (i = 0; isspace(s[i]); i++)
    ;
  sign = (s[i] == '-') ? -1 : 1;
  if (s[i] == '+' || s[i] == '-')
    i++;
  for (val = 0.0; isdigit(s[i]); i++)
    val = 10.0 * val + (s[i] - '0');
  if (s[i] == '.')
    i++;
  for (power = 1.0; isdigit(s[i]); i++) {
    val = 10.0 * val + (s[i] - '0');
    power *= 10.0;
  }
  if (s[i] == 'e' || s[i] == 'E') {
    i++;
    op = (s[i] == '-') ? 0 : 1; // if the e statement includes '-', we multiply
                                // power. Else divide power.
    if (!op)
      i++;
    for (e = 0.0; isdigit(s[i]); i++) {
      e = 10.0 * e + (s[i] - '0');
    }
    for (n = 0; n < e; n++) {
      power = op ? power / 10 : power * 10;
    }
  }
  return sign * val / power;
}

int main() {
  double result;
  result = atof("123.45");
  printf("result is %f\n", result);
  result = atof("123.45e-6");
  printf("result is %f\n", result);
  result = atof("123.45e6");
  printf("result is %f\n", result);

  return 0;
}
