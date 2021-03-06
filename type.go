package govips

// #cgo pkg-config: vips
// #include "vips/vips.h"
import "C"
import (
	"strings"
	"sync"
)

// ImageType represents an image type
type ImageType int

// ImageType enum
const (
	ImageTypeUnknown ImageType = iota
	ImageTypeGif
	ImageTypeJpeg
	ImageTypeMagick
	ImageTypePdf
	ImageTypePng
	ImageTypeSvg
	ImageTypeTiff
	ImageTypeWebp
)

var imageTypeExtensionMap = map[ImageType]string{
	ImageTypeGif:    ".gif",
	ImageTypeJpeg:   ".jpeg",
	ImageTypeMagick: ".magick",
	ImageTypePdf:    ".pdf",
	ImageTypePng:    ".png",
	ImageTypeSvg:    ".svg",
	ImageTypeTiff:   ".tiff",
	ImageTypeWebp:   ".webp",
}

// OutputExt returns the canonical extension for the ImageType
func (i ImageType) OutputExt() string {
	if ext, ok := imageTypeExtensionMap[i]; ok {
		return ext
	}
	return ""
}

// OperationMath represents VIPS_OPERATION_MATH type
type OperationMath int

// OperationMath enum
const (
	OperationMathSin   OperationMath = C.VIPS_OPERATION_MATH_SIN
	OperationMathCos   OperationMath = C.VIPS_OPERATION_MATH_COS
	OperationMathTan   OperationMath = C.VIPS_OPERATION_MATH_TAN
	OperationMathAsin  OperationMath = C.VIPS_OPERATION_MATH_ASIN
	OperationMathAcos  OperationMath = C.VIPS_OPERATION_MATH_ACOS
	OperationMathAtan  OperationMath = C.VIPS_OPERATION_MATH_ATAN
	OperationMathLog   OperationMath = C.VIPS_OPERATION_MATH_LOG
	OperationMathLog10 OperationMath = C.VIPS_OPERATION_MATH_LOG10
	OperationMathExp   OperationMath = C.VIPS_OPERATION_MATH_EXP
	OperationMathExp10 OperationMath = C.VIPS_OPERATION_MATH_EXP10
)

// OperationMath2 represents VIPS_OPERATION_MATH2 type
type OperationMath2 int

// OperationMath2 enum
const (
	OperationMath2Pow OperationMath2 = C.VIPS_OPERATION_MATH2_POW
	OperationMath2Wop OperationMath2 = C.VIPS_OPERATION_MATH2_WOP
)

// OperationRound represents VIPS_OPERATION_ROUND type
type OperationRound int

// OperationRound enum
const (
	OperationRoundRint  OperationRound = C.VIPS_OPERATION_ROUND_RINT
	OperationRoundCeil  OperationRound = C.VIPS_OPERATION_ROUND_CEIL
	OperationRoundFloor OperationRound = C.VIPS_OPERATION_ROUND_FLOOR
)

// OperationRelational represents VIPS_OPERATION_RELATIONAL type
type OperationRelational int

// OperationRelational enum
const (
	OperationRelationalEqual  OperationRelational = C.VIPS_OPERATION_RELATIONAL_EQUAL
	OperationRelationalNotEq  OperationRelational = C.VIPS_OPERATION_RELATIONAL_NOTEQ
	OperationRelationalLess   OperationRelational = C.VIPS_OPERATION_RELATIONAL_LESS
	OperationRelationalLessEq OperationRelational = C.VIPS_OPERATION_RELATIONAL_LESSEQ
	OperationRelationalMore   OperationRelational = C.VIPS_OPERATION_RELATIONAL_MORE
	OperationRelationalMoreEq OperationRelational = C.VIPS_OPERATION_RELATIONAL_MOREEQ
)

// OperationBoolean represents VIPS_OPERATION_BOOLEAN type
type OperationBoolean int

// OperationBoolean enum
const (
	OperationBooleanAnd    OperationBoolean = C.VIPS_OPERATION_BOOLEAN_AND
	OperationBooleanOr     OperationBoolean = C.VIPS_OPERATION_BOOLEAN_OR
	OperationBooleanEOr    OperationBoolean = C.VIPS_OPERATION_BOOLEAN_EOR
	OperationBooleanLShift OperationBoolean = C.VIPS_OPERATION_BOOLEAN_LSHIFT
	OperationBooleanRShift OperationBoolean = C.VIPS_OPERATION_BOOLEAN_RSHIFT
)

