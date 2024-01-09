package dao

import (
	"daily_plan_go/common/logger"
	"daily_plan_go/models"
	"encoding/json"
	"fmt"
	"time"
)

type Task struct{}

// TaskDetail detail结构
type TaskDetail struct {
	Label string
	Done  bool
}

// create 增加一条任务记录
func (d Task) create(date time.Time, detail []TaskDetail, userid int) error {
	detailJson, err := json.Marshal(detail)
	if err != nil {
		return err
	}

	task := models.Task{
		Date:   date,
		Detail: string(detailJson),
		Userid: userid,
	}

	result := Db.Create(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// readN 查询某一天开始之后，一共n天的任务
func (d Task) readN(date time.Time, n, userid int) error {
	if n < 1 {
		return fmt.Errorf("readN 参数错误， n：%v", n)
	}

	// 获取指定日期以及之后一共n-1天的日期数组
	lastDate := date.AddDate(0, 0, n-1)

	var taskList []models.Task
	err := Db.Model(&models.Task{}).
		Where("date BETWEEN ? And ?", date, lastDate).
		Where("userid = ?", userid).
		Order("date").
		Find(&taskList).
		Error
	if err != nil {
		logger.Log.Errorf("readN 查询记录出错 userid:%v date:%v n:%v err:%v", userid, date, n, err.Error())
		return err
	}

	return nil
}

// update 更新一条任务记录
func (d Task) update(date time.Time, detail []TaskDetail, userid int) error {
	task := &models.Task{}

	err := Db.Model(task).
		Where("userid = ? AND date = ?", userid, date).
		First(task).
		Error
	if err != nil {
		logger.Log.Errorf("update 查询记录出错 userid:%v date:%v err:%v", userid, date, err.Error())
		return err
	}

	detailJson, err := json.Marshal(detail)
	if err != nil {
		logger.Log.Errorf("update 序列化任务出错 userid:%v date:%v err:%v", userid, date, err.Error())
		return err
	}

	task.Detail = string(detailJson)
	err = Db.Update(task).Error
	if err != nil {
		logger.Log.Errorf("update 更新数据库出错 userid:%v date:%v err:%v", userid, date, err.Error())
		return err
	}

	return nil
}
