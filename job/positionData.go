package job

import "context"

func (j *Job) fetchAndSendPositionData(c context.Context, cancel context.CancelFunc) {
	//3분마다 포지션이 존재하면 데이터를 보내주고 저장
}
