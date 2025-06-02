package transport

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type response struct {
	Result   json.RawMessage `json:"result"`
	Timeinfo timeInfo        `json:"time"`
	ErrInfo  *errInfo        `json:"-"`
}

type timeInfo struct {
	Start            float64   `json:"start"`
	Finish           float64   `json:"finish"`
	Duration         float64   `json:"duration"`
	Processing       float64   `json:"processing"`
	DateStart        time.Time `json:"date_start"`
	DateFinish       time.Time `json:"date_finish"`
	OperatingResetAt int64     `json:"operating_reset_at"`
	Operating        float64   `json:"operating"`
}

type errInfo struct {
	Err     string `json:"error"`
	ErrDesc string `json:"error_description"`
}

func (r *response) UnmarshalJSON(data []byte) error {
	const op = "response.UnmarshalJSON"

	var raw struct {
		Error            any             `json:"error"`
		ErrorDescription string          `json:"error_description"`
		Time             timeInfo        `json:"time"`
		Result           json.RawMessage `json:"result"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("%s: base unmarshal failed: %w", op, err)
	}

	if raw.Error != nil {
		r.ErrInfo = &errInfo{
			ErrDesc: raw.ErrorDescription,
		}

		switch e := raw.Error.(type) {
		case string:
			r.ErrInfo.Err = e

		case float64:
			r.ErrInfo.Err = strconv.Itoa(int(e))

		default:
			r.ErrInfo.Err = fmt.Sprintf("%v", e)
		}

		return nil
	}

	r.Result = raw.Result
	r.Timeinfo = raw.Time

	return nil
}
