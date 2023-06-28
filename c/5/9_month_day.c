static char daytab[2][13] = {
    {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
    {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}};

// day_of_year: set dat of yeasr from month & day
int day_of_year(int year, int month, int day) {
  int i, leap;
  char *p;

  leap = year % 4 == 0 &&year % 100 != 0  || year % 400 == 0;
  p = &daytab[leap][0]; // P is a pointer to a single day, starting at 0 (illegal month)
  for (i = 1; i < month; i++)
    day += *++p;
  return day;
}

// month_day: set moht, day from day of year
void month_day(int year, int yearday, int *pmonth, int *pday)
{
  int i, leap;
  char *p;

  leap = year % 4 == 0 &&year % 100 != 0  || year % 400 == 0;
  p = &daytab[leap][0];
  for (i = 1; yearday > *++p; i++)
    yearday -= *p;
  *pmonth = i;
  *pday = yearday;
}
