package logger

import (
	"os"

	"github.com/johntdyer/slackrus"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/sirupsen/logrus"
)

//Setup sets up logger
func Setup(cfg datatype.Config) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
	if cfg.SlackWebhookURL != "" {
		logrus.AddHook(&slackrus.SlackrusHook{
			HookURL:        cfg.SlackWebhookURL,
			AcceptedLevels: slackrus.LevelThreshold(logrus.WarnLevel),
			Channel:        "#log",
			IconEmoji:      ":japanese_ogre:",
			Username:       "logrus",
		})
	} else {
		logrus.Warn("Running without SLACK_WEBHOOK_URL")
	}
}
