# x86 Assembly Tutorial 001

.global _start
.intel_syntax
.section .text

_start:
        # write syscall
        mov %eax, 4
        mov %ebx, 1 # stdout
        lea %ecx, [message]
        mov %edx, 13 # length
        int 0x80


        # exit syscall
        mov %eax, 1
        mov %ebx, 65

        int 0x80

.section .data
message:
        .ascii "Hello World\n"
