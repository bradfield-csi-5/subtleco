default rel

; v = pi * r * r * (h/3)
section .text
global volume

; xmm0 = r
; xmm1 = h

volume:
  movss xmm2, [pi]
  movss xmm3, [three]
  mulss xmm0, xmm0 ; square the radius
  divss xmm1, xmm3
  mulss xmm0, xmm1
  mulss xmm0, xmm2

  ret


section .rodata
  pi:
    dd 3.14159265
  three:
    dd 3.00000000
