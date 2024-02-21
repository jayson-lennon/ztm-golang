package main

import "testing"
import "github.com/stretchr/testify/assert"

// // Days of the week
// const (
// 	Monday    = 0
// 	Tuesday   = 1
// 	Wednesday = 2
// 	Thursday  = 3
// 	Friday    = 4
// 	Saturday  = 5
// 	Sunday    = 6
// )

func TestCheckAccess(t *testing.T) {
	// Admin / manager
	for i := 0; i <= 6; i++ { // <=6 will include every day
		assert.Equal(t, CheckUserAccess(i, Admin), true, "Admin should have access on %v", i)
		assert.Equal(t, CheckUserAccess(i, Manager), true, "Manager should have access on %v", i)
	}
	// Contractor
	for i := 0; i <= 4; i++ { // <=4 will only include weekdays
		assert.Equal(t, CheckUserAccess(i, Contractor), false, "Contractor should not have access on %v", i)
	}
	assert.Equal(t, CheckUserAccess(Saturday, Contractor), true, "Contractor should have access on saturday")
	assert.Equal(t, CheckUserAccess(Sunday, Contractor), true, "Contractor should have access on sunday")

	// Member
	for i := 0; i <= 4; i++ { // <=4 will only include weekdays
		assert.Equal(t, CheckUserAccess(i, Member), true, "Member should have access on %v", i)
	}
	assert.Equal(t, CheckUserAccess(Saturday, Member), false, "Member should not have access on Saturday")
	assert.Equal(t, CheckUserAccess(Sunday, Member), false, "Member should not have access on Sunday")

	// Guest
	assert.Equal(t, CheckUserAccess(Monday, Guest), true, "Guest should have access on Monday")
	assert.Equal(t, CheckUserAccess(Wednesday, Guest), true, "Guest should have access on Wednesday")
	assert.Equal(t, CheckUserAccess(Friday, Guest), true, "Guest should have access on Friday")

	assert.Equal(t, CheckUserAccess(Tuesday, Guest), false, "Guest should not have access on Tuesday")
	assert.Equal(t, CheckUserAccess(Thursday, Guest), false, "Guest should not have access on Thursday")
	assert.Equal(t, CheckUserAccess(Saturday, Guest), false, "Guest should not have access on Saturday")
	assert.Equal(t, CheckUserAccess(Sunday, Guest), false, "Guest should not have access on Sunday")
}
