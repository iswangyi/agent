package models

import (
	Mtils "agent/utils"
	"fmt"
)

type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}

func (this *MetricValue) String() string {
	return fmt.Sprintf("Endpoint:%s,Metric:%s,Value:%v,Step:%d,Type:%s,Tags:%s,Timestamp:%d",
		this.Endpoint,
		this.Metric,
		this.Value,
		this.Step,
		this.Type,
		this.Tags,
		this.Timestamp)
}

type MetaData struct {
	Metric      string            `json:"metric"`
	Endpoint    string            `json:"endpoint"`
	Timestamp   int64             `json:"timestamp"`
	Step        int64             `json:"step"`
	Value       float64           `json:"value"`
	CounterType string            `json:"counterType"`
	Tags        map[string]string `json:"tags"`
}

type JsonMetaData struct {
	Metric      string      `json:"metric"`
	Endpoint    string      `json:"endpoint"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
	Value       interface{} `json:"value"`
	CounterType string      `json:"counterType"`
	Tags        string      `json:"tags"`
}

func (t *JsonMetaData) String() string {
	return fmt.Sprintf("<JsonMetaData Endpoint:%s, Metric:%s, Tags:%s, DsType:%s, Step:%d, Value:%v, Timestamp:%d>",
		t.Endpoint, t.Metric, t.Tags, t.CounterType, t.Step, t.Value, t.Timestamp)
}

func (t *MetaData) String() string {
	return fmt.Sprintf("<MetaData Endpoint:%s, Metric:%s, Timestamp:%d, Step:%d, Value:%f, Tags:%v>",
		t.Endpoint, t.Metric, t.Timestamp, t.Step, t.Value, t.Tags)
}

func (t *MetaData) PK() string {
	return Mutils.PK(t.Endpoint, t.Metric, t.Tags)
}
