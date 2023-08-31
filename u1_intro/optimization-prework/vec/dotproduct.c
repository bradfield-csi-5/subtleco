#include "vec.h"
// unoptimized
// data_t dotproduct(vec_ptr u, vec_ptr v) {
//    data_t sum = 0, u_val, v_val;
//
//    for (long i = 0; i < vec_length(u); i++) { // we can assume both vectors
//    are same length
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


// plus move get vec element
data_t dotproduct(vec_ptr u, vec_ptr v) {
  data_t sum = 0, sum_0 = 0, sum_1 = 0, sum_2 = 0, sum_3 = 0;

  data_t *u_data = get_vec_start(u), *v_data = get_vec_start(v); 
  long vec_len = vec_length(u);
  long i;

  for (i = 0; i < vec_len - 3; i += 4) {
    // we can assume both vectors are same length
    sum_0 = u_data[i] * v_data[i];
    sum_1 = u_data[i + 1] * v_data[i + 1];
    sum_2 = u_data[i + 2] * v_data[i + 2];
    sum_3 = u_data[i + 3] * v_data[i + 3];

    sum += sum_0 + sum_1 + sum_2 + sum_3;
  }

  if (i < vec_len) {
    for (long j = vec_len - i; j > 0; j--, i++) {

      sum += u_data[i] * v_data[i];
    }
  }

  return sum;
}
