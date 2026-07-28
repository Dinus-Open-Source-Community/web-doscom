# MDM Launcher MVP - Technical Specification v1.0

## 1. Project Overview

### Project Name

MDM Launcher

### Objective

Develop an Android Device Owner based Mobile Device Management (MDM) platform that allows institutions to:

* Restrict device usage
* Enforce educational application access
* Measure learning activity
* Remotely manage Android devices
* Operate reliably in offline-first environments

### Target Platform

* Android 10+
* Android Enterprise
* Device Owner Mode
* Non-Root Devices

### Primary Customers

* Schools
* Homeschooling Organizations
* Training Centers
* Computer Labs
* Parents

---

# 2. System Architecture

## Components

### Android Client

Responsibilities:

* Device Owner enforcement
* Kiosk launcher
* Policy execution
* Local telemetry collection
* Offline operation
* OTA update handling

### Backend API

Responsibilities:

* Device registration
* Policy management
* Telemetry ingestion
* OTA distribution
* Authentication & authorization

### Dashboard

Responsibilities:

* Device monitoring
* Policy management
* User management
* OTA management

### PostgreSQL Database

Responsibilities:

* Tenant storage
* Device inventory
* Telemetry storage
* Policy storage
* Audit logging

---

# 3. Android Client Modules

## 3.1 Provisioning Module

Responsibilities:

* QR Enrollment
* Device Owner activation
* Initial registration

Inputs:

* Provisioning QR
* Staging Wi-Fi Profile

Outputs:

* Registered device
* Active Device Owner

Acceptance Criteria:

* Device Owner activated successfully
* Device registered in backend
* Launcher installed automatically

---

## 3.2 Network Bootstrap Module

Responsibilities:

* Connect to staging network
* Download institution Wi-Fi profile
* Auto migrate to production Wi-Fi

Acceptance Criteria:

* User interaction not required
* Automatic migration successful

---

## 3.3 Device Policy Controller (DPC)

Responsibilities:

* Apply Device Owner restrictions
* Manage Lock Task Mode
* Hide restricted packages
* Disable restricted system actions

Policies:

* DISALLOW_SAFE_BOOT
* DISALLOW_FACTORY_RESET
* DISALLOW_CREATE_WINDOWS

Restricted Packages:

* com.android.settings

Acceptance Criteria:

* Settings inaccessible
* Safe mode restricted
* Factory reset restricted

---

## 3.4 Kiosk Launcher

Responsibilities:

* Home screen replacement
* Whitelist rendering
* State-based application visibility

Features:

* App whitelist
* Hidden app drawer
* Lock Task enforcement

Acceptance Criteria:

* User cannot leave launcher
* Only allowed applications visible

---

## 3.5 State Engine

State Types:

### LOCKED

Allowed:

* Educational Apps

Blocked:

* Entertainment Apps

### REWARD

Allowed:

* Educational Apps
* Entertainment Apps

### IDLE

Triggered when:

* No interaction > 15 minutes

Acceptance Criteria:

* State transition performed locally
* State survives reboot

---

## 3.6 Activity Tracking Engine

Responsibilities:

* Measure active learning time

Data Sources:

Accessibility Events:

* TYPE_VIEW_CLICKED
* TYPE_VIEW_SCROLLED

Clock Source:

SystemClock.elapsedRealtime()

Idle Threshold:

15 minutes

Acceptance Criteria:

* Time manipulation does not affect timer
* Idle detection works offline

---

## 3.7 Watchdog Service

Responsibilities:

* Monitor Accessibility Service
* Monitor critical background services

Failure Response:

* Enter fail-secure mode
* Display full-screen lock overlay

Acceptance Criteria:

* Service failure detected within 30 seconds

---

## 3.8 Offline Emergency Access

Responsibilities:

* Local administrator access

Mechanism:

* Seven-corner-tap gesture
* TOTP authentication

Requirements:

* Offline operation
* No static PIN support

Acceptance Criteria:

* TOTP validated locally

---

## 3.9 OTA Update Module

Responsibilities:

* Version checking
* APK download
* Integrity verification
* Silent installation

Verification:

