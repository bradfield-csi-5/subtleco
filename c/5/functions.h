#ifndef FUNCTIONS_H
#define FUNCTIONS_H

#define MAXLINES 1000
#define MAXLEN 1000 
#define ALLOCSIZE 10000 

extern char *lineptr[]; 

int mygetline(char *, int);
int readlines(char *lineptr[], int nlines);
void writelines(char *lineptr[], int nlines);
void my_qsort(void *lineptr[], int, int, int (*comp)(void *, void *));
char *alloc(int);
int numcmp(char *, char *);
void swap(void *[], int, int);

#endif // FUNCTIONS_H
