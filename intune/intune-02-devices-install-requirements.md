# Devices and Installation Requirements

## Instalation Requirements

* Firewall Access
* Supported Operating Systems
* Supported Web Browsers

## Firewall Access

* You need to be able to communicate to cloud
* Firewalls & Proxy Server require access to

1. manage.microsoft.com (80/443)
2. microsoft.com (80/443)
3. social.technet.microsoft.com (80)
4. www.microsoft.com

### Older Clients

* Download the Intune client installation
* Download Client Enrollment Package

## Various other Bandwidth Requirements (Opt)

1. Endpoint Protection Agent - 65 MB (one time)
2. Operations Manageer Agent - 11 MB (one time)
3. Policy Agent - 3 MB (one time)
4. Remote Assistance via Microsoft Easy Assist agent - 6 MB (one time)
5. Endpoint Protection engine update - 5 MB (monthly)

## Additional Requirements

* Administrative permissions
* Windows Installer 3.1
* Remove SCCM or SMS

_From April 2018, Microsoft will turn off mobile
device management in the classic Intune console
for customers using Intune standalone_

## Methods for Installation of Intune PC Client

* Administrator deployment (physically sitting there)
* User Initiated Enrollment (end users enroll themselves)
* Image Installation (when an image is capture from an existing PC)

## Supported Windows Operating Systems

* You can manage modern OSes with installing client:

1. Windows 10 (Home, S mode, Pro, Education, Ent)
2. Windows 10 Mobile
3. Windows 10 IoT Ent, etc.

## Supported Web Browsers

1. Microsoft Edge
2. IE 11
3. Safari
4. Chrome
5. Firefox