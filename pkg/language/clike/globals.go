package clike

import (
	"sort"
	"time"

	"github.com/cznic/xc"
)

var dict = xc.Dict

// Kind is a type category. Kind formally implements Type the only method
// returning a non nil value is Kind.
type Kind int

// Values of type Kind.
const (
	Undefined Kind = iota
	Void
	Ptr
	UintPtr // Type used for pointer arithmetic.
	Char
	SChar
	UChar
	Short
	UShort
	Int
	UInt
	Long
	ULong
	LongLong
	ULongLong
	Float
	Double
	LongDouble
	Bool
	FloatComplex
	DoubleComplex
	LongDoubleComplex
	Struct
	Union
	Enum
	TypedefName
	Function
	Array
	typeof

	kindMax
)
const (
	tsBits = 5 // Values [0, 16]
	tsMask = 1<<tsBits - 1
)
const (
	tsVoid            = iota //  0: "void"
	tsChar                   //  1: "char"
	tsShort                  //  2: "short"
	tsInt                    //  3: "int"
	tsLong                   //  4: "long"
	tsFloat                  //  5: "float"
	tsDouble                 //  6: "double"
	tsSigned                 //  7: "signed"
	tsUnsigned               //  8: "unsigned"
	tsBool                   //  9: "_Bool"
	tsComplex                // 10: "_Complex"
	tsStructSpecifier        // 11: StructOrUnionSpecifier: struct
	tsUnionSpecifier         // 12: StructOrUnionSpecifier: union
	tsEnumSpecifier          // 13: EnumSpecifier
	tsTypedefName            // 14: TYPEDEFNAME
	tsTypeof                 // 15: "typeof"
	tsUintptr                // 16: Pseudo type
)

func tsEncode(a ...int) (r int) {
	sort.Ints(a)
	for _, v := range a {
		r = r<<tsBits | v
	}
	return r<<1 | 1 // Bit 0 set: value is valid.
}

const (
	yyDefault           = 57460
	yyEofCode           = 57344
	ADDASSIGN           = 57346
	ALIGNOF             = 57347
	ANDAND              = 57348
	ANDASSIGN           = 57349
	ARROW               = 57350
	ASM                 = 57351
	AUTO                = 57352
	BOOL                = 57353
	BREAK               = 57354
	CASE                = 57355
	CAST                = 57356
	CHAR                = 57357
	CHARCONST           = 57358
	COMPLEX             = 57359
	CONST               = 57360
	CONSTANT_EXPRESSION = 1048577
	CONTINUE            = 57361
	DDD                 = 57362
	DEC                 = 57363
	DEFAULT             = 57364
	DIVASSIGN           = 57365
	DO                  = 57366
	DOUBLE              = 57367
	ELSE                = 57368
	ENUM                = 57369
	EQ                  = 57370
	EXTERN              = 57371
	FLOAT               = 57372
	FLOATCONST          = 57373
	FOR                 = 57374
	GEQ                 = 57375
	GOTO                = 57376
	IDENTIFIER          = 57377
	IDENTIFIER_LPAREN   = 57378
	IDENTIFIER_NONREPL  = 57379
	IF                  = 57380
	INC                 = 57381
	INLINE              = 57382
	INT                 = 57383
	INTCONST            = 57384
	LEQ                 = 57385
	LONG                = 57386
	LONGCHARCONST       = 57387
	LONGSTRINGLITERAL   = 57388
	LSH                 = 57389
	LSHASSIGN           = 57390
	MODASSIGN           = 57391
	MULASSIGN           = 57392
	NEQ                 = 57393
	NOELSE              = 57394
	NORETURN            = 57395
	NOSEMI              = 57396
	ORASSIGN            = 57397
	OROR                = 57398
	PPDEFINE            = 57399
	PPELIF              = 57400
	PPELSE              = 57401
	PPENDIF             = 57402
	PPERROR             = 57403
	PPHASH_NL           = 57404
	PPHEADER_NAME       = 57405
	PPIF                = 57406
	PPIFDEF             = 57407
	PPIFNDEF            = 57408
	PPINCLUDE           = 57409
	PPINCLUDE_NEXT      = 57410
	PPLINE              = 57411
	PPNONDIRECTIVE      = 57412
	PPNUMBER            = 57413
	PPOTHER             = 57414
	PPPASTE             = 57415
	PPPRAGMA            = 57416
	PPUNDEF             = 57417
	PREPROCESSING_FILE  = 1048576
	REGISTER            = 57418
	RESTRICT            = 57419
	RETURN              = 57420
	RSH                 = 57421
	RSHASSIGN           = 57422
	SHORT               = 57423
	SIGNED              = 57424
	SIZEOF              = 57425
	STATIC              = 57426
	STATIC_ASSERT       = 57427
	STRINGLITERAL       = 57428
	STRUCT              = 57429
	SUBASSIGN           = 57430
	SWITCH              = 57431
	TRANSLATION_UNIT    = 1048578
	TYPEDEF             = 57432
	TYPEDEFNAME         = 57433
	TYPEOF              = 57434
	UNARY               = 57435
	UNION               = 57436
	UNSIGNED            = 57437
	VOID                = 57438
	VOLATILE            = 57439
	WHILE               = 57440
	XORASSIGN           = 57441
	yyErrCode           = 57345

	yyMaxDepth = 200
	yyTabOfs   = -328
)

