package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/camtrik/go-tools/internal/timer"
	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time format process",
	Long:  "Time format process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "Get current time",
	Long:  "Get current time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("Output: %s", nowTime.Format("2006-01-02 15:04:05"))
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "Get calculate time",
	Long:  "Get calculate time after duration",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			// if cannot parse time, try to parse as unix time
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatal("timer.GetCalculateTime err: ", err)
		}
		log.Printf("Output: %s", t.Format(layout))
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `Calculate time, e.g. "2022-01-01 12:00:00" or "2022-01-01"`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `Duration time, e.g. "1h" "1m" "1s"`)
}
