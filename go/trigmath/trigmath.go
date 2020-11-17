package trigmath

import "math"

const PI = 3.1415
const SQUARED_PI = PI * PI
const HALF_PI = PI / 2
const QUARTER_PI = HALF_PI / 2
const TWO_PI = 2 * PI
const THREE_PI_HALVES = TWO_PI - HALF_PI
const DEG_TO_RAD = PI / 180
const HALF_DEG_TO_RAD = PI / 360
const RAD_TO_DEG = 180 / PI
const SQRT_OF_TWO = 1.41421356237
const HALF_SQRT_OF_TWO = 0.70710678118

const sq2p1 = 2.414213562373095048802
const sq2m1 = 0.414213562373095048802
const p4 = 0.161536412982230228262e2
const p3 = 0.26842548195503973794141e3
const p2 = 0.11530293515404850115428136e4
const p1 = 0.178040631643319697105464587e4
const p0 = 0.89678597403663861959987488e3
const q4 = 0.5895697050844462222791e2
const q3 = 0.536265374031215315104235e3
const q2 = 0.16667838148816337184521798e4
const q1 = 0.207933497444540981287275926e4
const q0 = 0.89678597403663861962481162e3

const SIN_SIZE = 100000
const SIN_MASK = SIN_SIZE - 1
const SIN_CONVERSION_FACTOR = SIN_SIZE / TWO_PI
const COS_OFFSET = SIN_SIZE / 4

type TrigMath struct {
	SIN_TABLE [SIN_SIZE + 1]float64
}

func NewTrigMath() *TrigMath {
	var trigMath = TrigMath{}
	for i := 0; i <= SIN_SIZE; i++ {
		trigMath.SIN_TABLE[i] = math.Sin(float64(i) * TWO_PI / SIN_SIZE)
	}
	return &trigMath
}

func (t *TrigMath) Sin(angle float64) float64 {
	return t.sinRaw(t.floor(angle * SIN_CONVERSION_FACTOR))
}

func (t *TrigMath) Asin(value float64) float64 {
	var temp float64
	if value > 1 {
		return math.NaN()
	}
	if value < 0 {
		return -t.Asin(-value)
	}
	temp = math.Sqrt(1 - value*value)
	if value > 0.7 {
		return HALF_PI - t.msatan(temp/value)
	}
	return t.msatan(value / temp)
}

func (t *TrigMath) Cos(angle float64) float64 {
	return t.cosRaw(t.floor(angle * SIN_CONVERSION_FACTOR))
}

func (t *TrigMath) Tan(angle float64) float64 {
	var idx int
	idx = t.floor(angle * SIN_CONVERSION_FACTOR)
	return t.sinRaw(idx) / t.cosRaw(idx)
}

func (t *TrigMath) Csc(angle float64) float64 {
	return 1 / t.Sin(angle)
}

func (t *TrigMath) Sec(angle float64) float64 {
	return 1 / t.Cos(angle)
}

func (t *TrigMath) Cot(angle float64) float64 {
	var idx int
	idx = t.floor(angle * SIN_CONVERSION_FACTOR)
	return t.cosRaw(idx) / t.sinRaw(idx)
}

func (t *TrigMath) Acos(value float64) float64 {
	if value > 1 || value < -1 {
		return math.NaN()
	}
	return HALF_PI - t.Asin(value)
}

func (t *TrigMath) Atan(value float64) float64 {
	if value > 0 {
		return t.msatan(value)
	}
	return -t.msatan(-value)
}

func (t *TrigMath) Atan2(y float64, x float64) float64 {
	if y+x == y {
		if y >= 0 {
			return HALF_PI
		} else {
			return -HALF_PI
		}
	}
	y = t.Atan(y / x)
	if x < 0 {
		if y <= 0 {
			return y + PI
		} else {
			return y - PI
		}
	}
	return y
}

func (t *TrigMath) Acsc(value float64) float64 {
	if value == 0 {
		return math.NaN()
	}
	return t.Asin(1 / value)
}

func (t *TrigMath) Asec(value float64) float64 {
	if value == 0 {
		return math.NaN()
	}
	return t.Acos(1 / value)
}

func (t *TrigMath) Acot(value float64) float64 {
	if value == 0 {
		return math.NaN()
	}
	if value > 0 {
		return t.Atan(1 / value)
	}
	return t.Atan(1/value) + PI
}

// Private

func (t *TrigMath) sinRaw(idx int) float64 {
	return t.SIN_TABLE[(idx & SIN_MASK)]
}

func (t *TrigMath) cosRaw(idx int) float64 {
	return t.SIN_TABLE[(idx+COS_OFFSET)&SIN_MASK]
}

func (t *TrigMath) mxatan(arg float64) float64 {
	var argsq float64
	var value float64
	argsq = arg * arg
	value = (((p4*argsq+p3)*argsq+p2)*argsq+p1)*argsq + p0
	value /= ((((argsq+q4)*argsq+q3)*argsq+q2)*argsq+q1)*argsq + q0
	return value * arg
}

func (t *TrigMath) msatan(arg float64) float64 {
	if arg < sq2m1 {
		return t.mxatan(arg)
	}
	if arg > sq2p1 {
		return HALF_PI - t.mxatan(1/arg)
	}
	return HALF_PI/2 + t.mxatan((arg-1)/(arg+1))
}

func (t *TrigMath) floor(a float64) int {
	return int(a)
}
