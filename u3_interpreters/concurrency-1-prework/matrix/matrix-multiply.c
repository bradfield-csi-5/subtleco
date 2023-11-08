#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

/*
  A naive implementation of matrix multiplication.

  DO NOT MODIFY THIS FUNCTION, the tests assume it works correctly, which it
  currently does
*/
void matrix_multiply(double **C, double **A, double **B, int a_rows, int a_cols,
                     int b_cols) {
  for (int i = 0; i < a_rows; i++) {
    for (int j = 0; j < b_cols; j++) {
      C[i][j] = 0;
      for (int k = 0; k < a_cols; k++)
        C[i][j] += A[i][k] * B[k][j];
    }
  }
}

struct args {
  double **a;
  double **b;
  double **c;
  int start;
  int end;
  int dims;
};

void *thread_func(void *ptr);

void parallel_matrix_multiply(double **c, double **a, double **b, int a_rows,
                              int a_cols, int b_cols, int t_count) {

  pthread_t thread_id[t_count];   // store thread IDs
  int segment = a_cols / t_count; // For now, needs to divide cleanly - use a
                                  // t_count that does that.

  for (int t = 0; t < t_count; t++) {
    // build pthread args
    struct args *targs = malloc(sizeof(struct args));
    targs->a = a;
    targs->b = b;
    targs->c = c;
    targs->start = t * segment;
    targs->end = targs->start + segment;
    targs->dims = a_cols;

    pthread_create(&thread_id[t], NULL, &thread_func, targs);
  }

  // Wait for threads to finish
  for (int t = 0; t < t_count; t++) {
    pthread_join(thread_id[t], NULL);
  }
}

void *thread_func(void *targs) {
  struct args *this_args = (struct args *)targs;
  // each thread works on a set of rows
  for (int i = this_args->start; i < this_args->end; i++) {
    // each thread works on ALL columns
    for (int j = 0; j < this_args->dims; j++) {
      double sum = 0;
      //  sum over entire range of k
      for (int k = 0; k < this_args->dims; k++) {
        sum += this_args->a[i][k] * this_args->b[k][j];
      }
      this_args->c[i][j] = sum;
    }
  }
  free(targs);
  return NULL;
}
