# Managing Intune Users, Groups, & Roles

* Prior to setting up Users, you should setup Groups.
* You have 100 people in Sales team, and want to grant them access to the Sales app
* You can just add the users to the group, so you allocate access to the app to the Group instead

## Intune Groups now user Azure AD Security Groups

## Group Members

1. Devices
2. Users

* You coud have a group based on user location and a different one for device location
* Or one that is for specific devices in a specific location

## Devices

* Any type of enrolled device
* Apply policies, or deploy apps for specific devices types
* All devices in device group will be affected by applied policy
* For example, create a phone group called Android phones, then apply Policy to that group

## Users

* Logical group based on a grouping that alrady occurs within your organization
* Apply a policy, or deploy an app to all users in a group
* For example, a Finance team group, all members in Finance would be there, you can even add their devices, or a seperate Finance team devices group

## Group Owner

* Membership of Azure AD groups is managed by the group owner
* Resource owner delegates to the owner of the group

## Parent and Child Group Attributes

```
Azure AD -.
Child - Child - Child
```

```
Azure AD -.
Android Devices - Finance Team -. Northern Sales
Finance Team Users -.
Finance Team Devices
```

## Group Membership Types

* Assigned
* Dynamic Users
* Dynamic Devices

### Assigned Groups have members manually assigned to them

* An administrator has to assign

### Dynamic Groups are assigned members based on the evaluation of a query - dynamically

* Assigned based on the evaluation of a query, dynamically, using Azure AD for group membership
* First create the rules, then if any attribute of a device or user changes, Azure AD will evaluate the rules
* Uses query operators
* Example where a user belongs to a Sales department, but not a member of Marketing

## Managing Groups
* Only administrators can manage and maintain groups
* Change groups including:
1. Add / remove members (users, or devices)
2. Rename a group
3. Create a subgroup
* To change Dynamic Groups you will alter the query criteria for the group

## Demo: Create and Manage Intune Assigned and Dynamic Groups

* From Azure Home Screen, Navigate to `Groups`
* You see many groups based on Office 365 subscription
* Select if you want an `Office 365` group which is a collaborative group
* We don't want that right now, we want a `Security`
* Enter a group name `SG-FinanceTeam-USA`
* Enter a description `Finance Team USA`
* Enter a membership type `Assigned`
* Click `Create` to add users or devices or groups

* Under `Users and Groups` blade, go to `All Groups`, then `Finance`
* Click `Properties`  here you can amend the name
* Click `Members` to add additional members
* Modify group `Owners` that are people outside of administrators who can modify stuff
* Look at Dynamic Groups under `Group memberships` you cannot create a rule that has user and device dynamic
* Create a new dynamic group security group:
1. Group name `SG-MiddleManagers`
2. Group desc `Finance Executive Management & Marketing`
3. Membership type `Dynamic User`
4. Setup the rules, etc.