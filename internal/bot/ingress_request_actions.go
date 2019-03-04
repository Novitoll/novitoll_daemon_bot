// SPDX-License-Identifier: GPL-2.0
package bot

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (req *BotInReq) Process(app *App) {
	app.Logger.WithFields(logrus.Fields{
		"userId":    req.Message.From.Id,
		"username":  req.Message.From.Username,
		"chat":      req.Message.Chat.Username,
		"messageId": req.Message.MessageId,
	}).Info("Process: Running.")

	job := &Job{req, app}

	results, errors := FanOutProcessJobs(job, []ProcessJobFn{
		JobNewChatMemberDetector,
		JobNewChatMemberAuth,
		JobUrlDuplicationDetector,
		JobMsgStats,
		JobAdDetector,
		JobSentimentDetector,
		JobLeftParticipantDetector,
	})

	for _, e := range errors {
		app.Logger.Fatal(fmt.Sprintf("%v", e))
	}

	app.Logger.WithFields(logrus.Fields{
		"completed": len(results),
		"errors":    len(errors),
	}).Info("Process: Completed")
}