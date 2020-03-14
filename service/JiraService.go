package service

import (
	"fmt"
	"jira-reminder/model"
	"jira-reminder/repository"
	"jira-reminder/tools"
	"math"
	"strings"
	"time"
)

type JiraService struct {
}

func (JiraService) Remind() {

	//获取要推送的jira单
	//获取问题类型
	issueTypeRepository := new(repository.IssueTypeRepository)
	var issueTypes = issueTypeRepository.GetByPNames(strings.Split(repository.Cfg.Section("jira").Key("type").String(), ","))
	var issueTypeIds []int32
	for _, issueType := range issueTypes {
		issueTypeIds = append(issueTypeIds, issueType.ID)
	}

	//获取问题状态
	issueStatusRepository := new(repository.IssueStatusRepository)
	var issueStatuses = issueStatusRepository.GetByPNames(strings.Split(repository.Cfg.Section("jira").Key("status").String(), ","))
	var issueStatusIds []int32
	for _, issueStatus := range issueStatuses {
		issueStatusIds = append(issueStatusIds, issueStatus.ID)
	}

	//获取项目
	projectRepository := new(repository.ProjectRepository)
	var projects = projectRepository.GetByPKeys(strings.Split(repository.Cfg.Section("jira").Key("project_code").String(), ","))
	var projectIds []int32
	for _, project := range projects {
		projectIds = append(projectIds, project.ID)
	}

	jiraIssueRepository := new(repository.JiraIssueRepository)
	jiraIssues := jiraIssueRepository.WithProjectInAndIssueTypeInAndIssueStatusIn(projectIds, issueTypeIds, issueStatusIds)

	//经办人分类
	var assigneeMap = make(map[string][]model.JiraIssue)
	for _, jiraIssue := range jiraIssues {
		if _, ok := assigneeMap[jiraIssue.Assignee]; ok {
			assigneeMap[jiraIssue.Assignee] = append(assigneeMap[jiraIssue.Assignee], jiraIssue)
		} else {
			assigneeMap[jiraIssue.Assignee] = []model.JiraIssue{jiraIssue}
		}
	}

	//整合推送内容，按经办人推送
	for assignee, jiraIssues := range assigneeMap {
		var msg string
		var isAllContinue = 1

		for _, jiraIssue := range jiraIssues {
			var suffix string
			duration := jiraIssue.DueDate.Sub(time.Now())
			diffDays := math.Floor(duration.Hours() / 24)
			if diffDays == 2 {
				suffix = "任务后天提测，有问题请及时暴露风险"
			} else if diffDays == 1 {
				suffix = "任务明天提测，有问题请及时暴露风险"
			} else if diffDays == 0 {
				suffix = "任务今天提测，有问题请及时暴露风险"
			} else if diffDays < 0 {
				suffix = fmt.Sprintf("任务已过期%v天，迭代存在延期风险", math.Abs(diffDays))
			} else {
				continue
			}
			isAllContinue = 0
			msg += fmt.Sprintf("“%s”%s\n", jiraIssue.Summary, suffix)
		}
		if len(jiraIssues) == 0 {
			continue
		}
		if isAllContinue == 1 {
			continue
		}
		var pusher = new(tools.WechatPusher)
		pusher.Push(msg, []string{assignee})
	}
}
