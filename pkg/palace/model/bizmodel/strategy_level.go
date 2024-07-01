package bizmodel

import (
	"encoding/json"

	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/vobj"

	"google.golang.org/protobuf/types/known/durationpb"
)

const TableNameStrategyLevel = "strategy_level"

type StrategyLevel struct {
	model.AllFieldModel
	// 所属策略
	StrategyID uint32    `gorm:"column:strategy_id;type:int unsigned;not null;comment:策略ID" json:"strategy_id"`
	Strategy   *Strategy `gorm:"foreignKey:StrategyID" json:"strategy"`

	// 持续时间
	Duration *durationpb.Duration `gorm:"column:duration;type:varchar(64);not null;comment:告警持续时间" json:"duration"`
	// 持续次数
	Count uint32 `gorm:"column:count;type:int unsigned;not null;comment:持续次数" json:"count"`
	// 持续事件类型
	SustainType vobj.Sustain `gorm:"column:sustain_type;type:int(11);not null;comment:持续类型" json:"sustain_type"`
	// 执行频率
	Interval *durationpb.Duration `gorm:"column:interval;type:varchar(64);not null;comment:执行频率" json:"interval"`
	// 条件
	Condition string `gorm:"column:condition;type:varchar(2);not null;comment:条件" json:"condition"`
	// 阈值
	Threshold float64 `gorm:"column:threshold;type:text;not null;comment:阈值" json:"threshold"`
	// 告警等级
	LevelID uint32   `gorm:"column:level_id;type:int unsigned;not null;comment:告警等级" json:"level_id"`
	Level   *SysDict `gorm:"foreignKey:LevelID" json:"level"`

	// 状态
	Status vobj.Status `gorm:"column:status;type:int;not null;comment:策略状态" json:"status"`
}

// String json string
func (c *StrategyLevel) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func (c *StrategyLevel) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *StrategyLevel) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// TableName StrategyLevel's table name
func (*StrategyLevel) TableName() string {
	return TableNameStrategyLevel
}