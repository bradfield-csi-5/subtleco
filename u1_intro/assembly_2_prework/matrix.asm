section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: r_index
	; r8:  c_index
  imul rcx, rdx     ; skip_row_elements = r_index * cols
  add rcx, r8       ; c_index can be taken literally. Add skip_row_index to c_index to get total skip value
  imul rcx, 4       ; multiply skip value by size of int in bytes
  mov rax, rcx[rdi] ; dereference matrix with skip value size added to address
	ret               ; gimme
