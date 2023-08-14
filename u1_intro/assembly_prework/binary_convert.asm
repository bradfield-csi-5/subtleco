section .text
global binary_convert
binary_convert:
  movzx ecx, byte [rdi] ; grab a byte from the string
  cmp ecx, 0x00; test for null
  je END
  shl eax, 1  ; adjust the binary position
  sub ecx, '0' ; adjust for the ASCII code
  add eax, ecx ; move the result to eax
  inc rdi ; step the string pointer forward
  jmp binary_convert
END:
	ret


