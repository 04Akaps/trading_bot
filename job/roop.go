package job

func (j *Job) volumeTraceDiffChecker() {
	for {
		select {
		case symbol := <-j.volumeUpdateChannel:
			go func() {
				total, currentVolume, diff := j.mongoDB.GetVolumeInfo(symbol)

				j.slackClient.VolumeTracker(symbol, total, currentVolume, diff)
			}()
		}
	}
}
