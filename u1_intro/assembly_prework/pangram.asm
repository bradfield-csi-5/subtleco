section .text
global pangram
pangram:
  xor rbx, rbx
  xor rax, rax

LOOP_START:
    movzx ecx, byte [rdi]      ; Load the character
    cmp ecx, 0x00              ; Check for null terminator
    je CHECK_PANGRAM

    ; Convert to uppercase if it's a letter
    cmp ecx, 'a'
    jl CHECK_UPPER
    sub ecx, 'a'-'A'           ; Convert lowercase to uppercase
    jmp IS_LETTER

CHECK_UPPER:
    cmp ecx, 'A'
    jl NOT_LETTER
    cmp ecx, 'Z'
    jg NOT_LETTER

IS_LETTER:
    sub ecx, 'A'               ; set index to 0 for "A", etc
    BTS rbx, rcx

NOT_LETTER:
    inc rdi                    ; Move to the next character
    jmp LOOP_START

TRUE:
    mov rax, 1
    jmp END

CHECK_PANGRAM:
    cmp rbx, 0x03ffffff
    je TRUE

END:
    ret
