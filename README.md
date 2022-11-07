# Send incident to OpsGenie

GitHub action sending custom incidents to opsgenie.

For the action to integrate with your OpsGenie installation you need an API key. The API key may be obtained from the organisations OpsGenie frontpage under **Settings** (upper right) -> **APP SETTINGS** (section to the left) -> **API key management**. The API key management requires a privileged account.

You also need a configured **team** as a recipient of the incidents. If you misconfigure this part the incident will be sent to your OpsGenie system and this script will report success, however since there are no matching recipients on the other end no one will get notified.

## Parameters

**MESSAGE** (optional)

Main message. If not provided a generic message will be provided stating the repository generating the alert.

**DESCRIPTION** (optional)

Will be part of the description section of the incident. The description already contains the short SHA and workflow name.

**PRIORITY** (required)

The priority level of your incident (P1 -> P5)

**API_KEY** (required)

Your OpsGenie secret API key. Use the repository or organization secrets as your key storage.

**TEAM** (required)
 
The team to address this incident to 

**TAG** (optional)

Custom tag

## Prefilled information

The incident report alreday contains decorated information consiting of:

**Title:**

The repository name (override by providing **MESSAGE**)

**Decription:**

Workflow name

Short SHA (7 characters)

**Details:**

Runner OS

Runner Arch

Branch

## Usage

Sending incidents to OpsGenie may be triggered by just failing actions or by advanced expressions. GitHub has put together a great guide: 

[GitHub action expressions](https://docs.github.com/en/actions/learn-github-actions/expressions)

#### Example Usage:

*Minimal:*

```yaml
- name: Send OpsGenie incident
  if: success()
  uses: tickup-se/notify_opsgenie@v2
  with:
    API_KEY: ${{ secrets.OPS_GENIE }}
    PRIORITY: 'P5'
    TEAM: 'existing team name' 
``` 

*All parameters:*

```yaml
- name: Send OpsGenie incident
  if: success()
  uses: tickup-se/notify_opsgenie@v2
  with:
    MESSAGE: 'some message'
    DESCRIPTION: 'some description'
    API_KEY: ${{ secrets.OPS_GENIE }}
    PRIORITY: 'P5'
    TEAM: 'existing team name'
    TAG: 'some tag'
``` 


# Incident example on an iPhone

![Incident Example](example.png)