// OperationComplex represents VIPS_OPERATION_COMPLEX type
type OperationComplex int

// OperationComplex enum
const (
	OperationComplexPolar OperationComplex = C.VIPS_OPERATION_COMPLEX_POLAR
	OperationComplexRect  OperationComplex = C.VIPS_OPERATION_COMPLEX_RECT
	OperationComplexConj  OperationComplex = C.VIPS_OPERATION_COMPLEX_CONJ
)

// OperationComplex2 represents VIPS_OPERATION_COMPLEX2 type
type OperationComplex2 int

// OperationComplex2 enum
const (
	OperationComplex2CrossPhase OperationComplex2 = C.VIPS_OPERATION_COMPLEX2_CROSS_PHASE
)

// OperationComplexGet represents VIPS_OPERATION_COMPLEXGET type
type OperationComplexGet int

// OperationComplexGet enum
const (
	OperationComplexReal OperationComplexGet = C.VIPS_OPERATION_COMPLEXGET_REAL
	OperationComplexImag OperationComplexGet = C.VIPS_OPERATION_COMPLEXGET_IMAG
)

// Extend represents VIPS_EXTEND type
type Extend int

// Extend enum
const (
	ExtendBlack      Extend = C.VIPS_EXTEND_BLACK
	ExtendCopy       Extend = C.VIPS_EXTEND_COPY
	ExtendRepeat     Extend = C.VIPS_EXTEND_REPEAT
	ExtendMirror     Extend = C.VIPS_EXTEND_MIRROR
	ExtendWhite      Extend = C.VIPS_EXTEND_WHITE
	ExtendBackground Extend = C.VIPS_EXTEND_BACKGROUND
)

// Direction represents VIPS_DIRECTION type
type Direction int

// Direction enum
const (
	DirectionHorizontal Direction = C.VIPS_DIRECTION_HORIZONTAL
	DirectionVertical   Direction = C.VIPS_DIRECTION_VERTICAL
)

// Angle represents VIPS_ANGLE type
type Angle int

// Angle enum
const (
	Angle0   Angle = C.VIPS_ANGLE_D0
	Angle90  Angle = C.VIPS_ANGLE_D90
	Angle180 Angle = C.VIPS_ANGLE_D180
	Angle270 Angle = C.VIPS_ANGLE_D270
)

// Angle45 represents VIPS_ANGLE45 type
type Angle45 int

// Angle45 enum
const (
	Angle45_0   Angle45 = C.VIPS_ANGLE45_D0
	Angle45_45  Angle45 = C.VIPS_ANGLE45_D45
	Angle45_90  Angle45 = C.VIPS_ANGLE45_D90
	Angle45_135 Angle45 = C.VIPS_ANGLE45_D135
	Angle45_180 Angle45 = C.VIPS_ANGLE45_D180
	Angle45_225 Angle45 = C.VIPS_ANGLE45_D225
	Angle45_270 Angle45 = C.VIPS_ANGLE45_D270
	Angle45_315 Angle45 = C.VIPS_ANGLE45_D315
)

// Interpretation represents VIPS_INTERPRETATION type
type Interpretation int

// Interpretation enum
const (
	InterpretationError     Interpretation = C.VIPS_INTERPRETATION_ERROR
	InterpretationMultiband Interpretation = C.VIPS_INTERPRETATION_MULTIBAND
	InterpretationBw        Interpretation = C.VIPS_INTERPRETATION_B_W
	InterpretationHistogram Interpretation = C.VIPS_INTERPRETATION_HISTOGRAM
	InterpretationXyz       Interpretation = C.VIPS_INTERPRETATION_XYZ
	InterpretationLab       Interpretation = C.VIPS_INTERPRETATION_LAB
	InterpretationCMYK      Interpretation = C.VIPS_INTERPRETATION_CMYK
	InterpretationLabq      Interpretation = C.VIPS_INTERPRETATION_LABQ
	InterpretationRgb       Interpretation = C.VIPS_INTERPRETATION_RGB
	InterpretationCmd       Interpretation = C.VIPS_INTERPRETATION_CMC
	InterpretationLch       Interpretation = C.VIPS_INTERPRETATION_LCH
	InterpretationLabs      Interpretation = C.VIPS_INTERPRETATION_LABS
	InterpretationSrgb      Interpretation = C.VIPS_INTERPRETATION_sRGB
	InterpretationYxy       Interpretation = C.VIPS_INTERPRETATION_YXY
	InterpretationFourier   Interpretation = C.VIPS_INTERPRETATION_FOURIER
	InterpretationRgb16     Interpretation = C.VIPS_INTERPRETATION_RGB16
	InterpretationGrey16    Interpretation = C.VIPS_INTERPRETATION_GREY16
	InterpretationMatrix    Interpretation = C.VIPS_INTERPRETATION_MATRIX
	InterpretationScRgb     Interpretation = C.VIPS_INTERPRETATION_scRGB
	InterpretationHsv       Interpretation = C.VIPS_INTERPRETATION_HSV
)

