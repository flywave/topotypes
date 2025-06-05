package topo4d

type WorkStatus string

const (
	WorkStatusNotStarted WorkStatus = "NOT_STAETED"
	WorkStatusInProgress WorkStatus = "IN_PROGRESS"
	WorkStatusCompleted  WorkStatus = "COMPLETED"
)

type WorkTask struct {
	ID            string                 `json:"id"`
	PreviousID    string                 `json:"previousId"`
	Name          string                 `json:"name"`
	Status        WorkStatus             `json:"status"`
	Description   string                 `json:"description"`
	ScheduleStart string                 `json:"scheduleStart"`
	ScheduleEnd   string                 `json:"scheduleEnd"`
	ActualStart   string                 `json:"actualStart"`
	ActualEnd     string                 `json:"actualEnd"`
	Progress      float64                `json:"progress"`
	Metadata      map[string]interface{} `json:"metadata"`
}

type ProgressType string

const (
	ProgressByRatio    ProgressType = "RATIO"    // 按比例计算进度
	ProgressByDistance ProgressType = "DISTANCE" // 按距离计算进度
)

type WorkSchedule struct {
	ID          string       `json:"id"`
	Layer       string       `json:"layer,omitempty"`
	PlanID      string       `json:"planId"`
	Type        ProgressType `json:"type"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Creator     string       `json:"creator"`
	StartTime   string       `json:"startTime"`
	EndTime     string       `json:"endTime"`
	Tasks       []*WorkTask  `json:"tasks"`
	Total       float64      `json:"total"`
	StartValue  *float64     `json:"startValue,omitempty"`
	EndValue    *float64     `json:"endValue,omitempty"`
}

type WorkPlan struct {
	ID          string                 `json:"id"`
	Layer       string                 `json:"layer,omitempty"`
	Type        ProgressType           `json:"type"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	StartTime   string                 `json:"startTime"`
	EndTime     string                 `json:"endTime"`
	Total       float64                `json:"total"`
	StartValue  *float64               `json:"startValue,omitempty"`
	EndValue    *float64               `json:"endValue,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type TopoWork struct {
	Schedules []*WorkSchedule        `json:"schedules"`
	Plans     []*WorkPlan            `json:"plans"`
	Metadata  map[string]interface{} `json:"metadata"`
}
