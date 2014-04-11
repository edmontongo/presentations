// func Add(a, b int64) int64
TEXT ·Add(SB),4,$0
    MOVQ    x+0(FP),AX
    ADDQ    x+8(FP),AX
    MOVQ    AX,x+16(FP)
    RET
