package models

import (
	"testing"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
)

func Test_Fees(t *testing.T) {
	u := &User{
		LastMonthVolume: 0,
		FeeDiscountRate: decimaldt.NewFromFloat(0),
	}
	for _, entry := range takerMakerFees {
		u.LastMonthVolume = entry.MinVolume - 1
		if f := getMakerFeeRate(u); !f.Equal(entry.Maker) {
			t.Errorf(
				"Volume: %d, Discount: %s ->  expected: %s, got: %s",
				u.LastMonthVolume,
				u.FeeDiscountRate,
				entry.Maker.String(),
				f.String(),
			)
		}
		if f := getTakerFeeRate(u); !f.Equal(entry.Taker) {
			t.Errorf(
				"Volume: %d, Discount: %s ->  expected: %s, got: %s",
				u.LastMonthVolume,
				u.FeeDiscountRate,
				entry.Maker.String(),
				f.String(),
			)
		}
	}
}

func Test_DiscountedFees(t *testing.T) {
	u := &User{
		LastMonthVolume: 0,
		FeeDiscountRate: decimaldt.NewFromFloat(0.2),
	}
	for _, entry := range takerMakerFees {
		{
			u.LastMonthVolume = entry.MinVolume - 1
			expected := entry.Maker.Sub(
				entry.Maker.Mul(u.FeeDiscountRate),
			)
			if f := getMakerFeeRate(u); !f.Equal(expected) {
				t.Errorf(
					"Volume: %d, Discount: %s ->  expected: %s, got: %s",
					u.LastMonthVolume,
					u.FeeDiscountRate,
					entry.Maker.String(),
					f.String(),
				)
			}
		}
		{
			u.LastMonthVolume = entry.MinVolume - 1
			expected := entry.Taker.Sub(
				entry.Taker.Mul(u.FeeDiscountRate),
			)
			if f := getTakerFeeRate(u); !f.Equal(expected) {
				t.Errorf(
					"Volume: %d, Discount: %s ->  expected: %s, got: %s",
					u.LastMonthVolume,
					u.FeeDiscountRate,
					entry.Maker.String(),
					f.String(),
				)
			}
		}
	}
}
