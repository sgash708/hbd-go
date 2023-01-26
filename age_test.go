package age

import (
	"testing"
	"time"
)

func TestCalAge(t *testing.T) {
	t.Parallel()

	type input struct {
		Target time.Time
		Birth  time.Time
	}
	cases := []struct {
		Name  string
		Input input
		Exp   int
	}{
		{
			Name: "Success_85",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(1936, 8, 22, 0, 0, 0, 0, time.Local),
			},
			Exp: 85,
		},
		{
			Name: "Success_0_today",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 0,
		},
		{
			Name: "Success_0_future",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2023, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 0,
		},
		{
			Name: "Success_1",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 1,
		},
		{
			Name: "Success_21_leapYear",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2000, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 21,
		},
		{
			Name: "Success_121_leapYear",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(1900, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 121,
		},
		{
			Name: "Success_98_leapYear",
			Input: input{
				Target: time.Date(2000, 12, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(1901, 12, 31, 0, 0, 0, 0, time.Local),
			},
			Exp: 98,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			res := CalAge(c.Input.Birth, c.Input.Target)
			if res != c.Exp {
				t.Errorf("想定する値と違います。want: %v real: %v", c.Exp, res)
			}
		})
	}
}

func TestCalAgeNowTime(t *testing.T) {
	t.Parallel()

	type input struct {
		Target time.Time
		Birth  time.Time
	}
	cases := []struct {
		Name  string
		Input input
		Exp   int
	}{
		{
			Name: "Success_85",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(1936, 8, 22, 0, 0, 0, 0, time.Local),
			},
			Exp: 85,
		},
		{
			Name: "Success_0",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 0,
		},
		{
			Name: "Success_0_future",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2023, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 0,
		},
		{
			Name: "Success_1",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 1,
		},
		{
			Name: "Success_21_leapYear",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(2000, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 21,
		},
		{
			Name: "Success_121_leapYear",
			Input: input{
				Target: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(1900, 11, 1, 0, 0, 0, 0, time.Local),
			},
			Exp: 121,
		},
		{
			Name: "Success_98_leapYear",
			Input: input{
				Target: time.Date(2000, 12, 1, 0, 0, 0, 0, time.Local),
				Birth:  time.Date(1901, 12, 31, 0, 0, 0, 0, time.Local),
			},
			Exp: 98,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			fakeTime = c.Input.Target
			res := CalAgeNowTime(c.Input.Birth)
			if res != c.Exp {
				t.Errorf("想定する値と違います。want: %v real: %v", c.Exp, res)
			}
		})
	}
}

func TestNow(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name string
		Exp  time.Time
	}{
		{
			Name: "Success_nowtime",
			Exp:  time.Time{},
		},
		{
			Name: "Success_faketime",
			Exp:  time.Date(1901, 12, 31, 0, 0, 0, 0, time.Local),
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			if c.Exp.IsZero() {
				fakeTime = time.Time{}
				minusTenMin := time.Now().Add(-10 * time.Minute)
				if now().Before(minusTenMin) {
					t.Errorf("現在時刻が正しく取得できていません")
				}
			} else {
				fakeTime = c.Exp
				if now() != c.Exp {
					t.Errorf("現在時刻が正しく取得できていません")
				}
			}
		})
	}
}
