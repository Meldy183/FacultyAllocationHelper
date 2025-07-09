# Acceptance tests

## Simplify the course's card, add information which track it belongs to

Issue #103 - Accepted by the customer

**AC**
- **GIVEN**: I am an authorized user with access to courses info
- **WHEN**: I am viewing a course's card
- **THEN**: I can see which program the course belongs to

## Transform course/TA management forms to pop-up windows without redirections

Issue #104 - Accepted by the customer

**AC**
- **GIVEN**: I am an authorized user able to manage TAs/courses
- **WHEN**: I am updating the info about a course/TA
- **THEN**: I am not redirected to another page AND management is hold in a pop-up window

## Refactor workload from hours to percent of working rate

Issue #117 - Accepted by the customer

**AC**
- **GIVEN**: I am an authorized user able to update faculty members
- **WHEN**: I update workload of a faculty member
- **THEN**: The workload is measured in percents and not in hours

## Implement the prediction of accordance of a faculty for a course

Discussed during the meeting, not planned yet

**AC**
- **GIVEN**: I am an authorized user able to manage TAs/courses
- **WHEN**: I am using a list of faculties to assign to a course
- **THEN**: I see the faculties ordered by the accordance which is calculated by special formula

## Implement a "time machine"

Discussed during the meeting, not planned yet

**AC**
- **GIVEN**: I am an authorized user able to manage TAs/courses
- **WHEN**: I am viewing the info about the current year
- **THEN**: I am able to switch from the current year to one of the previous and view the same information about it but without ability to edit