var (
	cwords = map[int]rune{
		dict.SID("auto"):     AUTO,
		dict.SID("_Bool"):    BOOL,
		dict.SID("break"):    BREAK,
		dict.SID("case"):     CASE,
		dict.SID("char"):     CHAR,
		dict.SID("_Complex"): COMPLEX,
		dict.SID("const"):    CONST,
		dict.SID("continue"): CONTINUE,
		dict.SID("default"):  DEFAULT,
		dict.SID("do"):       DO,
		dict.SID("double"):   DOUBLE,
		dict.SID("else"):     ELSE,
		dict.SID("enum"):     ENUM,
		dict.SID("extern"):   EXTERN,
		dict.SID("float"):    FLOAT,
		dict.SID("for"):      FOR,
		dict.SID("goto"):     GOTO,
		dict.SID("if"):       IF,
		dict.SID("inline"):   INLINE,
		dict.SID("int"):      INT,
		dict.SID("long"):     LONG,
		dict.SID("register"): REGISTER,
		dict.SID("restrict"): RESTRICT,
		dict.SID("return"):   RETURN,
		dict.SID("short"):    SHORT,
		dict.SID("signed"):   SIGNED,
		dict.SID("sizeof"):   SIZEOF,
		dict.SID("static"):   STATIC,
		dict.SID("struct"):   STRUCT,
		dict.SID("switch"):   SWITCH,
		dict.SID("typedef"):  TYPEDEF,
		dict.SID("union"):    UNION,
		dict.SID("unsigned"): UNSIGNED,
		dict.SID("void"):     VOID,
		dict.SID("volatile"): VOLATILE,
		dict.SID("while"):    WHILE,
	}

	tokConstVals = map[rune]int{
		ADDASSIGN:     dict.SID("+="),
		ALIGNOF:       dict.SID("_Alignof"),
		ANDAND:        dict.SID("&&"),
		ANDASSIGN:     dict.SID("&="),
		ARROW:         dict.SID("->"),
		ASM:           dict.SID("asm"),
		AUTO:          dict.SID("auto"),
		BOOL:          dict.SID("_Bool"),
		BREAK:         dict.SID("break"),
		CASE:          dict.SID("case"),
		CHAR:          dict.SID("char"),
		COMPLEX:       dict.SID("_Complex"),
		CONST:         dict.SID("const"),
		CONTINUE:      dict.SID("continue"),
		DDD:           dict.SID("..."),
		DEC:           dict.SID("--"),
		DEFAULT:       dict.SID("default"),
		DIVASSIGN:     dict.SID("/="),
		DO:            dict.SID("do"),
		DOUBLE:        dict.SID("double"),
		ELSE:          dict.SID("else"),
		ENUM:          dict.SID("enum"),
		EQ:            dict.SID("=="),
		EXTERN:        dict.SID("extern"),
		FLOAT:         dict.SID("float"),
		FOR:           dict.SID("for"),
		GEQ:           dict.SID(">="),
		GOTO:          dict.SID("goto"),
		IF:            dict.SID("if"),
		INC:           dict.SID("++"),
		INLINE:        dict.SID("inline"),
		INT:           dict.SID("int"),
		LEQ:           dict.SID("<="),
		LONG:          dict.SID("long"),
		LSH:           dict.SID("<<"),
		LSHASSIGN:     dict.SID("<<="),
		MODASSIGN:     dict.SID("%="),
		MULASSIGN:     dict.SID("*="),
		NEQ:           dict.SID("!="),
		ORASSIGN:      dict.SID("|="),
		OROR:          dict.SID("||"),
		PPPASTE:       dict.SID("##"),
		REGISTER:      dict.SID("register"),
		RESTRICT:      dict.SID("restrict"),
		RETURN:        dict.SID("return"),
		RSH:           dict.SID(">>"),
		RSHASSIGN:     dict.SID(">>="),
		SHORT:         dict.SID("short"),
		SIGNED:        dict.SID("signed"),
		SIZEOF:        dict.SID("sizeof"),
		STATIC:        dict.SID("static"),
		STATIC_ASSERT: dict.SID("_Static_assert"),
		STRUCT:        dict.SID("struct"),
		SUBASSIGN:     dict.SID("-="),
		SWITCH:        dict.SID("switch"),
		TYPEDEF:       dict.SID("typedef"),
		TYPEOF:        dict.SID("typeof"),
		UNION:         dict.SID("union"),
		UNSIGNED:      dict.SID("unsigned"),
		VOID:          dict.SID("void"),
		VOLATILE:      dict.SID("volatile"),
		WHILE:         dict.SID("while"),
		XORASSIGN:     dict.SID("^="),
	}

	id0                      = dict.SID("0")
	id1                      = dict.SID("1")
	idAlignof                = dict.SID("_Alignof")
	idAlignofAlt             = dict.SID("__alignof__")
	idAsm                    = dict.SID("asm")
	idAsmAlt                 = dict.SID("__asm__")
	idBuiltinClasifyType     = dict.SID("__builtin_classify_type")
	idBuiltinConstantP       = dict.SID("__builtin_constant_p")
	idBuiltinTypesCompatible = dict.SID("__builtin_types_compatible__") // Implements __builtin_types_compatible_p
	idChar                   = dict.SID("char")
	idConst                  = dict.SID("const")
	idDate                   = dict.SID("__DATE__")
	idDefined                = dict.SID("defined")
	idEmptyString            = dict.SID(`""`)
	idFile                   = dict.SID("__FILE__")
	idID                     = dict.SID("ID")
	idInlineAlt              = dict.SID("__inline__")
	idL                      = dict.SID("L")
	idLine                   = dict.SID("__LINE__")
	idMagicFunc              = dict.SID("__func__")
	idNoreturn               = dict.SID("_Noreturn")
	idPopMacro               = dict.SID("pop_macro")
	idPragma                 = dict.SID("_Pragma")
	idPushMacro              = dict.SID("push_macro")
	idRestrictAlt            = dict.SID("__restrict__")
	idSTDC                   = dict.SID("__STDC__")
	idSTDCHosted             = dict.SID("__STDC_HOSTED__")
	idSTDCMBMightNeqWc       = dict.SID("__STDC_MB_MIGHT_NEQ_WC__")
	idSTDCVersion            = dict.SID("__STDC_VERSION__")
	idSignedAlt              = dict.SID("__signed__")
	idSpace                  = dict.SID(" ")
	idStatic                 = dict.SID("static")
	idStaticAssert           = dict.SID("_Static_assert")
	idTDate                  = dict.SID(tuTime.Format("Jan _2 2006")) // The date of translation of the preprocessing translation unit.
	idTTime                  = dict.SID(tuTime.Format("15:04:05"))    // The time of translation of the preprocessing translation unit.
	idTime                   = dict.SID("__TIME__")
	idTypeof                 = dict.SID("typeof")
	idTypeofAlt              = dict.SID("__typeof__")
	idVAARGS                 = dict.SID("__VA_ARGS__")
	idVolatileAlt            = dict.SID("__volatile__")
	tuTime                   = time.Now()

	tokHasVal = map[rune]bool{
		CHARCONST:         true,
		FLOATCONST:        true,
		IDENTIFIER:        true,
		IDENTIFIER_LPAREN: true,
		INTCONST:          true,
		LONGCHARCONST:     true,
		LONGSTRINGLITERAL: true,
		PPHEADER_NAME:     true,
		PPNUMBER:          true,
		PPOTHER:           true,
		STRINGLITERAL:     true,
		TYPEDEFNAME:       true,
	}

	// Valid combinations of TypeSpecifier.Case ([0], 6.7.2, 2)
	tsValid = map[int]Kind{
		tsEncode(tsBool):                            Bool,              // _Bool
		tsEncode(tsChar):                            Char,              // char
		tsEncode(tsComplex):                         DoubleComplex,     // _Complex
		tsEncode(tsDouble):                          Double,            // double
		tsEncode(tsDouble, tsComplex):               DoubleComplex,     // double _Complex
		tsEncode(tsEnumSpecifier):                   Enum,              // enum specifier
		tsEncode(tsFloat):                           Float,             // float
		tsEncode(tsFloat, tsComplex):                FloatComplex,      // float _Complex
		tsEncode(tsInt):                             Int,               // int
		tsEncode(tsLong):                            Long,              // long
		tsEncode(tsLong, tsDouble):                  LongDouble,        // long double
		tsEncode(tsLong, tsDouble, tsComplex):       LongDoubleComplex, // long double _Complex
		tsEncode(tsLong, tsInt):                     Long,              // long int
		tsEncode(tsLong, tsLong):                    LongLong,          // long long
		tsEncode(tsLong, tsLong, tsInt):             LongLong,          // long long int
		tsEncode(tsShort):                           Short,             // short
		tsEncode(tsShort, tsInt):                    Short,             // short int
		tsEncode(tsSigned):                          Int,               // signed
		tsEncode(tsSigned, tsChar):                  SChar,             // signed char
		tsEncode(tsSigned, tsInt):                   Int,               // signed int
		tsEncode(tsSigned, tsLong):                  Long,              // signed long
		tsEncode(tsSigned, tsLong, tsInt):           Long,              // signed long int
		tsEncode(tsSigned, tsLong, tsLong):          LongLong,          // signed long long
		tsEncode(tsSigned, tsLong, tsLong, tsInt):   LongLong,          // signed long long int
		tsEncode(tsSigned, tsShort):                 Short,             // signed short
		tsEncode(tsSigned, tsShort, tsInt):          Short,             // signed short int
		tsEncode(tsStructSpecifier):                 Struct,            // struct
		tsEncode(tsTypedefName):                     TypedefName,       // typedef name
		tsEncode(tsTypeof):                          typeof,            // typeof name
		tsEncode(tsUintptr):                         UintPtr,           // Pseudo type.
		tsEncode(tsUnionSpecifier):                  Union,             // union
		tsEncode(tsUnsigned):                        UInt,              // unsigned
		tsEncode(tsUnsigned, tsChar):                UChar,             // unsigned char
		tsEncode(tsUnsigned, tsInt):                 UInt,              // unsigned int
		tsEncode(tsUnsigned, tsLong):                ULong,             // unsigned long
		tsEncode(tsUnsigned, tsLong, tsInt):         ULong,             // unsigned long int
		tsEncode(tsUnsigned, tsLong, tsLong):        ULongLong,         // unsigned long long
		tsEncode(tsUnsigned, tsLong, tsLong, tsInt): ULongLong,         // unsigned long long int
		tsEncode(tsUnsigned, tsShort):               UShort,            // unsigned short
		tsEncode(tsUnsigned, tsShort, tsInt):        UShort,            // unsigned short int
		tsEncode(tsVoid):                            Void,              // void
	}
)
