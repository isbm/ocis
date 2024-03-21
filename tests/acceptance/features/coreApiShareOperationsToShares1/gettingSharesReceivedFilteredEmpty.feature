@skipOnReva
Feature: get the received shares filtered by type (user, group etc)
  As a user
  I want to filter the shares that I have received of a particular type (user, group etc)
  So that I can know about the status of the shares I've received

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
    And group "grp1" has been created
    And user "Brian" has been added to group "grp1"
    And user "Alice" has created folder "/folderToShareWithUser"
    And user "Alice" has created folder "/folderToShareWithGroup"
    And user "Alice" has created folder "/folderToShareWithPublic"
    And user "Alice" has uploaded file with content "file to share with user" to "/fileToShareWithUser.txt"
    And user "Alice" has uploaded file with content "file to share with group" to "/fileToShareWithGroup.txt"
    And user "Alice" has uploaded file with content "file to share with public" to "/fileToShareWithPublic.txt"


  Scenario Outline: getting shares received from users when there are none
    Given using OCS API version "<ocs-api-version>"
    And user "Alice" has shared folder "/folderToShareWithGroup" with group "grp1"
    And user "Alice" has created a public link share with settings
      | path        | /folderToShareWithPublic |
      | permissions | read                     |
      | password    | %public%                 |
    And user "Alice" has shared file "/fileToShareWithGroup.txt" with group "grp1"
    And user "Alice" has created a public link share with settings
      | path        | /fileToShareWithPublic.txt |
      | permissions | read                       |
      | password    | %public%                   |
    When user "Brian" gets the user shares shared with him using the sharing API
    Then the OCS status code should be "<ocs-status-code>"
    And the HTTP status code should be "200"
    And no files or folders should be included in the response
    Examples:
      | ocs-api-version | ocs-status-code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: getting shares received from groups when there are none
    Given using OCS API version "<ocs-api-version>"
    And user "Alice" has shared folder "/folderToShareWithUser" with user "Brian"
    And user "Alice" has created a public link share with settings
      | path        | /folderToShareWithPublic |
      | permissions | read                     |
      | password    | %public%                 |
    And user "Alice" has shared file "/fileToShareWithUser.txt" with user "Brian"
    And user "Alice" has created a public link share with settings
      | path        | /fileToShareWithPublic.txt |
      | permissions | read                       |
      | password    | %public%                   |
    When user "Brian" gets the group shares shared with him using the sharing API
    Then the OCS status code should be "<ocs-status-code>"
    And the HTTP status code should be "200"
    And no files or folders should be included in the response
    Examples:
      | ocs-api-version | ocs-status-code |
      | 1               | 100             |
      | 2               | 200             |


  Scenario Outline: getting shares received from public links when there are none
    # Note: public links are purposely created in this scenario
    #       users do not receive public links, so asking for a list of public links
    #       that are "shared with me" should always return an empty list.
    Given using OCS API version "<ocs-api-version>"
    And user "Alice" has shared folder "/folderToShareWithUser" with user "Brian"
    And user "Alice" has shared folder "/folderToShareWithGroup" with group "grp1"
    And user "Alice" has created a public link share with settings
      | path        | /folderToShareWithPublic |
      | permissions | read                     |
      | password    | %public%                 |
    And user "Alice" has shared file "/fileToShareWithUser.txt" with user "Brian"
    And user "Alice" has shared file "/fileToShareWithGroup.txt" with group "grp1"
    And user "Alice" has created a public link share with settings
      | path        | /fileToShareWithPublic.txt |
      | permissions | read                       |
      | password    | %public%                   |
    When user "Brian" gets the public link shares shared with him using the sharing API
    Then the OCS status code should be "<ocs-status-code>"
    And the HTTP status code should be "200"
    And no files or folders should be included in the response
    Examples:
      | ocs-api-version | ocs-status-code |
      | 1               | 100             |
      | 2               | 200             |
