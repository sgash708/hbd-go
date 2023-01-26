package age

import "time"

var fakeTime time.Time

// CalAge is ...
func CalAge(birthdayTime, targetTime time.Time) int {
	return calAge(birthdayTime, targetTime)
}

// CalAgeNowTime is ...
func CalAgeNowTime(birthdayTime time.Time) int {
	targetTime := now()

	return calAge(birthdayTime, targetTime)
}

func now() time.Time {
	if !fakeTime.IsZero() {
		return fakeTime
	}
	return time.Now()
}

func calAge(bTime, tTime time.Time) int {
	age := tTime.Year() - bTime.Year()
	if age < 0 {
		age = 0
		return age
	}

	birthDay := getAdjustedBirthDay(bTime, tTime)
	if tTime.YearDay() < birthDay {
		age--
	}

	return age
}

func getAdjustedBirthDay(birthDate time.Time, target time.Time) int {
	birthDay := birthDate.YearDay()
	birthYear := birthDate.Year()
	targetDay := target.YearDay()
	targetYear := target.Year()

	if isLeap(birthYear) && !isLeap(targetYear) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(targetYear) && !isLeap(birthYear) && targetDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

func isLeap(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
