#include "textflag.h"

// func Sum(s []int32) int64
TEXT ·Sum(SB), NOSPLIT, $0
    MOVQ x_base+0(FP), AX
    MOVQ x_len+8(FP), CX
    XORQ R8, R8

loop:
    CMPQ CX, $0
    JE done
    MOVLQSX (AX), R9
    ADDQ R9, R8
    ADDQ $4, AX
    SUBQ $1, CX
    JMP loop

done:
    MOVQ R8, ret+24(FP)
    RET

// func Fibonacci(n uint64) uint64
TEXT ·Fibonacci(SB), NOSPLIT, $0
    MOVQ n+0(FP), AX
    XORQ R8, R8
    XORQ R9, R9
    XORQ R10, R10

    CMPQ AX, $0
    JE done_fib

    MOVQ $1, R10
    MOVQ $1, R9

loop_fib:
    CMPQ R10, AX
    JAE done_fib

    MOVQ R9, R11
    ADDQ R8, R9
    MOVQ R11, R8
    ADDQ $1, R10
    JMP loop_fib

done_fib:
    MOVQ R9, ret+8(FP)
    RET
