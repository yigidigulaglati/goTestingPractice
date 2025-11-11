package mathutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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

func TestAverage_TableDriven(t *testing.T){


	tests := []struct{
		name  string
		input []float64
		want float64
		wantErr bool
		errMsg string
	}{
		{
			name: `basic average`,
			input: []float64{1,2,3},
			want: 2.0,
			wantErr: false,
			errMsg: ``,
		},
		{
			name: `empty slice`,
			input: []float64{},
			want: 0.0,
			wantErr: true,
			errMsg: "cannot average empty slice",
		},{
			name: `negative numbers`,
			input: []float64{-1,-2,-3},
			want: -2.0,
			wantErr: false,
			errMsg: ``,
		},
	}

	for _, tt := range tests{

		t.Run(tt.name, func(t *testing.T){
			got, err := Average(tt.input);

			if tt.wantErr{
				require.Error(t, err);
				require.EqualError(t, err, tt.errMsg);
				require.Equal(t, tt.want, got)
			}else{
				require.NoError(t, err);
				require.Equal(t, tt.want, got);
			}
		})
	}
}


type AverageTestSuite struct{
	suite.Suite
}

func (s *AverageTestSuite) TestBasic(){

	got, err := Average([]float64{1.0, 2.0, 3.0});

	require.NoError(s.T(), err);
	require.Equal(s.T(), 2.0, got);
}

func (s *AverageTestSuite) TestEmpty(){
	got, err := Average([]float64{});

	require.Error(s.T(), err);
	require.Equal(s.T(), 0.0, got);
	require.EqualError(s.T(), err, "cannot average empty slice");
}

func (s *AverageTestSuite) TestNegative(){
	got, err := Average([]float64{-1.0, -2.0, -3.0});

	require.NoError(s.T(), err);
	require.Equal(s.T(), -2.0, got);
}

func TestAverageTestSuite(t *testing.T){
	suite.Run(t, new(AverageTestSuite));
}


