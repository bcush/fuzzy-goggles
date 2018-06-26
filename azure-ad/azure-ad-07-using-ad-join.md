# Leveraging Azure AD Join and Azure AD Domain Services

* Supported by Windows 10 and builds on `Workplace Join` that was in 2012 R2
* A machine becomes a _known asset_; joined but not in traditional sense
* For example, no GPO, no strict management, but basic governance
* You can however also apply policy and apply software

## Azure AD Join aimed at devices leveraging cloud services and managed via MDM

* You want to use this in the following areas:

1. Your apps and services are mostly in the cloud.
2. Your users use Office 365, or other partner account services.
3. You have part-time, high turnover employees like seasonal or students
4. Some organizations are just into bring your own device models

## Must be enabled per Azure AD instance

* Log into Azure to review these settings
`Portal -> Azure AD -> Devices -> Device settings -> Users may [...]`

* You can use `All Users` or a subset of users
* You may specify who are `Administrators` on those Azure joined devices
* You can setup whether `multi-factor authentication` is required
* You may configure the number of devices that a user can add
* You can specify MDM or Office 365 applications

### Users May Register Their Devices with Azure AD

* Important, you can also set who can register with Intune

## In Windows 10 the option to Join Azure AD can be used

* On a Windows 10 Machine go to `Settings` then `System` then go to `About`
* Finally click `Join Azure AD` which will ask for creds, or MFA if set
* Type in your account and password
* At this point it will check to see if there are policies set with Intune
* You will receive a warning if there are changes to take effect
* Click `Finish` and it is now added to your Azure AD instance

## Users can see and perform actions on all their registered devices

* Check the devices registered to an account
`Portal -> Azure AD -> Users -> ${user} -> Devices`
* From here you can also `Block Device` or `Delete Device`