// BandFormat represents VIPS_FORMAT type
type BandFormat int

// BandFormat enum
const (
	BandFormatNotSet    BandFormat = C.VIPS_FORMAT_NOTSET
	BandFormatUchar     BandFormat = C.VIPS_FORMAT_UCHAR
	BandFormatChar      BandFormat = C.VIPS_FORMAT_CHAR
	BandFormatUshort    BandFormat = C.VIPS_FORMAT_USHORT
	BandFormatShort     BandFormat = C.VIPS_FORMAT_SHORT
	BandFormatUint      BandFormat = C.VIPS_FORMAT_UINT
	BandFormatInt       BandFormat = C.VIPS_FORMAT_INT
	BandFormatFloat     BandFormat = C.VIPS_FORMAT_FLOAT
	BandFormatComplex   BandFormat = C.VIPS_FORMAT_COMPLEX
	BandFormatDouble    BandFormat = C.VIPS_FORMAT_DOUBLE
	BandFormatDpComplex BandFormat = C.VIPS_FORMAT_DPCOMPLEX
)

// Coding represents VIPS_CODING type
type Coding int

// Coding enum
const (
	CodingError Coding = C.VIPS_CODING_ERROR
	CodingNone  Coding = C.VIPS_CODING_NONE
	CodingLabq  Coding = C.VIPS_CODING_LABQ
	CodingRad   Coding = C.VIPS_CODING_RAD
)

// Access represents VIPS_ACCESS
type Access int

// Access enum
const (
	AccessRandom               Access = C.VIPS_ACCESS_RANDOM
	AccessSequential           Access = C.VIPS_ACCESS_SEQUENTIAL
	AccessSequentialUnbuffered Access = C.VIPS_ACCESS_SEQUENTIAL_UNBUFFERED
)

// OperationMorphology represents VIPS_OPERATION_MORPHOLOGY
type OperationMorphology int

// OperationMorphology enum
const (
	MorphologyErode  OperationMorphology = C.VIPS_OPERATION_MORPHOLOGY_ERODE
	MorphologyDilate OperationMorphology = C.VIPS_OPERATION_MORPHOLOGY_DILATE
)

var imageTypes = map[ImageType]string{
	ImageTypeGif:    "gif",
	ImageTypeJpeg:   "jpeg",
	ImageTypeMagick: "magick",
	ImageTypePdf:    "pdf",
	ImageTypePng:    "png",
	ImageTypeSvg:    "svg",
	ImageTypeTiff:   "tiff",
	ImageTypeWebp:   "webp",
}

var (
	once                sync.Once
	typeLoaders         = make(map[string]ImageType)
	supportedImageTypes = make(map[ImageType]bool)
)

// DetermineImageType attempts to determine the image type of the given buffer
func DetermineImageType(buf []byte) ImageType {
	startupIfNeeded()

	size := len(buf)
	if size == 0 {
		return ImageTypeUnknown
	}

	cName := C.vips_foreign_find_load_buffer(
		byteArrayPointer(buf),
		C.size_t(size))

	if cName == nil {
		return ImageTypeUnknown
	}

	imageType := ImageTypeUnknown
	name := strings.ToLower(C.GoString(cName))
	if imageType, ok := typeLoaders[name]; ok {
		return imageType
	}

	return imageType
}

// InitTypes initializes caches and figures out which image types are supported
func initTypes() {
	once.Do(func() {
		cType := C.CString("VipsOperation")
		defer freeCString(cType)

		for k, v := range imageTypes {
			name := strings.ToLower("VipsForeignLoad" + v)
			typeLoaders[name] = k
			typeLoaders[name+"buffer"] = k

			cFunc := C.CString(v + "load")
			defer freeCString(cFunc)

			ret := C.vips_type_find(
				cType,
				cFunc)
			supportedImageTypes[k] = int(ret) != 0
		}
	})
}
