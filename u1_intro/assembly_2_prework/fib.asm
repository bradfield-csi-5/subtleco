; #include <stdio.h>
;
; int fib(int n) {
;   if (n <= 1) return n;
;
;   return fib(n-1) + fib(n-2);
; }
;
;
; int main(int argc, char const *argv[]) {
;   int r = fib(10);
;   printf("%d", r); // 55
; }

section .text
global fib
fib:
  cmp rdi, 1          ; Compare n to 1
  jg .loop            ; if (n > 1) jump to "loop"
  mov rax, rdi        ; else move n to rax and return
  ret

.loop:
  push rdi            ; store pre-call "state" of process

  dec rdi             ; prep fib(n-1)
  call fib            ; run fib(n-1)
  push rax            ; store the return value of fib(n-1), as rax is volatile 

  dec rdi             ; prep fib(n-2)
  call fib            ; run fib(n-2)
  mov rbp, rax        ; store return value of fib(n-2) to rbp, a callee saved register

  pop rax             ; retrieve return of fib(n-1)
  add rax, rbp        ; add return values of fib(n-1) and fib(n-2)
  pop rdi             ; restore pre-call "state"
  ret

