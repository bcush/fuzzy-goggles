# Management of Desktop Devices

* Desktops enrolled as if they were cell phones
* Once client is installed, it can provide detailed inventory of hardware/software
* Install apps, malware scans, data wipes, update firewall settings, etc.
* Unlike phones, you cannot do a full wipe

## Installing the Intune PC Client

* Three ways:
1. Administrator Deployment
2. Image Installation
3. User Initiated Enrollment

### Administrator Deployment

* Administrator downloads the Intune Client software from the Intune Classic console
* Manually install the software on target device:
1. Microsoft_Intune_Setup.exe
2. MicrosoftIntune.accountcert
* Group Policy can automate installation

### Image Installation

* Install Intune client as part of an image
* Install to multiple computers
* PCs will be automatically enrolled in Intune after Windows Setup finishes

### User Initiated Enrollment

* End users self-enroll their PCs
* Intune client available from Company Portal website
* Users need to choose the client software option

## Bandwidth Implications for the Intune Client

* Expect about 200MB for client enrollment package
* Malware stuff will get updated, and policies
* Might take a few hours before it appears available to be managed from the Intune console

## Summary

* Setup tasks
1. Adding a custom domain name
2. Adding users, assigning licenses, creating groups
3. Administrator roles
* Configure devices, creating policies
* Alerts and notifications
* Company Portal
* Enabling device enrollment
* Manage desktop devices
* Installing the Intune client
