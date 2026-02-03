TEXT ·Add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX    // cargar primer argumento en AX
    MOVQ b+8(FP), BX    // cargar segundo argumento en BX
    ADDQ BX, AX         // AX = AX + BX
    MOVQ AX, ret+16(FP) // guardar resultado
    RET
