package ir

// IPred is an integer comparison predicate.
type IPred string

// Integer predicates.
const (
	IPredEQ  IPred = "eq"  // eq
	IPredNE  IPred = "ne"  // ne
	IPredSGE IPred = "sge" // sge
	IPredSGT IPred = "sgt" // sgt
	IPredSLE IPred = "sle" // sle
	IPredSLT IPred = "slt" // slt
	IPredUGE IPred = "uge" // uge
	IPredUGT IPred = "ugt" // ugt
	IPredULE IPred = "ule" // ule
	IPredULT IPred = "ult" // ult
)

// FPred is a floating-point comparison predicate.
type FPred string

// Floating-point predicates.
const (
	FPredFalse FPred = "false" // false
	FPredOEQ   FPred = "oeq"   // oeq
	FPredOGE   FPred = "oge"   // oge
	FPredOGT   FPred = "ogt"   // ogt
	FPredOLE   FPred = "ole"   // ole
	FPredOLT   FPred = "olt"   // olt
	FPredONE   FPred = "one"   // one
	FPredORD   FPred = "ord"   // ord
	FPredTrue  FPred = "true"  // true
	FPredUEQ   FPred = "ueq"   // ueq
	FPredUGE   FPred = "uge"   // uge
	FPredUGT   FPred = "ugt"   // ugt
	FPredULE   FPred = "ule"   // ule
	FPredULT   FPred = "ult"   // ult
	FPredUNE   FPred = "une"   // une
	FPredUNO   FPred = "uno"   // uno
)
