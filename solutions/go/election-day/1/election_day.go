// Package electionday provides utilities for counting votes
// and handling simple election results.
package electionday

import "fmt"

// NewVoteCounter returns a pointer to a new vote counter
// initialized with the given number of votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount returns the number of votes in counter.
// It returns 0 if counter is nil.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount adds increment to the value in counter.
// It does nothing if counter is nil.
func IncrementVoteCount(counter *int, increment int) {
	if counter == nil {
		return
	}
	*counter += increment
}

// NewElectionResult returns a pointer to a new ElectionResult
// for the named candidate with the given vote count.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult formats an election result as "Name (Votes)".
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decreases the vote count of candidate
// in results by one, if the candidate is present in the map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	if _, ok := results[candidate]; ok {
		results[candidate]--
	}
}
