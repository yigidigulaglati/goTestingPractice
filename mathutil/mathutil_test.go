package mathutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage_Basic(t *testing.T){
	got, err := Average([]float64{1.0, 2.0, 3.0});


	require.NoError(t, err);
	require.Equal(t, 2.0, got);
}

func TestAverage_EmptySlice(t *testing.T){
	got, err := Average([]float64{});

	require.Error(t, err);
	require.Equal(t, 0.0, got);
	require.EqualError(t, err, "cannot average empty slice");
}

func TestAverage_NegativeNumbers(t *testing.T){
	got, err := Average([]float64{-1.0, -2.0, -3.0});

	require.NoError(t, err);
	require.Equal(t, -2.0, got);
}