* SHA-256

Rollback Conditions:

* Crash > 3 times within 5 minutes

Retry Strategy:

* Exponential Backoff

Intervals:

* 1 hour
* 2 hours
* 4 hours
* 8 hours

Acceptance Criteria:

* Failed update automatically reverted

---

# 4. Local Storage

## Database

Technology:

Room Database (SQLite)

### Tables

devices

* id
* tenant_id
* device_uuid
* registered_at

policy_cache

* id
* version
* payload
* updated_at

telemetry

* id
* device_id
* event_type
* timestamp
* payload

system_logs

* id
* severity
* message
* timestamp

---

## Log Retention Policy

Maximum Size:

50 MB

Maximum Age:

30 days

Deletion Strategy:

FIFO

Acceptance Criteria:

* Storage exhaustion prevented

---

# 5. Backend Services

## Authentication Service

Functions:

* Login
* Session management
* RBAC enforcement

---

## Device Service

Functions:

* Device registration
* Device inventory
* Device status

---

## Policy Service

Functions:

* Policy creation
* Policy versioning
* Policy deployment

---

## Telemetry Service

Functions:

* Telemetry ingestion
* Device analytics

---

## OTA Service

Functions:

* APK hosting
* Version distribution
* Rollback tracking

---

# 6. API Specification

## Register Device

POST /api/v1/device/register

Request:

```json
{
  "device_uuid":"string",
  "institution_id":"uuid"
}
```

Response:

```json
{
  "device_id":"uuid",
  "status":"registered"
}
```

---

## Get Policy

GET /api/v1/policy/{device_id}

Response:

```json
{
  "policy_version":1,
  "policy":{}
}
```

---

## Telemetry Upload

POST /api/v1/telemetry

Request:

```json
{
  "device_id":"uuid",
  "events":[]
}
```

---

## OTA Check

GET /api/v1/ota/check

Response:

```json
{
  "version":"1.2.0",
  "sha256":"...",
  "download_url":"..."
}
```

---

# 7. Realtime Communication

## Primary Channel

Firebase Cloud Messaging (FCM)

Supported Actions:

```json
{
  "action":"SYNC_POLICY"
}
```

```json
{
  "action":"LOCK_DEVICE"
}
```

```json
{
  "action":"UNLOCK_DEVICE"
}
```

---

## Fallback Channel

WorkManager Polling

Interval:

15 minutes

Endpoints:

GET /api/v1/policy/check

Acceptance Criteria:

* Policy update eventually delivered even if FCM unavailable

---

# 8. Multi-Tenant Model

Hierarchy:

Institution
└── Group/Class
└── Device

Rules:

* Tenant isolation mandatory
* Cross-tenant access prohibited

---

# 9. RBAC

## Role: Super Admin

Permissions:

* Institution Management
* OTA Management
* Device Owner Operations
* Device Retirement
* TOTP Seed Management

---

## Role: Institution Admin

Permissions:

* Device Monitoring
* Schedule Management
* Policy Management
* Temporary Lock
* Temporary Unlock

Restrictions:

* No OTA Access
* No Device Owner Access
* No TOTP Seed Access

---

# 10. Device Lifecycle

States:

UNREGISTERED

↓

PROVISIONING

↓

REGISTERED

↓

ACTIVE

↓

OFFLINE

↓

RETIRED

Acceptance Criteria:

* Lifecycle changes auditable
* Every transition logged

---

# 11. Security Requirements

Protected Against:

* System clock manipulation
* Settings access
* App switching
* Split screen bypass
* Overlay bypass
* Safe mode bypass

Out of Scope:

* Root exploits
* Bootloader unlock
* Fastboot flashing
* Custom recovery
* Hardware attacks
* Kernel exploits

---

# 12. MVP Scope

Included:

* Device Owner
* Kiosk Launcher
* Policy Engine
* Telemetry
* OTA Update
* Dashboard
* RBAC
* Offline Mode

Excluded:

* DNS Filtering
* VPN Enforcement
* Mobile Threat Defense
* SIEM Integration
* DPI
* SSL Inspection
* Play Integrity Attestation
