package usecases

import (
	"github.com/A-ndrey/leetcode-client/internal/client"
	"github.com/A-ndrey/leetcode-client/internal/types"
)

type ProblemsSolvedStat struct {
	AllScore int
	AllTotal int

	EasyScore int
	EasyTotal int

	MediumScore int
	MediumTotal int

	HardScore int
	HardTotal int
}

func UserProblemsSolved(username string) (ProblemsSolvedStat, error) {
	op := types.Operation{
		Name: "userProblemsSolved",
		Fields: []types.Field{
			{
				Name: "allQuestionsCount",
				Fields: []types.Field{
					{
						Name: "difficulty",
					},
					{
						Name: "count",
					},
				},
			},
			{
				Name: "matchedUser",
				Args: map[string]any{"username": username},
				Fields: []types.Field{
					{
						Name: "submitStatsGlobal",
						Fields: []types.Field{
							{
								Name: "acSubmissionNum",
								Fields: []types.Field{
									{
										Name: "difficulty",
									},
									{
										Name: "count",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	type DifCount struct {
		Difficulty string `json:"difficulty"`
		Count      int    `json:"count"`
	}

	type RespData struct {
		Totals []DifCount `json:"allQuestionsCount"`
		User   struct {
			Stats struct {
				Scores []DifCount `json:"acSubmissionNum"`
			} `json:"submitStatsGlobal"`
		} `json:"matchedUser"`
	}

	req := client.NewRequest(op)

	respData := RespData{}

	resp, err := client.Do(req, &respData)
	if err != nil {
		return ProblemsSolvedStat{}, err
	}

	data := resp.Data.(*RespData)

	var stat ProblemsSolvedStat
	for _, t := range data.Totals {
		switch t.Difficulty {
		case "All":
			stat.AllTotal = t.Count
		case "Easy":
			stat.EasyTotal = t.Count
		case "Medium":
			stat.MediumTotal = t.Count
		case "Hard":
			stat.HardTotal = t.Count
		}
	}

	for _, s := range data.User.Stats.Scores {
		switch s.Difficulty {
		case "All":
			stat.AllScore = s.Count
		case "Easy":
			stat.EasyScore = s.Count
		case "Medium":
			stat.MediumScore = s.Count
		case "Hard":
			stat.HardScore = s.Count
		}
	}

	return stat, nil
}
