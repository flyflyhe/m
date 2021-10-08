"".foo STEXT size=298 args=0x10 locals=0x48 funcid=0x0
	0x0000 00000 (pro.go:3)	TEXT	"".foo(SB), ABIInternal, $72-16
	0x0000 00000 (pro.go:3)	MOVQ	TLS, CX
	0x0009 00009 (pro.go:3)	PCDATA	$0, $-2
	0x0009 00009 (pro.go:3)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (pro.go:3)	PCDATA	$0, $-1
	0x0010 00016 (pro.go:3)	CMPQ	SP, 16(CX)
	0x0014 00020 (pro.go:3)	PCDATA	$0, $-2
	0x0014 00020 (pro.go:3)	JLS	288
	0x001a 00026 (pro.go:3)	PCDATA	$0, $-1
	0x001a 00026 (pro.go:3)	SUBQ	$72, SP
	0x001e 00030 (pro.go:3)	MOVQ	BP, 64(SP)
	0x0023 00035 (pro.go:3)	LEAQ	64(SP), BP
	0x0028 00040 (pro.go:3)	FUNCDATA	$0, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
	0x0028 00040 (pro.go:3)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0028 00040 (pro.go:4)	MOVQ	$11, "".foo_val1+48(SP)
	0x0031 00049 (pro.go:5)	MOVQ	$12, "".foo_val2+40(SP)
	0x003a 00058 (pro.go:6)	LEAQ	type.int(SB), AX
	0x0041 00065 (pro.go:6)	MOVQ	AX, (SP)
	0x0045 00069 (pro.go:6)	PCDATA	$1, $0
	0x0045 00069 (pro.go:6)	CALL	runtime.newobject(SB)
	0x004a 00074 (pro.go:6)	MOVQ	8(SP), AX
	0x004f 00079 (pro.go:6)	MOVQ	AX, "".&foo_val3+56(SP)
	0x0054 00084 (pro.go:6)	MOVQ	$13, (AX)
	0x005b 00091 (pro.go:7)	MOVQ	$14, "".foo_val4+32(SP)
	0x0064 00100 (pro.go:8)	MOVQ	$15, "".foo_val5+24(SP)
	0x006d 00109 (pro.go:8)	XORL	CX, CX
	0x006f 00111 (pro.go:12)	JMP	260
	0x0074 00116 (pro.go:12)	MOVQ	CX, "".i+16(SP)
	0x0079 00121 (pro.go:13)	PCDATA	$1, $1
	0x0079 00121 (pro.go:13)	CALL	runtime.printlock(SB)
	0x007e 00126 (pro.go:13)	LEAQ	"".arg_val+80(SP), AX
	0x0083 00131 (pro.go:13)	MOVQ	AX, (SP)
	0x0087 00135 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x008c 00140 (pro.go:13)	CALL	runtime.printsp(SB)
	0x0091 00145 (pro.go:13)	LEAQ	"".foo_val1+48(SP), AX
	0x0096 00150 (pro.go:13)	MOVQ	AX, (SP)
	0x009a 00154 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x009f 00159 (pro.go:13)	NOP
	0x00a0 00160 (pro.go:13)	CALL	runtime.printsp(SB)
	0x00a5 00165 (pro.go:13)	LEAQ	"".foo_val2+40(SP), AX
	0x00aa 00170 (pro.go:13)	MOVQ	AX, (SP)
	0x00ae 00174 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00b3 00179 (pro.go:13)	CALL	runtime.printsp(SB)
	0x00b8 00184 (pro.go:13)	MOVQ	"".&foo_val3+56(SP), AX
	0x00bd 00189 (pro.go:13)	MOVQ	AX, (SP)
	0x00c1 00193 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00c6 00198 (pro.go:13)	CALL	runtime.printsp(SB)
	0x00cb 00203 (pro.go:13)	LEAQ	"".foo_val4+32(SP), AX
	0x00d0 00208 (pro.go:13)	MOVQ	AX, (SP)
	0x00d4 00212 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00d9 00217 (pro.go:13)	CALL	runtime.printsp(SB)
	0x00de 00222 (pro.go:13)	LEAQ	"".foo_val5+24(SP), AX
	0x00e3 00227 (pro.go:13)	MOVQ	AX, (SP)
	0x00e7 00231 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00ec 00236 (pro.go:13)	CALL	runtime.printnl(SB)
	0x00f1 00241 (pro.go:13)	CALL	runtime.printunlock(SB)
	0x00f6 00246 (pro.go:12)	MOVQ	"".i+16(SP), AX
	0x00fb 00251 (pro.go:12)	LEAQ	1(AX), CX
	0x00ff 00255 (pro.go:16)	MOVQ	"".&foo_val3+56(SP), AX
	0x0104 00260 (pro.go:12)	CMPQ	CX, $5
	0x0108 00264 (pro.go:12)	JLT	116
	0x010e 00270 (pro.go:16)	MOVQ	AX, "".~r1+88(SP)
	0x0113 00275 (pro.go:16)	MOVQ	64(SP), BP
	0x0118 00280 (pro.go:16)	ADDQ	$72, SP
	0x011c 00284 (pro.go:16)	RET
	0x011d 00285 (pro.go:16)	NOP
	0x011d 00285 (pro.go:3)	PCDATA	$1, $-1
	0x011d 00285 (pro.go:3)	PCDATA	$0, $-2
	0x011d 00285 (pro.go:3)	NOP
	0x0120 00288 (pro.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0125 00293 (pro.go:3)	PCDATA	$0, $-1
	0x0125 00293 (pro.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 06 01 00 00 48 83 ec 48 48 89  H;a.......H..HH.
	0x0020 6c 24 40 48 8d 6c 24 40 48 c7 44 24 30 0b 00 00  l$@H.l$@H.D$0...
	0x0030 00 48 c7 44 24 28 0c 00 00 00 48 8d 05 00 00 00  .H.D$(....H.....
	0x0040 00 48 89 04 24 e8 00 00 00 00 48 8b 44 24 08 48  .H..$.....H.D$.H
	0x0050 89 44 24 38 48 c7 00 0d 00 00 00 48 c7 44 24 20  .D$8H......H.D$ 
	0x0060 0e 00 00 00 48 c7 44 24 18 0f 00 00 00 31 c9 e9  ....H.D$.....1..
	0x0070 90 00 00 00 48 89 4c 24 10 e8 00 00 00 00 48 8d  ....H.L$......H.
	0x0080 44 24 50 48 89 04 24 e8 00 00 00 00 e8 00 00 00  D$PH..$.........
	0x0090 00 48 8d 44 24 30 48 89 04 24 e8 00 00 00 00 90  .H.D$0H..$......
	0x00a0 e8 00 00 00 00 48 8d 44 24 28 48 89 04 24 e8 00  .....H.D$(H..$..
	0x00b0 00 00 00 e8 00 00 00 00 48 8b 44 24 38 48 89 04  ........H.D$8H..
	0x00c0 24 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 20  $..........H.D$ 
	0x00d0 48 89 04 24 e8 00 00 00 00 e8 00 00 00 00 48 8d  H..$..........H.
	0x00e0 44 24 18 48 89 04 24 e8 00 00 00 00 e8 00 00 00  D$.H..$.........
	0x00f0 00 e8 00 00 00 00 48 8b 44 24 10 48 8d 48 01 48  ......H.D$.H.H.H
	0x0100 8b 44 24 38 48 83 f9 05 0f 8c 66 ff ff ff 48 89  .D$8H.....f...H.
	0x0110 44 24 58 48 8b 6c 24 40 48 83 c4 48 c3 0f 1f 00  D$XH.l$@H..H....
	0x0120 e8 00 00 00 00 e9 d6 fe ff ff                    ..........
	rel 12+4 t=17 TLS+0
	rel 61+4 t=16 type.int+0
	rel 70+4 t=8 runtime.newobject+0
	rel 122+4 t=8 runtime.printlock+0
	rel 136+4 t=8 runtime.printpointer+0
	rel 141+4 t=8 runtime.printsp+0
	rel 155+4 t=8 runtime.printpointer+0
	rel 161+4 t=8 runtime.printsp+0
	rel 175+4 t=8 runtime.printpointer+0
	rel 180+4 t=8 runtime.printsp+0
	rel 194+4 t=8 runtime.printpointer+0
	rel 199+4 t=8 runtime.printsp+0
	rel 213+4 t=8 runtime.printpointer+0
	rel 218+4 t=8 runtime.printsp+0
	rel 232+4 t=8 runtime.printpointer+0
	rel 237+4 t=8 runtime.printnl+0
	rel 242+4 t=8 runtime.printunlock+0
	rel 289+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=331 args=0x0 locals=0x50 funcid=0x0
	0x0000 00000 (pro.go:19)	TEXT	"".main(SB), ABIInternal, $80-0
	0x0000 00000 (pro.go:19)	MOVQ	TLS, CX
	0x0009 00009 (pro.go:19)	PCDATA	$0, $-2
	0x0009 00009 (pro.go:19)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (pro.go:19)	PCDATA	$0, $-1
	0x0010 00016 (pro.go:19)	CMPQ	SP, 16(CX)
	0x0014 00020 (pro.go:19)	PCDATA	$0, $-2
	0x0014 00020 (pro.go:19)	JLS	321
	0x001a 00026 (pro.go:19)	PCDATA	$0, $-1
	0x001a 00026 (pro.go:19)	SUBQ	$80, SP
	0x001e 00030 (pro.go:19)	MOVQ	BP, 72(SP)
	0x0023 00035 (pro.go:19)	LEAQ	72(SP), BP
	0x0028 00040 (pro.go:19)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (pro.go:19)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (pro.go:20)	MOVQ	$666, "".arg_val+56(SP)
	0x0031 00049 (<unknown line number>)	NOP
	0x0031 00049 (pro.go:4)	MOVQ	$11, "".foo_val1+48(SP)
	0x003a 00058 (pro.go:5)	MOVQ	$12, "".foo_val2+40(SP)
	0x0043 00067 (pro.go:6)	MOVQ	$13, "".foo_val3+32(SP)
	0x004c 00076 (pro.go:7)	MOVQ	$14, "".foo_val4+24(SP)
	0x0055 00085 (pro.go:8)	MOVQ	$15, "".foo_val5+16(SP)
	0x005e 00094 (pro.go:8)	XORL	AX, AX
	0x0060 00096 (pro.go:12)	JMP	242
	0x0065 00101 (pro.go:12)	MOVQ	AX, "".i+8(SP)
	0x006a 00106 (pro.go:13)	PCDATA	$1, $0
	0x006a 00106 (pro.go:13)	CALL	runtime.printlock(SB)
	0x006f 00111 (pro.go:13)	LEAQ	"".arg_val+56(SP), AX
	0x0074 00116 (pro.go:13)	MOVQ	AX, (SP)
	0x0078 00120 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x007d 00125 (pro.go:13)	NOP
	0x0080 00128 (pro.go:13)	CALL	runtime.printsp(SB)
	0x0085 00133 (pro.go:13)	LEAQ	"".foo_val1+48(SP), AX
	0x008a 00138 (pro.go:13)	MOVQ	AX, (SP)
	0x008e 00142 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x0093 00147 (pro.go:13)	CALL	runtime.printsp(SB)
	0x0098 00152 (pro.go:13)	LEAQ	"".foo_val2+40(SP), AX
	0x009d 00157 (pro.go:13)	MOVQ	AX, (SP)
	0x00a1 00161 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00a6 00166 (pro.go:13)	CALL	runtime.printsp(SB)
	0x00ab 00171 (pro.go:13)	LEAQ	"".foo_val3+32(SP), AX
	0x00b0 00176 (pro.go:13)	MOVQ	AX, (SP)
	0x00b4 00180 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00b9 00185 (pro.go:13)	CALL	runtime.printsp(SB)
	0x00be 00190 (pro.go:13)	LEAQ	"".foo_val4+24(SP), AX
	0x00c3 00195 (pro.go:13)	MOVQ	AX, (SP)
	0x00c7 00199 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00cc 00204 (pro.go:13)	CALL	runtime.printsp(SB)
	0x00d1 00209 (pro.go:13)	LEAQ	"".foo_val5+16(SP), AX
	0x00d6 00214 (pro.go:13)	MOVQ	AX, (SP)
	0x00da 00218 (pro.go:13)	CALL	runtime.printpointer(SB)
	0x00df 00223 (pro.go:13)	NOP
	0x00e0 00224 (pro.go:13)	CALL	runtime.printnl(SB)
	0x00e5 00229 (pro.go:13)	CALL	runtime.printunlock(SB)
	0x00ea 00234 (pro.go:12)	MOVQ	"".i+8(SP), AX
	0x00ef 00239 (pro.go:12)	INCQ	AX
	0x00f2 00242 (pro.go:12)	CMPQ	AX, $5
	0x00f6 00246 (pro.go:12)	JLT	101
	0x00fc 00252 (pro.go:22)	MOVQ	"".foo_val3+32(SP), AX
	0x0101 00257 (pro.go:22)	MOVQ	AX, ""..autotmp_17+64(SP)
	0x0106 00262 (pro.go:22)	CALL	runtime.printlock(SB)
	0x010b 00267 (pro.go:22)	MOVQ	""..autotmp_17+64(SP), AX
	0x0110 00272 (pro.go:22)	MOVQ	AX, (SP)
	0x0114 00276 (pro.go:22)	CALL	runtime.printint(SB)
	0x0119 00281 (pro.go:22)	CALL	runtime.printsp(SB)
	0x011e 00286 (pro.go:22)	LEAQ	"".foo_val3+32(SP), AX
	0x0123 00291 (pro.go:22)	MOVQ	AX, (SP)
	0x0127 00295 (pro.go:22)	CALL	runtime.printpointer(SB)
	0x012c 00300 (pro.go:22)	CALL	runtime.printnl(SB)
	0x0131 00305 (pro.go:22)	CALL	runtime.printunlock(SB)
	0x0136 00310 (pro.go:23)	MOVQ	72(SP), BP
	0x013b 00315 (pro.go:23)	ADDQ	$80, SP
	0x013f 00319 (pro.go:23)	NOP
	0x0140 00320 (pro.go:23)	RET
	0x0141 00321 (pro.go:23)	NOP
	0x0141 00321 (pro.go:19)	PCDATA	$1, $-1
	0x0141 00321 (pro.go:19)	PCDATA	$0, $-2
	0x0141 00321 (pro.go:19)	CALL	runtime.morestack_noctxt(SB)
	0x0146 00326 (pro.go:19)	PCDATA	$0, $-1
	0x0146 00326 (pro.go:19)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 27 01 00 00 48 83 ec 50 48 89  H;a...'...H..PH.
	0x0020 6c 24 48 48 8d 6c 24 48 48 c7 44 24 38 9a 02 00  l$HH.l$HH.D$8...
	0x0030 00 48 c7 44 24 30 0b 00 00 00 48 c7 44 24 28 0c  .H.D$0....H.D$(.
	0x0040 00 00 00 48 c7 44 24 20 0d 00 00 00 48 c7 44 24  ...H.D$ ....H.D$
	0x0050 18 0e 00 00 00 48 c7 44 24 10 0f 00 00 00 31 c0  .....H.D$.....1.
	0x0060 e9 8d 00 00 00 48 89 44 24 08 e8 00 00 00 00 48  .....H.D$......H
	0x0070 8d 44 24 38 48 89 04 24 e8 00 00 00 00 0f 1f 00  .D$8H..$........
	0x0080 e8 00 00 00 00 48 8d 44 24 30 48 89 04 24 e8 00  .....H.D$0H..$..
	0x0090 00 00 00 e8 00 00 00 00 48 8d 44 24 28 48 89 04  ........H.D$(H..
	0x00a0 24 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 20  $..........H.D$ 
	0x00b0 48 89 04 24 e8 00 00 00 00 e8 00 00 00 00 48 8d  H..$..........H.
	0x00c0 44 24 18 48 89 04 24 e8 00 00 00 00 e8 00 00 00  D$.H..$.........
	0x00d0 00 48 8d 44 24 10 48 89 04 24 e8 00 00 00 00 90  .H.D$.H..$......
	0x00e0 e8 00 00 00 00 e8 00 00 00 00 48 8b 44 24 08 48  ..........H.D$.H
	0x00f0 ff c0 48 83 f8 05 0f 8c 69 ff ff ff 48 8b 44 24  ..H.....i...H.D$
	0x0100 20 48 89 44 24 40 e8 00 00 00 00 48 8b 44 24 40   H.D$@.....H.D$@
	0x0110 48 89 04 24 e8 00 00 00 00 e8 00 00 00 00 48 8d  H..$..........H.
	0x0120 44 24 20 48 89 04 24 e8 00 00 00 00 e8 00 00 00  D$ H..$.........
	0x0130 00 e8 00 00 00 00 48 8b 6c 24 48 48 83 c4 50 90  ......H.l$HH..P.
	0x0140 c3 e8 00 00 00 00 e9 b5 fe ff ff                 ...........
	rel 12+4 t=17 TLS+0
	rel 107+4 t=8 runtime.printlock+0
	rel 121+4 t=8 runtime.printpointer+0
	rel 129+4 t=8 runtime.printsp+0
	rel 143+4 t=8 runtime.printpointer+0
	rel 148+4 t=8 runtime.printsp+0
	rel 162+4 t=8 runtime.printpointer+0
	rel 167+4 t=8 runtime.printsp+0
	rel 181+4 t=8 runtime.printpointer+0
	rel 186+4 t=8 runtime.printsp+0
	rel 200+4 t=8 runtime.printpointer+0
	rel 205+4 t=8 runtime.printsp+0
	rel 219+4 t=8 runtime.printpointer+0
	rel 225+4 t=8 runtime.printnl+0
	rel 230+4 t=8 runtime.printunlock+0
	rel 263+4 t=8 runtime.printlock+0
	rel 277+4 t=8 runtime.printint+0
	rel 282+4 t=8 runtime.printsp+0
	rel 296+4 t=8 runtime.printpointer+0
	rel 301+4 t=8 runtime.printnl+0
	rel 306+4 t=8 runtime.printunlock+0
	rel 322+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
go.info."".foo$abstract SDWARFABSFCN dupok size=106
	0x0000 04 2e 66 6f 6f 00 01 01 11 61 72 67 5f 76 61 6c  ..foo....arg_val
	0x0010 00 00 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 31 00  .......foo_val1.
	0x0020 04 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 32 00 05  ......foo_val2..
	0x0030 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 33 00 06 00  .....foo_val3...
	0x0040 00 00 00 0c 66 6f 6f 5f 76 61 6c 34 00 07 00 00  ....foo_val4....
	0x0050 00 00 0c 66 6f 6f 5f 76 61 6c 35 00 08 00 00 00  ...foo_val5.....
	0x0060 00 0c 69 00 0c 00 00 00 00 00                    ..i.......
	rel 0+0 t=24 type.*int+0
	rel 0+0 t=24 type.int+0
	rel 18+4 t=31 go.info.int+0
	rel 33+4 t=31 go.info.int+0
	rel 48+4 t=31 go.info.int+0
	rel 63+4 t=31 go.info.int+0
	rel 78+4 t=31 go.info.int+0
	rel 93+4 t=31 go.info.int+0
	rel 101+4 t=31 go.info.int+0
gclocals·2589ca35330fc0fce83503f4569854a0 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
