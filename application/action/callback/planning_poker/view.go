package planning_poker

import (
	"fmt"
	"planning_pocker_bot/domain/entity"
	"strings"
)

func viewVotes(voters []entity.Vote) string {
	votersView := make([]string, len(voters))
	for i, vote := range voters {
		votersView[i] = fmt.Sprintf("%s %s (@%s)", vote.Icon, vote.FirstName, vote.Name)
	}

	return strings.Join(votersView, "\n")
}

func viewReveal(voters []entity.Vote) string {
	votersView := make([]string, len(voters))
	for i, vote := range voters {
		votersView[i] = fmt.Sprintf("%s %s (@%s) → <b>%s</b>", vote.Icon, vote.FirstName, vote.Name, vote.Value)
	}

	return strings.Join(votersView, "\n")
}
