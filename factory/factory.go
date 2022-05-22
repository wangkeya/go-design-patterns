package factory

type JobExecutor interface {
	Init()
	Start() string
	Stop() string
}

type LogJob struct {
}

func (l *LogJob) Init() {

}

func (l *LogJob) Start() string {
	return "logStart"
}

func (l *LogJob) Stop() string {
	return "logStop"
}

func Factory(jobType string) JobExecutor {
	switch jobType {
	case "log":
		return &LogJob{}
	}
	return nil
}
