# file: features/notification.feature

# http://localhost:8088/
# http://reg_service:8088/

Feature: Email notification sending
	As API client of registration service
	In order to understand that the user was informed about registration
	I want to receive event from notifications queue

	Scenario: Registration service is available
		When I send "GET" request to "http://reg_service:8088/"
		Then The response code should be 200
		And The response should match text "OK"

	Scenario: Notification event is received
		When I send "POST" request to "http://reg_service:8088/api/v1/registration" with "application/json" data:
		"""
		{
			"first_name": "otus",
			"email": "otus@otus.ru",
			"age": 27
		}
		"""
		Then The response code should be 200
		And I receive event with text "otus@otus.ru"
