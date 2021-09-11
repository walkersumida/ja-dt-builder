package jadtbuilder

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func timeParse(year int, month time.Month, day int, hour int) time.Time {
	layout := "2006-1-2 15:04 MST"
	datetime, _ := time.Parse(layout,
		fmt.Sprintf("%d-%d-%d %d:00 JST", year, month, day, hour))

	return datetime
}

func outPutFmtDt(month time.Month, day int, weekday string, hour int) string {
	return fmt.Sprintf("%d月%d日(%s) %d:00", int(month), day, weekday, hour)
}

func Build(dt string) string {
	now := time.Now()
	wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)

	year := nowJST.Year()
	month := nowJST.Month()
	day := nowJST.Day()
	hour := nowJST.Hour()
	weekday := wdays[nowJST.Weekday()]

	splitedDt := strings.Fields(dt)

	if len(splitedDt) >= 1 {
		hour, _ = strconv.Atoi(splitedDt[len(splitedDt)-1])
		datetime := timeParse(year, month, day, hour)

		if nowJST.Unix() > datetime.Unix() {
			day += 1
			datetime = timeParse(year, month, day, hour)
			weekday = wdays[datetime.Weekday()]
		}
	}

	if len(splitedDt) >= 2 {
		day, _ = strconv.Atoi(splitedDt[len(splitedDt)-2])
		datetime := timeParse(year, month, day, hour)
		weekday = wdays[datetime.Weekday()]

		if nowJST.Unix() > datetime.Unix() {
			monthNum := int(month) + 1
			month = time.Month(monthNum)
			datetime = timeParse(year, month, day, hour)
			weekday = wdays[datetime.Weekday()]
		}
	}

	if len(splitedDt) >= 3 {
		monthNum, _ := strconv.Atoi(splitedDt[len(splitedDt)-3])
		month = time.Month(monthNum)
		datetime := timeParse(year, month, day, hour)
		weekday = wdays[datetime.Weekday()]

		if nowJST.Unix() > datetime.Unix() {
			year += 1
			datetime = timeParse(year, month, day, hour)
			weekday = wdays[datetime.Weekday()]
		}
	}

	if len(splitedDt) >= 4 {
		year, _ = strconv.Atoi(splitedDt[len(splitedDt)-4])
		datetime := timeParse(year, month, day, hour)
		weekday = wdays[datetime.Weekday()]
	}

	return outPutFmtDt(month, day, weekday, hour)
}
