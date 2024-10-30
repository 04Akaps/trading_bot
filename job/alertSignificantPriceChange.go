package job

import (
	"context"
	"fmt"
)

func (j *Job) alertSignificantPriceChange(c context.Context, cancel context.CancelFunc) {
	//2분마다 오전 09시에 갱신된 일별 가격 (dailyPrice) 에 비해서 2% 급상승하는 자산에 대한 알림
	fmt.Println("2분?")
}
