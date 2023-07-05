// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package guardduty provides the client and types for making API
// requests to Amazon GuardDuty.
//
// Amazon GuardDuty is a continuous security monitoring service that analyzes
// and processes the following data sources: VPC Flow Logs, Amazon Web Services
// CloudTrail event logs, and DNS logs. It uses threat intelligence feeds (such
// as lists of malicious IPs and domains) and machine learning to identify unexpected,
// potentially unauthorized, and malicious activity within your Amazon Web Services
// environment. This can include issues like escalations of privileges, uses
// of exposed credentials, or communication with malicious IPs, URLs, or domains.
// For example, GuardDuty can detect compromised EC2 instances that serve malware
// or mine bitcoin.
//
// GuardDuty also monitors Amazon Web Services account access behavior for signs
// of compromise. Some examples of this are unauthorized infrastructure deployments
// such as EC2 instances deployed in a Region that has never been used, or unusual
// API calls like a password policy change to reduce password strength.
//
// GuardDuty informs you of the status of your Amazon Web Services environment
// by producing security findings that you can view in the GuardDuty console
// or through Amazon CloudWatch events. For more information, see the Amazon
// GuardDuty User Guide (https://docs.aws.amazon.com/guardduty/latest/ug/what-is-guardduty.html) .
//
// See https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28 for more information on this service.
//
// See guardduty package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/guardduty/
//
// # Using the Client
//
// To contact Amazon GuardDuty with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon GuardDuty client GuardDuty for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/guardduty/#New
package guardduty
