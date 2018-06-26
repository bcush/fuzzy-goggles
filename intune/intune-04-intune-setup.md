# Intune Setup

## Intune Setup Tasks (New Subscription)

1. Signing-up
2. Custom Domain Name
3. Users and Roles
4. Licenses
5. Groups
6. Administrators

## Signing Up for Intune

* Sign up for a new Intune subscription or add it to an existing
  Microsoft subscription
* User account which sets up Intune is setup up as the Global Administrator
* "yourcompany.onmicrosoft.com"
* Overlary a custom domain name
* https://docs.microsoft.com/intune/free-trial-sign-up

## Adding a Custom Domain Name

* Many organizations use a custom domain name for branding purposes
* Other Microsoft services may already use your custom domain name
* Effectively this is an alias on top of onmicrosoft.com domain

## Demo

### Assigning a custom domain name to Intune

* Login to the Azure Intune Portal
* Add your Domain
* Verify records in the DNS register

## Users and Roles

### Adding Users

* Add users directory using the Azure Intune Portal
* Automatically sync on-premises domain users using Azure AD Connect
* Bulk upload user data using .csv files and scripting tools such as PowerShell

### Administrators

* Global Administrator
* Other admin built-in roles include:
1. Intune Service Administrator
2. Conditional Access Administrator
3. User Administrator
* Assign admin roles to user account
* You cannot delete or edit the built-in roles

### Intune Roles

* Grant permissions to users using Roles
* Each role has:
1. Role definition
2. Members
3. Scope
4. Assignment

* Intune built-in roles include:
1. Intune Role Administrator
2. Help Desk Operator
3. Policy and Profile Manager
4. Read-only Operator
5. Application Manager

## Licenses

* Each user requires a license to access Intune services
* Assign licenses using the Azure Intune Portal or in bulk by using scripts
* Enterprise Mobility + Security Intune licenses can be assigned in the Office 365 Admin Portal
* Purchase extra licenses or revoke a license so that it can be reused
* https://www.microsoft.com/cloud-platform/microsoft-intune-pricing

## Groups

* Security and Office 365 groups can be created for both users and devices
* Can include users, device and other grops
* Group membership can be allocated or dynamically assigned based on queried criteria

## Additional Intune Setup Tasks

* Apps
* Configure Devices
* Policies
* Alerts and Notifications
* Company Portal
* Enable Device Enrollment

### Apps

* Manage and assign apps to users
* Intune supports managing apps on Windows, iOS, and Android devices
* Type of apps available:
1. web apps
2. in-house
3. purchased apps
* Categories help organize apps

### Intune Device Configuration

* Use Intune device profiles to manage settings and features of enrolled devices
* Device profiles manage:
1. Features and Restrictions
2. Email, Wi-Fi and VPN access
3. Education settings
4. Certificates
5. Upgrades
6. Endpoint Protection and Windows Information Protection

### Policies

* Policies protect corporate data and to ensure devices are compliant
* Only your users access resources in a secure manner
* Different policies available:
1. Configuration
2. Device Compliance
3. Conditional Access
4. Corporate Device Enrollment
5. Resource Access

### Alerts and Notifications

* Hundreds of alerts available in the classic portal
* Alerts for:
1. Monitoring
2. Policy
3. System & Updates
* View alerts by date, category and severity
* Create alert notifications
* Set notification rules

### Setting up the Company Portal
* End user portal website or app
* Users enroll and manage devices
* Users access apps and support
* Company Portal can be customized with their branding and logo

### Enable Device Enrollment

* Need to configure MDM authority to allow devices to be enrolled
* Decide if you will use Intune as a standalone system, as a hybrid one, or as Office 365 with MDM features
* Need to set up an enrollment server too
* Can use DEM to enroll up to 1,000 devices at once

### Device Enrollment Manager

* Users can enroll only 5 devices
* Device Enrollment Manager account (DEM) can be used to enroll up to 1,000 devices
* Used by Administrators when setting up Intune, or when provisioning a new set of devices