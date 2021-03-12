package loto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistory_OrderRecent(t *testing.T) {
	h := History{RecentDraws: []Draw{
		{ID: "2018153"},
		{ID: "2019024"},
		{ID: "20211028"},
		{ID: "20199132"},
		{ID: "2019131"},
		{ID: "20200153"},
	}}
	t.Run("Should be ok", func(t *testing.T) {
		expectedDraws := []Draw{
			{ID: "2018153"},
			{ID: "2019024"},
			{ID: "2019131"},
			{ID: "20199132"},
			{ID: "20200153"},
			{ID: "20211028"},
		}
		h.OrderRecent()

		assert.EqualValues(t, expectedDraws, h.RecentDraws)
	})
}

func TestHistory_OrderOld(t *testing.T) {
	h := History{OldDraws: []OldDraw{
		{ID: "2018153", Tirage: 1},
		{ID: "2018153", Tirage: 2},
		{ID: "2019024", Tirage: 2},
		{ID: "2019024", Tirage: 1},
		{ID: "20211028", Tirage: 2},
		{ID: "20211028", Tirage: 1},
		{ID: "20199132", Tirage: 2},
		{ID: "20199132", Tirage: 1},
		{ID: "2019131", Tirage: 2},
		{ID: "2019131", Tirage: 1},
		{ID: "20200153", Tirage: 2},
		{ID: "20200153", Tirage: 1},
	}}
	t.Run("Should be ok", func(t *testing.T) {
		expectedDraws := []OldDraw{
			{ID: "2018153", Tirage: 1},
			{ID: "2018153", Tirage: 2},
			{ID: "2019024", Tirage: 1},
			{ID: "2019024", Tirage: 2},
			{ID: "2019131", Tirage: 1},
			{ID: "2019131", Tirage: 2},
			{ID: "20199132", Tirage: 1},
			{ID: "20199132", Tirage: 2},
			{ID: "20200153", Tirage: 1},
			{ID: "20200153", Tirage: 2},
			{ID: "20211028", Tirage: 1},
			{ID: "20211028", Tirage: 2},
		}
		h.OrderOld()

		assert.EqualValues(t, expectedDraws, h.OldDraws)
	})
}

func TestHistory_DrawByIndex(t *testing.T) {
	h := History{RecentDraws: []Draw{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
		{ID: "4"},
	}}
	t.Run("Should return an error because index < 0", func(t *testing.T) {
		expectedErr := "invalid index < 0"
		expectedDraw := Draw{}
		d, err := h.DrawByIndex(-1)

		if assert.EqualError(t, err, expectedErr) {
			assert.EqualValues(t, d, expectedDraw)
		}
	})
	t.Run("Should return an error because index is out of range", func(t *testing.T) {
		expectedErr := "invalid index to get the draw by index"
		expectedDraw := Draw{}
		d, err := h.DrawByIndex(4)

		if assert.EqualError(t, err, expectedErr) {
			assert.EqualValues(t, d, expectedDraw)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		expectedDraw := Draw{ID: "4"}
		d, err := h.DrawByIndex(3)

		if assert.NoError(t, err) {
			assert.EqualValues(t, d, expectedDraw)
		}
	})
}

func TestHistory_Len(t *testing.T) {
	h := History{}
	t.Run("Should be ok with an history not initialized", func(t *testing.T) {
		expectedLen := 0
		v := h.Len()

		assert.EqualValues(t, expectedLen, v)
	})

	t.Run("Should be ok with 2 draw added to the recent list", func(t *testing.T) {
		h.RecentDraws = []Draw{{ID: "first"}, {ID: "two"}}
		expectedLen := 2
		v := h.Len()

		assert.EqualValues(t, expectedLen, v)
	})
}
