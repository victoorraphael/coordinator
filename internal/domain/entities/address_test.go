package entities_test

import (
	"testing"
)

//nolint:funnel
func TestAddress(t *testing.T) {
	//cases := []struct {
	//	Name        string
	//	Addr        entities.Address
	//	ErrExpected error
	//	ErrMsgField string
	//}{
	//	{
	//		Name: "correct address",
	//		Addr: entities.Address{
	//			ID:     1,
	//			UUID:   uid.NewUUID().String(),
	//			Street: "teste",
	//			City:   "garnahuns",
	//			Zip:    "192389182",
	//			Number: 12,
	//		},
	//		ErrExpected: nil,
	//		ErrMsgField: "",
	//	}, {
	//		Name: "address missing street",
	//		Addr: entities.Address{
	//			ID:     1,
	//			UUID:   uid.NewUUID().String(),
	//			Street: "",
	//			City:   "garnahuns",
	//			Zip:    "192389182",
	//			Number: 12,
	//		},
	//		ErrExpected: errs.ErrFieldViolation,
	//		ErrMsgField: "street",
	//	}, {
	//		Name: "address missing city",
	//		Addr: entities.Address{
	//			ID:     1,
	//			UUID:   uid.NewUUID().String(),
	//			Street: "teste",
	//			City:   "",
	//			Zip:    "192389182",
	//			Number: 12,
	//		},
	//		ErrExpected: errs.ErrFieldViolation,
	//		ErrMsgField: "city",
	//	}, {
	//		Name: "address missing zip",
	//		Addr: entities.Address{
	//			ID:     1,
	//			UUID:   uid.NewUUID().String(),
	//			Street: "teste",
	//			City:   "garnahuns",
	//			Zip:    "",
	//			Number: 12,
	//		},
	//		ErrExpected: errs.ErrFieldViolation,
	//		ErrMsgField: "zip",
	//	}, {
	//		Name: "address missing number",
	//		Addr: entities.Address{
	//			ID:     1,
	//			UUID:   uid.NewUUID().String(),
	//			Street: "",
	//			City:   "garnahuns",
	//			Zip:    "192389182",
	//			Number: 0,
	//		},
	//		ErrExpected: errs.ErrFieldViolation,
	//		ErrMsgField: "number",
	//	}, {
	//		Name:        "address missing everything",
	//		Addr:        entities.Address{},
	//		ErrExpected: errs.ErrFieldViolation,
	//		ErrMsgField: "",
	//	},
	//}

	//for _, tc := range cases {
	//	t.Run(tc.Name, func(t *testing.T) {
	//		err := tc.Addr.Validate()
	//		assert.True(t, errors.Is(err, tc.ErrExpected))
	//		if tc.ErrExpected != nil {
	//			assert.True(t, strings.Contains(err.Error(), tc.ErrMsgField))
	//		}
	//	})
	//}
}
