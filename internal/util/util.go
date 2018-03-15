package util



func CastInt(intsPtr []*int) (ints []int) {
	for _, v := range intsPtr {
		if v != nil {
			ints = append(ints, *v)
		}
	}
	return ints
}

func CastInt8(int8Ptrs []*int8) (int8s []int8){
	for _, v := range int8Ptrs {
		if v != nil {
			int8s = append(int8s, *v)
		}
	}
	return int8s
}

func CastInt32(int32Ptrs []*int32) (int32s []int32) {
	for _, v := range int32Ptrs {
		if v != nil {
			int32s = append(int32s, *v)
		}
	}
	return int32s
}


func CastInt64(int64Ptrs []*int64) (int64s []int64) {
	for _, v := range int64Ptrs {
		if v != nil {
			int64s = append(int64s, *v)
		}
	}
	return int64s
}

func CastUint(uintPtrs []*uint) (uints []uint) {
	for _, v := range uintPtrs {
		if v != nil {
			uints = append(uints, *v)
		}
	}
	return uints
}

func CastUint8(uint8Ptrs []*uint8) (uint8s []uint8) {
	for _, v := range uint8Ptrs {
		if v != nil {
			uint8s = append(uint8s, *v)
		}
	}
	return uint8s
}

func CastUint32(uint32Ptrs []*uint32) (uint32s []uint32) {
	for _, v := range uint32Ptrs {
		if v != nil {
			uint32s = append(uint32s, *v)
		}
	}
	return uint32s
}

func CastUint64(uint64Ptrs []*uint64) (uint64s []uint64) {
	for _, v := range uint64Ptrs {
		if v != nil {
			uint64s = append(uint64s, *v)
		}
	}
	return uint64s
}


func CastFloat32(float32Ptrs []*float32) (float32s []float32) {
	for _, v := range float32Ptrs {
		if v != nil {
			float32s = append(float32s, *v)
		}
	}
	return float32s
}

func CastFloat64(float64Ptrs []*float64) (float64s []float64) {
	for _, v := range float64Ptrs {
		if v != nil {
			float64s = append(float64s, *v)
		}
	}
	return float64s
}
