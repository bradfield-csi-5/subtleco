/*
Naive code for multiplying two matrices together.

There must be a better way!
*/

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

// void fast_matrix_multiply(double **C, double **A, double **B, int a_rows,
//                           int a_cols, int b_cols) {
//   for (int i = 0; i < a_rows; i++) {
//     for (int j = 0; j < b_cols; j++) {
//       double sum_0 = 0, sum_1 = 0, sum_2 = 0, sum_3 = 0;
//       C[i][j] = 0;
//       int k = 0;
//       for (; k < a_cols - 3; k += 4) {
//         sum_0 += A[i][k + 0] * B[k + 0][j];
//         sum_1 += A[i][k + 1] * B[k + 1][j];
//         sum_2 += A[i][k + 2] * B[k + 2][j];
//         sum_3 += A[i][k + 3] * B[k + 3][j];
//       }
//       
//       for (; k < a_cols; k++)
//         sum_0 += A[i][k] * B[k][j];
//
//       C[i][j] += sum_0 + sum_1 + sum_2 + sum_3;
//
//     }
//   }
// }



void fast_matrix_multiply(double **C, double **A, double **B, int a_rows,
                          int a_cols, int b_cols) {
  for (int i = 0; i < a_rows; i++) {
    for (int j = 0; j < b_cols; j++) {
      double sum_0 = 0, sum_1 = 0, sum_2 = 0, sum_3 = 0;
      int k = 0;
      for (; k < a_cols - 3; k += 4) {
        sum_0 += A[i][k + 0] * B[k + 0][j];
        sum_1 += A[i][k + 1] * B[k + 1][j];
        sum_2 += A[i][k + 2] * B[k + 2][j];
        sum_3 += A[i][k + 3] * B[k + 3][j];
      }
      
      for (; k < a_cols; k++)
        sum_0 += A[i][k] * B[k][j];

      C[i][k] = sum_0 + sum_1 + sum_2 + sum_3;

    }
  }
}
