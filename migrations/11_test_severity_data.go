package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func testSeverityDataUp(db dbx.Builder) error {
	dao := daos.New(db)

	// Get the tickets collection
	ticketsCollection, err := dao.FindCollectionByNameOrId(TicketCollectionName)
	if err != nil {
		return err
	}

	// Get the vulnerability type
	vulnerabilityType, err := dao.FindRecordById(TypeCollectionName, "vulnerability")
	if err != nil {
		return err
	}

	// Create test tickets with different severities
	testTickets := []struct {
		name        string
		severity    string
		description string
	}{
		{"SQL Injection in Login Form", "critical", "Critical vulnerability found in authentication system"},
		{"Cross-Site Scripting (XSS) Vulnerability", "critical", "XSS vulnerability in user input fields"},
		{"Outdated SSL Certificate", "high", "SSL certificate needs immediate renewal"},
		{"Weak Password Policy", "high", "Current password requirements are insufficient"},
		{"Missing Security Headers", "high", "HTTP security headers not properly configured"},
		{"Unpatched Dependencies", "medium", "Several npm packages need security updates"},
		{"Exposed API Endpoints", "medium", "Some API endpoints lack proper authentication"},
		{"Information Disclosure", "medium", "Error messages reveal sensitive system information"},
		{"Missing Rate Limiting", "low", "API endpoints should implement rate limiting"},
		{"Outdated Documentation", "low", "Security documentation needs updating"},
		{"Verbose Error Messages", "low", "Error messages could be more generic"},
	}

	for _, ticket := range testTickets {
		record := models.NewRecord(ticketsCollection)
		record.Set("name", ticket.name)
		record.Set("type", vulnerabilityType.Id)
		record.Set("description", ticket.description)
		record.Set("open", true)
		record.Set("severity", ticket.severity)

		if err := dao.SaveRecord(record); err != nil {
			return err
		}
	}

	return nil
}

func testSeverityDataDown(db dbx.Builder) error {
	// No need to clean up test data in down migration
	return nil
}
