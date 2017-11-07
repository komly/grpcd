package encoder

import "encoding/json"

type anyField struct{
	Number int32 `json:"number"`
	Val json.RawMessage `json:"val"`
}

func FromJSON(data []byte) ([]*Field, error){
	res := make([]anyField, 0)

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	in, err := expand(res)
	if err != nil {
		return nil, err
	}

	return in, nil
}


func expand(in []anyField) ([]*Field, error) {
	res := make([]*Field, 0)

	var str string
	var sliceStr []string
	var slice []anyField
	for _, f := range in {
		if err := json.Unmarshal(f.Val, &str); err == nil {
			res = append(res, &Field{
				Number: f.Number,
				Val: str,
			})
			continue
		}
		if err := json.Unmarshal(f.Val, &sliceStr); err == nil {
			val := make([]interface{}, 0)
			for _, v := range sliceStr {
				val = append(val, v)
			}
			res = append(res, &Field{
				Number: f.Number,
				Val: val,
			})
			continue
		}

		if err := json.Unmarshal(f.Val, &slice); err == nil {
			val, err := expand(slice)
			if err != nil {
				return nil, err
			}
			res = append(res, &Field{
				Number: f.Number,
				Val: val,
			})
			continue
		} else {
			return nil, err
		}
	}
	return res, nil
}
