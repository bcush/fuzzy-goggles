# Managing Intune Users

## Intune User Types

* End Users
* Administrators

### End Users

* Can self-enroll devices in Intune
* Manage their Intune enrolled devices
* Access apps, data and support help from the Company Portal
* Enroll up to five devices

### Administrators

* Setup, manage and maintain Intune
* Manage and support users and devices
* First Administrator is the Global Administrator
* Administrators add users, roles, create groups, and remove devices and manage the Company Portal

## Adding Users

* Administrators can add users Office 365 portal or the Azure Portal
* Enter their details such as name, username and profile
* Allocate the user to groups
* Add roles
* Assign licenses

## Adding new user accounts using the Azure Intune Portal

* Select `Intune` from left hand side
* Select `Users` and click `New User`
* Enter the name and the username
* Name is `Don Pilling` and user name is `dpilling@yaya.com`
* Under Profile you can enter more detailed info
* Click `Groups` and assign him to `Finance` group
* If you want him to be admin, select your `Directory role`
* Click `Create` and your new user is added to `Intune`
* Need to add a license for this user
* Search for Dom and then click Licenses then Assign
* Click the EM+ES and Office 365 and click Select
* Then on assignment options, fine tune what Dom can use
* The user can now auto-enroll their devices into Intune