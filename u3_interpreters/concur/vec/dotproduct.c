#include "vec.h"
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

// with basic threading
#define NTHREADS 10
void *thread_func(void *ptr);
pthread_mutex_t sum_mutex = PTHREAD_MUTEX_INITIALIZER;
int count;
data_t big_sum;

struct args {
  vec_ptr u;
  vec_ptr v;
  long start;
  long end;
};

data_t dotproduct(vec_ptr u, vec_ptr v) {
  long length =
      vec_length(u); // Assuming vec_length gives you the number of elements
  long segment = length / NTHREADS;   // Number of elements per thread
  long remainder = length % NTHREADS; // Remainder elements
  pthread_t thread_id[NTHREADS];

  big_sum = 0; // Reset sum for each call to dotproduct

  for (int t = 0; t < NTHREADS; t++) {
    struct args *targs =
        malloc(sizeof(struct args)); // Allocate unique args for each thread
    targs->u = u;
    targs->v = v;
    targs->start = t * segment + (t < remainder ? t : remainder);
    targs->end = targs->start + segment + (t < remainder ? 1 : 0);

    if (pthread_create(&thread_id[t], NULL, &thread_func, targs) != 0) {
      perror("Thread create failed");
      free(targs);
      return 0;
    }
  }

  // Wait for threads to finish
  for (int t = 0; t < NTHREADS; t++) {
    if (pthread_join(thread_id[t], NULL) != 0) {
      perror("Thread join failed");
      return 0;
    }
  }

  return big_sum;
}

void *thread_func(void *ptr) {
  struct args *args = (struct args *)ptr;
  data_t this_sum = 0, sum_0 = 0, sum_1 = 0, sum_2 = 0, sum_3 = 0;

  data_t *u_data = get_vec_start(args->u), *v_data = get_vec_start(args->v);
  long i;

  for (i = args->start; i < args->end - 3; i += 4) {
    sum_0 = u_data[i] * v_data[i];
    sum_1 = u_data[i + 1] * v_data[i + 1];
    sum_2 = u_data[i + 2] * v_data[i + 2];
    sum_3 = u_data[i + 3] * v_data[i + 3];
    this_sum += sum_0 + sum_1 + sum_2 + sum_3;
  }

  if (i < args->end){
    for (long j = args->end - i; j > 0; j--, i++){
      this_sum += u_data[i] * v_data[i];
    }

  }

  pthread_mutex_lock(&sum_mutex);
  big_sum += this_sum;
  pthread_mutex_unlock(&sum_mutex);

  free(args); // Free the allocated memory for args
  return NULL;
}
// plus move get vec element
// data_t dotproduct(vec_ptr u, vec_ptr v) {
//   data_t sum = 0, sum_0 = 0, sum_1 = 0, sum_2 = 0, sum_3 = 0;
//
//   data_t *u_data = get_vec_start(u), *v_data = get_vec_start(v);
//   long vec_len = vec_length(u);
//   long i;
//
//   for (i = 0; i < vec_len - 3; i += 4) {
//     // we can assume both vectors are same length
//     sum_0 = u_data[i] * v_data[i];
//     sum_1 = u_data[i + 1] * v_data[i + 1];
//     sum_2 = u_data[i + 2] * v_data[i + 2];
//     sum_3 = u_data[i + 3] * v_data[i + 3];
//
//     sum += sum_0 + sum_1 + sum_2 + sum_3;
//   }
//
//   if (i < vec_len) {
//     for (long j = vec_len - i; j > 0; j--, i++) {
//
//       sum += u_data[i] * v_data[i];
//     }
//   }
//
//   return sum;
// }

// unoptimized
// data_t dotproduct(vec_ptr u, vec_ptr v) {
//    data_t sum = 0, u_val, v_val;
//
//    for (long i = 0; i < vec_length(u); i++) { // we can assume both vectors
//         get_vec_element(u, i, &u_val);
//         get_vec_element(v, i, &v_val);
//         sum += u_val * v_val;
//    }
//    return sum;
// }

