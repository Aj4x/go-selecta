package selecta

func Select[I ~[]IT, O ~[]OT, IT, OT any](input I, action func(IT) (OT, error)) (O, error) {
	output := make(O, len(input))
	for index, item := range input {
		o, err := action(item)
		if err != nil {
			return nil, err
		}
		output[index] = o
	}
	return output, nil
}

func Where[I ~[]IT, IT any](input I, action func(IT) (bool, error)) (I, error) {
	var output I
	for _, item := range input {
		match, err := action(item)
		if err != nil {
			return nil, err
		}
		if match {
			output = append(output, item)
		}
	}
	return output, nil
}

func SelectWhere[I ~[]IT, O ~[]OT, IT, OT any](input I, action func(IT) (bool, OT, error)) (O, error) {
	var output O
	for _, item := range input {
		match, out, err := action(item)
		if err != nil {
			return nil, err
		}
		if match {
			output = append(output, out)
		}
	}
	return output, nil
}

func Any[I ~[]IT, IT any](input I, action func(IT) (bool, error)) (bool, error) {
	for _, item := range input {
		match, err := action(item)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}

func All[I ~[]IT, IT any](input I, action func(IT) (bool, error)) (bool, error) {
	for _, item := range input {
		match, err := action(item)
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
	}
	return true, nil
}

func IndexOf[I ~[]IT, IT any](input I, action func(IT) (bool, error)) (int, error) {
	for i, item := range input {
		match, err := action(item)
		if err != nil {
			return -1, err
		}
		if match {
			return i, nil
		}
	}
	return -1, nil
}

func ForEach[I ~[]IT, IT any](input I, action func(IT) error) error {
	for _, item := range input {
		if err := action(item); err != nil {
			return err
		}
	}
	return nil
}

func GroupToMap[I ~[]IT, IT any, K comparable](input I, getKey func(IT) K) (map[K][]IT, error) {
	output := make(map[K][]IT)
	for _, item := range input {
		key := getKey(item)
		_, ok := output[key]
		if !ok {
			keyData, err := Where(input, func(it IT) (bool, error) { return getKey(it) == key, nil })
			if err != nil {
				return nil, err
			}
			output[key] = keyData
		}
	}
	return output, nil
}

func GroupBy[I ~[]IT, IT any, K comparable](input I, getKey func(IT) K) ([][]IT, error) {
	m, err := GroupToMap(input, getKey)
	if err != nil {
		return nil, err
	}
	output := make([][]IT, len(m))
	i := 0
	for _, v := range m {
		output[i] = v
		i++
	}
	return output, nil
}

func MapToSlice[I ~map[K]V, K comparable, V, N any](input I, action func(K, V) (N, error)) ([]N, error) {
	output := make([]N, len(input))
	i := 0
	for k, v := range input {
		n, err := action(k, v)
		if err != nil {
			return nil, err
		}
		output[i] = n
		i++
	}
	return output, nil
}
