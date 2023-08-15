section .text
global sum_to_n

sum_to_n:
  xor rax, rax
  xor rsi, rsi
LOOP:
  cmp rsi, rdi
  jg END
  add rax, rsi
  inc rsi
  jmp LOOP
END:
  ret