// remove vec_len from loop
//  data_t dotproduct(vec_ptr u, vec_ptr v) {
//    data_t sum = 0, u_val, v_val;
//
//   long vec_len = vec_length(u);
//
//    for (long i = 0; i < vec_len; i++) { // we can assume both vectors are
//    same length
//         get_vec_element(u, i, &u_val);
//         get_vec_element(v, i, &v_val);
//         sum += u_val * v_val;
//    }
//    return sum;
// }

// loop unroll: 2 elements
// data_t dotproduct(vec_ptr u, vec_ptr v) {
//   data_t sum = 0, u_val, v_val;
//
//   long vec_len = vec_length(u);
//   long i;
//
//   for (i = 0; i < vec_len - 1; i += 2) {
//     // we can assume both vectors are same length
//     get_vec_element(u, i, &u_val);
//     get_vec_element(v, i, &v_val);
//     sum += u_val * v_val;
//
//     get_vec_element(u, i+1, &u_val);
//     get_vec_element(v, i+1, &v_val);
//     sum += u_val * v_val;
//   }
//   if (i < vec_len) {
//     get_vec_element(u, i, &u_val);
//     get_vec_element(v, i, &v_val);
//     sum += u_val * v_val;
//
//   }
//
//     return sum;
// }

// four loop unroll
// data_t dotproduct(vec_ptr u, vec_ptr v) {
//   data_t sum = 0, u_val, v_val;
//
//   long vec_len = vec_length(u);
//   long i;
//
//   for (i = 0; i < vec_len - 3; i += 4) {
//     // we can assume both vectors are same length
//     get_vec_element(u, i, &u_val);
//     get_vec_element(v, i, &v_val);
//     sum += u_val * v_val;
//
//     get_vec_element(u, i + 1, &u_val);
//     get_vec_element(v, i + 1, &v_val);
//     sum += u_val * v_val;
//
//     get_vec_element(u, i + 2, &u_val);
//     get_vec_element(v, i + 2, &v_val);
//     sum += u_val * v_val;
//
//     get_vec_element(u, i + 3, &u_val);
//     get_vec_element(v, i + 3, &v_val);
//     sum += u_val * v_val;
//   }
//
//   if (i < vec_len) {
//     for (long j = vec_len - i; j > 0; j--, i++) {
//
//       get_vec_element(u, i, &u_val);
//       get_vec_element(v, i, &v_val);
//       sum += u_val * v_val;
//     }
//   }
//
//   return sum;
// }

// four loop unroll, multi acc
// data_t dotproduct(vec_ptr u, vec_ptr v) {
//   data_t sum = 0, sum_0 = 0, sum_1 = 0, sum_2 = 0, sum_3 = 0, u_val, v_val;
//
//   long vec_len = vec_length(u);
//   long i;
//
//   for (i = 0; i < vec_len - 3; i += 4) {
//     // we can assume both vectors are same length
//     get_vec_element(u, i, &u_val);
//     get_vec_element(v, i, &v_val);
//     sum_0 = u_val * v_val;
//
//     get_vec_element(u, i + 1, &u_val);
//     get_vec_element(v, i + 1, &v_val);
//     sum_1 = u_val * v_val;
//
//     get_vec_element(u, i + 2, &u_val);
//     get_vec_element(v, i + 2, &v_val);
//     sum_2 = u_val * v_val;
//
//     get_vec_element(u, i + 3, &u_val);
//     get_vec_element(v, i + 3, &v_val);
//     sum_3 = u_val * v_val;
//
//     sum += sum_0 + sum_1 + sum_2 + sum_3;
//   }
//
//   if (i < vec_len) {
//     for (long j = vec_len - i; j > 0; j--, i++) {
//
//       get_vec_element(u, i, &u_val);
//       get_vec_element(v, i, &v_val);
//       sum += u_val * v_val;
//     }
//   }
//
//   return sum;
// }
