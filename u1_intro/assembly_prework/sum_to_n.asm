section .text
global sum_to_n

sum_to_n:
  mov rax, 0x00
  mov rsi, 0x00
LOOP:
  cmp rsi, rdi
  jg END
  add rax, rsi
  inc rsi
  jmp LOOP
END:
  ret
