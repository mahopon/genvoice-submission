# genvoice-submission

## Frontend - Hosted on Github Pages
> React-router with Antd and Axios

### Home
Shows writeup and links when not logged in. Shows two tables when logged in, one for surveys that user has created and another one for surveys that others have done. Users will be able to create surveys, which start from 1 question and can have more added.

1. For surveys that users have created, they are not able to enter their input, however they can see all answers by other users (I didn't have time to add in the identifier for the answers). They can also choose to delete the survey, which will delete all related questions and answers in the DB.
2. For surveys that others have created, users can record their voice input. After saving, they can go back to the survey to delete or re-record, which will override the audio.

### Register
Register will redirect the user if they have been authenticated before (they should not be able to reach this page!)

Validation is not done well here, only requiring to have values in the inputs, and not able to have same username as others.

No password policy is imposed here.

### Login
Login will redirect the user if they have been authenticated before (they should not be able to reach this page!)

### Settings
Settings will redirect the user if they have not been authenticated before (they should not be able to access this page)

No password policy is imposed here.

### Admin
Admin will redirect the user if 1. they have not been authenticated before and/or 2. they do not have admin role

Admin will be able to perform on other users:
1. Change username - Must always have value
2. Change name - Must always have value
3. Change password - Need not be filled, just won't be update
4. Change role - Select of either USER or ADMIN

Admins also can create new users within the admin page.

Be aware that deleting has no confirmation popup, and that deleting users will trigger cascade deletion of surveys and answers that they have made.

## Backend - Hosted on my DigitalOcean droplet
> Golang with Echo

### Authentication & Authorization
Done by JSON Web Tokens (JWT) over HTTPOnly Cookies (as best as I can, I wanted to try implementing it)

### Password hashing
Using ARGON2 hashing algorithm to secure passwords.

## Database - Hosted on my DigitalOcean droplet
Using PostgreSQL, accessing with GORM.

## Additional Notes
I am aware that the return status codes/messages are not being used properly. Due to a lack of time (incoming exams), I'm not able to find time to refactor the backend for proper use. Some of the API calls on the frontend are also not done well as a result of this.

I used a lot of GPT for the frontend as I'm not familiar with React (but learning) and don't think I could've done as much as I have without doing so.
